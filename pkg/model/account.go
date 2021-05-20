package model

import "math/big"

type Account struct {
	ID                       int       `json:"id"`
	Currency                 Currency  `json:"currency"`
	CurrencySymbol           string    `json:"currency_symbol"`
	Balance                  float64   `json:"balance"`
	PusherChannel            string    `json:"pusher_channel"`
	LowestOfferInterestRate  float64   `json:"lowest_offer_interest_rate"`
	HighestOfferInterestRate float64   `json:"highest_offer_interest_rate"`
	SendToBtcAddress         string    `json:"send_to_btc_address"`
	CurrencyType             string    `json:"currency_type"`
	ExchangeRate             big.Float `json:"exchange_rate,string"`
}
