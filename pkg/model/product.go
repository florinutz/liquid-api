package model

type Product struct {
	ID                      int      `json:"id,string"`
	ProductType             string   `json:"product_type,omitempty"`
	Code                    string   `json:"code,omitempty"`
	Name                    string   `json:"name,omitempty"`
	MarketAsk               Price    `json:"market_ask,string"`
	MarketBid               Price    `json:"market_bid,string"`
	Indicator               int      `json:"indicator,omitempty"`
	Currency                Currency `json:"currency,omitempty"`
	CurrencyPairCode        string   `json:"currency_pair_code,omitempty"`
	Symbol                  string   `json:"symbol,omitempty"`
	BtcMinimumWithdraw      string   `json:"btc_minimum_withdraw,omitempty"`
	FiatMinimumWithdraw     string   `json:"fiat_minimum_withdraw,omitempty"`
	PusherChannel           string   `json:"pusher_channel,omitempty"`
	TakerFee                Price    `json:"taker_fee,string"`
	MakerFee                Price    `json:"maker_fee,string"`
	LowMarketBid            Price    `json:"low_market_bid,string"`
	HighMarketAsk           Price    `json:"high_market_ask,string"`
	Volume24h               Price    `json:"volume_24h,string"`
	LastPrice24h            Price    `json:"last_price_24h,string"`
	LastTradedPrice         Price    `json:"last_traded_price,string"`
	LastTradedQuantity      Price    `json:"last_traded_quantity,string"`
	AveragePrice            Price    `json:"average_price,string"`
	QuotedCurrency          Currency `json:"quoted_currency,omitempty"`
	BaseCurrency            Currency `json:"base_currency,omitempty"`
	TickSize                string   `json:"tick_size,omitempty"`
	Disabled                bool     `json:"disabled,omitempty"`
	MarginEnabled           bool     `json:"margin_enabled,omitempty"`
	CfdEnabled              bool     `json:"cfd_enabled,omitempty"`
	PerpetualEnabled        bool     `json:"perpetual_enabled,omitempty"`
	LastEventTimestamp      string   `json:"last_event_timestamp,omitempty"`
	Timestamp               string   `json:"timestamp,omitempty"`
	MultiplierUp            string   `json:"multiplier_up,omitempty"`
	MultiplierDown          string   `json:"multiplier_down,omitempty"`
	AverageTimeInterval     int      `json:"average_time_interval,omitempty"`
	ProgressiveTierEligible bool     `json:"progressive_tier_eligible,omitempty"`
}
