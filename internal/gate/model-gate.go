package gate

import (
	"crypto/configs"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

type errorResponse struct {
	Label   string `json:"label"`
	Message string `json:"message"`
}

// Client struct
type GateClient struct {
	Host          string
	Prefix        string
	Endpoints     configs.GateEndpoints
	CommonHeaders configs.GateCommonHeaders
	Pair          string
	HTTPClient    *http.Client
	WSConn        *websocket.Conn
}

type CurrencyPairDetails struct {
	ID              string `json:"id"`
	Base            string `json:"base"`
	Quote           string `json:"quote"`
	Fee             string `json:"fee"`
	MinBaseAmount   string `json:"min_base_amount"`
	MinQuoteAmount  string `json:"min_quote_amount"`
	AmountPrecision int    `json:"amount_precision"`
	Precision       int    `json:"precision"`
	TradeStatus     string `json:"trade_status"`
	SellStart       int    `json:"sell_start"`
	BuyStart        int    `json:"buy_start"`
}

type orderBook struct {
	ID      int        `json:"id"`
	Current int64      `json:"current"`
	Update  int64      `json:"update"`
	Asks    [][]string `json:"asks"`
	Bids    [][]string `json:"bids"`
}

type Balance struct {
	BTC  string `json:"btc"`
	ETH  string `json:"eth"`
	USDT string `json:"usdt"`
}

type CurrencyChain []struct {
	Chain              string `json:"chain"`
	NameCn             string `json:"name_cn"`
	NameEn             string `json:"name_en"`
	IsDisabled         int    `json:"is_disabled"`
	IsDepositDisabled  int    `json:"is_deposit_disabled"`
	IsWithdrawDisabled int    `json:"is_withdraw_disabled"`
}

type WithdrawalRecords []struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Currency  string `json:"currency"`
	Address   string `json:"address"`
	Txid      string `json:"txid"`
	Amount    string `json:"amount"`
	Memo      string `json:"memo"`
	Status    string `json:"status"`
	Chain     string `json:"chain"`
}

type TotalBalance struct {
	Details Details `json:"details"`
	Total   Total   `json:"total"`
}

type Total struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type SpotBalance []struct {
	Currency  string `json:"currency"`
	Available string `json:"available"`
	Locked    string `json:"locked"`
}

type Details struct {
	CrossMargin CrossMargin `json:"cross_margin"`
	Spot        Spot        `json:"spot"`
	Finance     Finance     `json:"finance"`
	Margin      Margin      `json:"margin"`
	Quant       Quant       `json:"quant"`
	Futures     Futures     `json:"futures"`
	Delivery    Delivery    `json:"delivery"`
	Warrant     Warrant     `json:"warrant"`
	Cbbc        Cbbc        `json:"cbbc"`
}

type CrossMargin struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Spot struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Finance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Margin struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Quant struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Futures struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Delivery struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Warrant struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Cbbc struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Prices struct {
	Sells       float64
	SellsVolume float64
	Buys        float64
	BuysVolume  float64
}

type WSUpdateNotification struct {
	Time    int             `json:"time"`
	Channel string          `json:"channel"`
	Event   string          `json:"event"`
	Result  json.RawMessage `json:"result"`
}

type WalletUpdateNotification []struct {
	Timestamp   string `json:"timestamp"`
	TimestampMs string `json:"timestamp_ms"`
	User        string `json:"user"`
	Currency    string `json:"currency"`
	Change      string `json:"change"`
	Total       string `json:"total"`
	Available   string `json:"available"`
}

type orderBookUpdateNotification struct {
	T            int64      `json:"t"`
	LastUpdateID int64      `json:"lastUpdateId"`
	S            string     `json:"s"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type WSMsg struct {
	Time    int64    `json:"time"`
	Channel string   `json:"channel"`
	Event   string   `json:"event"`
	Payload []string `json:"payload"`
	Auth    *Auth    `json:"auth"`
}

type Auth struct {
	Method string `json:"method"`
	KEY    string `json:"KEY"`
	SIGN   string `json:"SIGN"`
}
