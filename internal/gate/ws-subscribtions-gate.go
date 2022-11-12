package gate

import (
	"crypto/configs"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
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

func GateWSClient() {

	msg := WSNotificationEvent{}

	u := url.URL{Scheme: configs.Conf.Gate.WS.Schema,
		Host: configs.Conf.Gate.WS.Host,
		Path: configs.Conf.Gate.WS.Path}
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{RootCAs: nil, InsecureSkipVerify: true}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	c.SetPingHandler(nil)

	// read msg
	go func() {
		for {
			// _, message, err := c.ReadMessage()
			err := c.ReadJSON(&msg)
			if err != nil {
				c.Close()
				panic(err)
			}
			// c.ReadJSON()
			// _ = json.Unmarshal(message, &msg)
			// fmt.Printf("recv: %v\n", msg)
			switch {
			case msg.Channel == "spot.order_book":
				fmt.Printf("recv: %v\n", msg)
				orderBook, ok := msg.Result.(OrderBookEventResults)
				if !ok {
					fmt.Printf("recv: %v\n", ok)
				}
				fmt.Printf("recv: %v\n", orderBook.Asks)
			case msg.Channel == "spot.balances":
				fmt.Printf("recv: %v\n", msg)

			}
		}
	}()

	t := time.Now().Unix()
	pingMsg := NewWSMsg("spot.ping", "", t, []string{})
	err = pingMsg.send(c)
	if err != nil {
		panic(err)
	}

	// subscribe order book
	orderBookMsg := NewWSMsg("spot.order_book", "subscribe", t, []string{"ETH_USDT", "5", "100ms"})
	err = orderBookMsg.send(c)
	if err != nil {
		panic(err)
	}

	// // subscribe positions
	// ordersMsg := NewWSMsg("spot.orders", "subscribe", t, []string{"BTC_USDT"})
	// ordersMsg.signWSMsg()
	// err = ordersMsg.send(c)
	// if err != nil {
	// 	panic(err)
	// }

	select {}
}
