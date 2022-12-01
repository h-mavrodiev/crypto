package stex

type Balance struct {
	BTC  string `json:"btc"`
	ETH  string `json:"eth"`
	USDT string `json:"usdt"`
}

type ProfileBalance struct {
	Success bool `json:"success"`
	Data    data `json:"data"`
}

type data []struct {
	ID                            int                      `json:"id"`
	CurrencyID                    int                      `json:"currency_id"`
	CurrencyCode                  string                   `json:"currency_code"`
	CurrencyName                  string                   `json:"currency_name"`
	Balance                       string                   `json:"balance"`
	FrozenBalance                 string                   `json:"frozen_balance"`
	BonusBalance                  string                   `json:"bonus_balance"`
	HoldBalance                   string                   `json:"hold_balance"`
	TotalBalance                  string                   `json:"total_balance"`
	DisableDeposits               bool                     `json:"disable_deposits"`
	DisableWithdrawals            bool                     `json:"disable_withdrawals"`
	WithdrawalLimit               string                   `json:"withdrawal_limit"`
	Delisted                      bool                     `json:"delisted"`
	Disabled                      bool                     `json:"disabled"`
	DepositAddress                depositAddress           `json:"deposit_address"`
	MultiDepositAddresses         multiDepositAddresses    `json:"multi_deposit_addresses"`
	WithdrawalAdditionalFieldName string                   `json:"withdrawal_additional_field_name"`
	CurrencyTypeID                int                      `json:"currency_type_id"`
	ProtocolSpecificSettings      protocolSpecificSettings `json:"protocol_specific_settings"`
	CoinInfo                      coinInfo                 `json:"coin_info"`
	Rates                         rates                    `json:"rates"`
}

type depositAddress struct {
	Address                        string `json:"address"`
	AddressName                    string `json:"address_name"`
	AdditionalAddressParameter     string `json:"additional_address_parameter"`
	AdditionalAddressParameterName string `json:"additional_address_parameter_name"`
	Notification                   string `json:"notification"`
	ProtocolID                     int    `json:"protocol_id"`
	ProtocolName                   string `json:"protocol_name"`
	SupportsNewAddressCreation     bool   `json:"supports_new_address_creation"`
}

type multiDepositAddresses struct {
	Address                        string `json:"address"`
	AddressName                    string `json:"address_name"`
	AdditionalAddressParameter     string `json:"additional_address_parameter"`
	AdditionalAddressParameterName string `json:"additional_address_parameter_name"`
	Notification                   string `json:"notification"`
	ProtocolID                     int    `json:"protocol_id"`
	ProtocolName                   string `json:"protocol_name"`
	SupportsNewAddressCreation     bool   `json:"supports_new_address_creation"`
}

type protocolSpecificSettings []struct {
	ProtocolName                  string  `json:"protocol_name"`
	ProtocolID                    int     `json:"protocol_id"`
	Active                        bool    `json:"active"`
	DisableDeposits               bool    `json:"disable_deposits"`
	DisableWithdrawals            bool    `json:"disable_withdrawals"`
	WithdrawalLimit               int     `json:"withdrawal_limit"`
	DepositFeeCurrencyID          int     `json:"deposit_fee_currency_id"`
	DepositFeeCurrencyCode        string  `json:"deposit_fee_currency_code"`
	DepositFeePercent             int     `json:"deposit_fee_percent"`
	DepositFeeConst               int     `json:"deposit_fee_const"`
	WithdrawalFeeCurrencyID       int     `json:"withdrawal_fee_currency_id"`
	WithdrawalFeeCurrencyCode     string  `json:"withdrawal_fee_currency_code"`
	WithdrawalFeeConst            float64 `json:"withdrawal_fee_const"`
	WithdrawalFeePercent          int     `json:"withdrawal_fee_percent"`
	BlockExplorerURL              string  `json:"block_explorer_url"`
	WithdrawalAdditionalFieldName string  `json:"withdrawal_additional_field_name"`
}

type coinInfo struct {
	Twitter               string `json:"twitter"`
	Version               string `json:"version"`
	Facebook              string `json:"facebook"`
	Telegram              string `json:"telegram"`
	IconLarge             string `json:"icon_large"`
	IconSmall             string `json:"icon_small"`
	Description           string `json:"description"`
	OfficialSite          string `json:"official_site"`
	OfficialBlockExplorer string `json:"official_block_explorer"`
}

type rates struct {
	Btc float64 `json:"BTC"`
}
