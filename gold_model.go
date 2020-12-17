package nbpapi

import (
	"net/http"
)

// GoldRate type
type GoldRate struct {
	Data  string  `json:"data"`
	Price float64 `json:"cena"`
}

// NBPGold type
type NBPGold struct {
	GoldRates []GoldRate
	result    []byte
	Client    *http.Client
}
