package nbpapi

import "net/http"

// NBPTable type
type NBPTable struct {
	tableType string
	result    []byte
	Exchange  []ExchangeTable
	ExchangeC []ExchangeTableC
	Client    *http.Client
}

type rateTable struct {
	Currency string  `json:"currency"`
	Code     string  `json:"code"`
	Mid      float64 `json:"mid"`
}

// ExchangeTable type
type ExchangeTable struct {
	Table         string      `json:"table"`
	No            string      `json:"no"`
	EffectiveDate string      `json:"effectiveDate"`
	Rates         []rateTable `json:"rates"`
}

type rateTableC struct {
	Currency string  `json:"currency"`
	Code     string  `json:"code"`
	Bid      float64 `json:"bid"`
	Ask      float64 `json:"ask"`
}

// ExchangeTableC type
type ExchangeTableC struct {
	Table         string       `json:"table"`
	No            string       `json:"no"`
	TradingDate   string       `json:"tradingDate"`
	EffectiveDate string       `json:"effectiveDate"`
	Rates         []rateTableC `json:"rates"`
}
