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
	HTTPClient    *http.Client
}

type StexInfo struct {
	Sells       float64
	SellsVolume float64
	Buys        float64
	BuysVolume  float64
}

type InfoData struct {
	Email                 string           `json:"email"`
	Username              string           `json:"username"`
	UserID                int              `json:"user_id"`
	Verifications         Verifications    `json:"verifications"`
	TradingFeeLevels      TradingFeeLevels `json:"trading_fee_levels"`
	APIWithdrawalsAllowed bool             `json:"api_withdrawals_allowed"`
	ReferralProgram       ReferalProgram   `json:"referral_program"`
	ApproxBalance         ApproxBalance    `json:"approx_balance"`
	Settings              Settings         `json:"settings"`
}

type Verifications struct {
	Cryptonomica bool `json:"cryptonomica"`
	Privatbank   bool `json:"privatbank"`
	Stex         bool `json:"stex"`
}

type TradingFeeLevels struct {
	NotVerified  string `json:"not_verified"`
	Cryptonomica string `json:"cryptonomica"`
	Privatbank   string `json:"privatbank"`
	Stex         string `json:"stex"`
}

type ReferalProgram struct {
	ReferralCode string `json:"referral_code"`
	Members      int    `json:"members"`
	Invited      bool   `json:"invited"`
}

type ApproxBalance struct {
	Usd Usd `json:"USD"`
	Btc Btc `json:"BTC"`
}

type Usd struct {
	Balance       string `json:"balance"`
	FrozenBalance string `json:"frozen_balance"`
	BonusBalance  string `json:"bonus_balance"`
	HoldBalance   string `json:"hold_balance"`
	TotalBalance  string `json:"total_balance"`
}

type Btc struct {
	Balance       string `json:"balance"`
	FrozenBalance string `json:"frozen_balance"`
	BonusBalance  string `json:"bonus_balance"`
	HoldBalance   string `json:"hold_balance"`
	TotalBalance  string `json:"total_balance"`
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

type OrderBookDetails struct {
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
