// simple translation system

package nbpapi

type lTexts map[string]string

// variable declaration for strings translation
var l lTexts

// Get - method of lTexts type, return localized string for key,
// if no key is found, return key - english text
func (t *lTexts) Get(key string) string {
	result := (*t)[key]
	if result != "" {
		return result
	}
	return key
}

var langTexts = map[string]lTexts{
	"pl": {
		"Table type:":                                           "Typ tabeli:",
		"Table number:":                                         "Numer tabeli:",
		"Publication date:":                                     "Data publikacji:",
		"CODE \t NAME \t AVERAGE (PLN)":                         "KOD \t NAZWA \t ŚREDNI (PLN)",
		"---- \t ---- \t -------------":                         "--- \t ----- \t ------------",
		"Trading date:":                                         "Data notowania:",
		"CODE \t NAME \t BUY (PLN) \t SELL (PLN) ":              "KOD \t NAZWA \t KUPNO (PLN) \t SPRZEDAŻ (PLN) ",
		"---- \t ---- \t --------- \t ---------- ":              "--- \t ----- \t ----------- \t -------------- ",
		"TABLE,CODE,NAME,AVERAGE (PLN)":                         "TABELA,KOD,NAZWA,ŚREDNI (PLN)",
		"TABLE,CODE,NAME,BUY (PLN),SELL (PLN)":                  "TABELA,KOD,NAZWA,KUPNO (PLN),SPRZEDAŻ (PLN)",
		"The price of 1g of gold (of 1000 millesimal fineness)": "Cena złota (1 g złota w próbie 1000)",
		"DATE \t PRICE (PLN)":                                   "DATA \t CENA (PLN)",
		"---- \t ---------- ":                                   "---- \t ---------- ",
		"DATE,PRICE (PLN)":                                      "DATA,CENA (PLN)",
		"Currency name:":                                        "Nazwa waluty:",
		"Currency code:":                                        "Kod waluty:",
		"TABLE \t DATE \t AVERAGE (PLN)":                        "TABELA \t DATA \t ŚREDNI (PLN)",
		"----- \t ---- \t -------------":                        "------ \t ---- \t ------------",
		"TABLE \t DATE \t BUY (PLN) \t SELL (PLN) ":             "TABELA \t DATA \t KUPNO (PLN) \t SPRZEDAŻ (PLN) ",
		"----- \t ---- \t --------- \t ---------- ":             "------ \t ---- \t ----------- \t -------------- ",
		"TABLE,DATE,AVERAGE (PLN)":                              "TABELA,DATA,ŚREDNI (PLN)",
		"TABLE,DATE,BUY (PLN),SELL (PLN)":                       "TABELA,DATA,KUPNO (PLN),SPRZEDAŻ (PLN)",
		"The average of the last quotations is: ":               "Średnia z ostatnich notowań wynosi: ",
		"The average exchange rate (from downloaded): ":         "Średni kurs waluty (z pobranych): ",
		"Buy: the average exchange rate (from downloaded): ":    "Zakup: średni kurs waluty (z pobranych): ",
		"Sell: the average exchange rate (from downloaded): ":   "Sprzedaż: średni kurs waluty (z pobranych): ",
	},
	"en": {
		"Table type:":                  "Table type:",
		"Table number:":                "Table number:",
		"Publication date:":            "Publication date:",
		"CODE \t NAME \t AVERAGE":      "CODE \t NAME \t AVERAGE (PLN)",
		"---- \t ---- \t -------":      "---- \t ---- \t -------------",
		"Trading date:":                "Trading date:",
		"CODE \t NAME \t BUY \t SELL ": "CODE \t NAME \t BUY (PLN) \t SELL (PLN) ",
		"---- \t ---- \t --- \t ---- ": "---- \t ---- \t --------- \t ---------- ",
		"TABLE,CODE,NAME,AVERAGE":      "TABLE,CODE,NAME,AVERAGE (PLN)",
		"TABLE,CODE,NAME,BUY,SELL":     "TABLE,CODE,NAME,BUY (PLN),SELL (PLN)",
		"The price of 1g of gold (of 1000 millesimal fineness)": "The price of 1g of gold (of 1000 millesimal fineness)",
		"DATE \t PRICE (PLN)":                                 "DATE \t PRICE (PLN)",
		"---- \t ---------- ":                                 "---- \t ----------- ",
		"DATE,PRICE (PLN)":                                    "DATE,PRICE (PLN)",
		"Currency name:":                                      "Currency name:",
		"Currency code:":                                      "Currency code:",
		"TABLE \t DATE \t AVERAGE (PLN)":                      "TABLE \t DATE \t AVERAGE (PLN)",
		"----- \t ---- \t -------------":                      "----- \t ---- \t -------------",
		"TABLE \t DATE \t BUY (PLN) \t SELL (PLN) ":           "TABLE \t DATE \t BUY (PLN) \t SELL (PLN) ",
		"----- \t ---- \t --------- \t ---------- ":           "----- \t ---- \t --------- \t ---------- ",
		"TABLE,DATE,AVERAGE (PLN)":                            "TABLE,DATE,AVERAGE (PLN)",
		"TABLE,DATE,BUY (PLN),SELL (PLN)":                     "TABLE,DATE,BUY,SELL (PLN)",
		"The average of the last quotations is: ":             "The average of the last quotations is: ",
		"The average exchange rate (from downloaded): ":       "The average exchange rate (from downloaded): ",
		"Buy: the average exchange rate (from downloaded): ":  "Buy: the average exchange rate (from downloaded): ",
		"Sell: the average exchange rate (from downloaded): ": "Sell: the average exchange rate (from downloaded): ",
	},
}
