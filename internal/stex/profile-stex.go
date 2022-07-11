package stex

import "errors"

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

func (c *StexClient) GetProfileInfo(ch chan<- interface{}) error {

	resource := "info"

	req, err := c.CreateGetRequest(c.Endpoints.Profile, resource, "", "")
	if err != nil {
		return errors.New("faild create get request for stex profile info")
	}

	c.Authenticate(req)

	res := InfoData{}
	if err = c.SendRequest(req, &res); err != nil {
		return errors.New("failed get request for stex profile info")
	}

	ch <- res

	return nil
}
