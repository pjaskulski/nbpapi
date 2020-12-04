package nbpapi

import "net/http"

// NBPCurrency type
type NBPCurrency struct {
	tableType string
	result    []byte
	Exchange  exchangeCurrency
	ExchangeC exchangeCurrencyC
	Client    *http.Client
}

type rateCurrency struct {
	No            string  `json:"no"`
	EffectiveDate string  `json:"effectiveDate"`
	Mid           float64 `json:"mid"`
}

type exchangeCurrency struct {
	Table    string         `json:"table"`
	Currency string         `json:"currency"`
	Code     string         `json:"code"`
	Rates    []rateCurrency `json:"rates"`
}

type rateCurrencyC struct {
	No            string  `json:"no"`
	EffectiveDate string  `json:"effectiveDate"`
	Bid           float64 `json:"bid"`
	Ask           float64 `json:"ask"`
}

type exchangeCurrencyC struct {
	Table    string          `json:"table"`
	Currency string          `json:"currency"`
	Code     string          `json:"code"`
	Rates    []rateCurrencyC `json:"rates"`
}

// Rate type
type Rate struct {
	No            string
	EffectiveDate string
	Mid           float64
	Bid           float64
	Ask           float64
}
