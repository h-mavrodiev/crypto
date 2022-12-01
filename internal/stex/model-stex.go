package stex

import (
	"crypto/configs"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type successResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Client struct
type StexClient struct {
	Host          string
	Endpoints     configs.StexEndpoints
	CommonHeaders configs.StexCommonHeaders
	Pair          int
	HTTPClient    *http.Client
}

type Prices struct {
	Sells       float64
	SellsVolume float64
	Buys        float64
	BuysVolume  float64
}

type Settings struct {
	GoogleAuthEnabled           bool `json:"google_auth_enabled"`
	SmsOtpEnabled               bool `json:"sms_otp_enabled"`
	SecurityKeyEnabled          bool `json:"security_key_enabled"`
	EmailEncryptionEnabled      bool `json:"email_encryption_enabled"`
	ConfirmWithdrawalsWith2Fa   bool `json:"confirm_withdrawals_with_2fa"`
	ConfirmWithdrawalsWithEmail bool `json:"confirm_withdrawals_with_email"`
}

type CurrencyPairDetails []struct {
	ID                int    `json:"id"`
	CurrencyID        int    `json:"currency_id"`
	CurrencyCode      string `json:"currency_code"`
	CurrencyName      string `json:"currency_name"`
	MarketCurrencyID  int    `json:"market_currency_id"`
	MarketCode        string `json:"market_code"`
	MarketName        string `json:"market_name"`
	MinOrderAmount    string `json:"min_order_amount"`
	MinBuyPrice       string `json:"min_buy_price"`
	MinSellPrice      string `json:"min_sell_price"`
	BuyFeePercent     string `json:"buy_fee_percent"`
	SellFeePercent    string `json:"sell_fee_percent"`
	Active            bool   `json:"active"`
	Delisted          bool   `json:"delisted"`
	Message           string `json:"message"`
	CurrencyPrecision int    `json:"currency_precision"`
	MarketPrecision   int    `json:"market_precision"`
	Symbol            string `json:"symbol"`
	GroupName         string `json:"group_name"`
	GroupID           int    `json:"group_id"`
	AmountMultiplier  int    `json:"amount_multiplier"`
	TradingPrecision  int    `json:"trading_precision"`
}

type orderBook struct {
	Ask            Ask     `json:"ask"`
	Bid            Bid     `json:"bid"`
	AskTotalAmount float64 `json:"ask_total_amount"`
	BidTotalAmount float64 `json:"bid_total_amount"`
}

type Ask []struct {
	CurrencyPairID   int     `json:"currency_pair_id"`
	Amount           string  `json:"amount"`
	Price            string  `json:"price"`
	Amount2          string  `json:"amount2"`
	Count            int     `json:"count"`
	CumulativeAmount float64 `json:"cumulative_amount"`
}

type Bid []struct {
	CurrencyPairID   int     `json:"currency_pair_id"`
	Amount           string  `json:"amount"`
	Price            string  `json:"price"`
	Amount2          string  `json:"amount2"`
	Count            int     `json:"count"`
	CumulativeAmount float64 `json:"cumulative_amount"`
}

type CurrencyPairFees struct {
	SellFee string `json:"sell_fee"`
	BuyFee  string `json:"buy_fee"`
}
