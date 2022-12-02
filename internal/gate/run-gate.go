package gate

import (
	"crypto/configs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

const (
	numRetry int = 5
)

var (
	PriceInfo   SafePrices
	BalanceInfo SafeBalance = SafeBalance{Balance: Balance{}}

	msg WSUpdateNotification
)

func NewWSMsg(channel, event string, t int64, payload []string) *WSMsg {
	return &WSMsg{
		Time:    t,
		Channel: channel,
		Event:   event,
		Payload: payload,
	}
}

func (msg *WSMsg) send(c *websocket.Conn) error {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return c.WriteMessage(websocket.TextMessage, msgByte)
}

func (c *GateClient) subscribeOrderBook(pair string, unixTime int64) error {
	var err error

	// subscribe order book
	orderBookMsg := NewWSMsg("spot.order_book", "subscribe", unixTime, []string{pair, "5", "100ms"})
	log.Println("|| GATE || SEND WS || -> : Sending subscribe message to spot.order_book channel ... ")
	err = orderBookMsg.send(c.WSConn)
	if err != nil {
		log.Printf("|| GATE || SEND WS || -> : Failed to send subscribe message to spot.order_book channel %v\n", err)
		return err
	}
	return nil
}

func (c *GateClient) subscribeBalances(unixTime int64) error {
	var err error
	// subscribe positions
	ordersMsg := NewWSMsg("spot.balances", "subscribe", unixTime, []string{})
	ordersMsg.signWSMsg()
	log.Println("|| GATE || SEND WS || -> : Sending subscribe message to spot.balances channel ... ")
	err = ordersMsg.send(c.WSConn)
	if err != nil {
		log.Printf("|| GATE || SEND WS || -> : Failed to send subscribe message to spot.balances channel %v\n", err)
		return err
	}
	return nil
}

func (c *GateClient) reconnect(maxRetryConn int) error {
	stop := false
	retry := 0
	u := url.URL{Scheme: configs.Conf.Gate.WS.Schema,
		Host: configs.Conf.Gate.WS.WSHost,
		Path: configs.Conf.Gate.WS.Path}
	for !stop {
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			if retry >= maxRetryConn {
				log.Printf("max reconnect time %d reached, give it up", maxRetryConn)
				return err
			}
			retry++
			log.Printf("failed to connect to server for the %d time, try again later", retry)
			time.Sleep(time.Millisecond * (time.Duration(retry) * 500))
			continue
		} else {
			stop = true
			c.WSConn = conn
		}
	}
	t := time.Now().Unix()
	err := c.subscribeOrderBook(c.Pair, t)
	if err != nil {
		log.Println(err)
		return err
	}

	err = c.subscribeBalances(t)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func listenWSChan(msgCh <-chan *WSUpdateNotification, ob *safeOrderBook, prices *SafePrices, balance *SafeBalance, quitWs <-chan bool) {
	var (
		orderBookNotification  orderBookUpdateNotification
		spotWalletNotification WalletUpdateNotification
	)

	for {
		msg := <-msgCh

		switch {
		case msg.Channel == "spot.order_book":
			err := json.Unmarshal(msg.Result, &orderBookNotification)
			if err != nil {
				log.Printf("recv: Could NOT parse ws update notification to order book update notification %v\n", err)
			}
			ob.updateOrderBookFromWS(&orderBookNotification)
			prices.updatePrices(ob)

		case msg.Channel == "spot.balances":
			err := json.Unmarshal(msg.Result, &spotWalletNotification)
			if err != nil {
				log.Printf("recv: Could NOT parse ws update notification to wallet update notification %v\n", err)
			}
			balance.updateBalanceFromWS(&spotWalletNotification)
		case <-quitWs:
			return
		}
	}

}

func (c *GateClient) readMessages(ch chan<- *WSUpdateNotification, errs chan<- error, quitWs chan<- bool) {
	for {
		_, message, err := c.WSConn.ReadMessage()
		c.WSConn.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			ne, hasNetErr := err.(net.Error)
			noe, hasNetOpErr := err.(*net.OpError)
			if websocket.IsUnexpectedCloseError(err) || (hasNetErr && ne.Timeout()) || (hasNetOpErr && noe != nil) ||
				websocket.IsCloseError(err) || io.ErrUnexpectedEOF == err {

				log.Printf("websocket err:%s", err.Error())
				if e := c.reconnect(10); e != nil {
					log.Printf("reconnect err:%s", err.Error())
					errs <- fmt.Errorf("reconnect err:%s", err.Error())
					quitWs <- true
					return
				} else {
					log.Printf("reconnect success, continue read message")
					continue
				}
			} else {
				log.Printf("wsRead err:%s, type:%T", err.Error(), err)
				errs <- fmt.Errorf("wsRead err:%s, type:%T", err.Error(), err)
				quitWs <- true
				return
			}
		}

		err = json.Unmarshal(message, &msg)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal msg into WSUpdateNotification: %v", err)
			log.Println(err)
			errs <- err
		}
		ch <- &msg
	}

}

func (c *GateClient) RunGate(prices *SafePrices, balance *SafeBalance, errs chan error) {
	msgCh := make(chan *WSUpdateNotification)
	quitWs := make(chan bool)
	var (
		ob  safeOrderBook = safeOrderBook{orderBook: orderBook{}}
		err error
	)

	c.WSConn.SetPingHandler(nil)

	// Initial call to obtain balance
	err = c.GetSpotWalletBalanceDetails(balance)
	if err != nil {
		log.Println(err)
		errs <- err
	}

	t := time.Now().Unix()
	err = c.subscribeOrderBook(c.Pair, t)
	if err != nil {
		log.Println(err)
		errs <- err
	}

	err = c.subscribeBalances(t)
	if err != nil {
		log.Println(err)
		errs <- err
	}

	// read msg
	go c.readMessages(msgCh, errs, quitWs)
	go listenWSChan(msgCh, &ob, prices, balance, quitWs)

	if <-quitWs {
		go c.CallGetOrderBookDetails(&ob, prices, errs)
		go c.CallSpotWalletBalanceDetails(balance, errs)
	}
}
