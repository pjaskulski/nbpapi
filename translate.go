// simple translation system

package nbpapi

type lTexts map[string]string

// variable declaration for strings translation, the value
// of the variable is determined at program startup by the value of --lang flag
var l lTexts

// Get - method of Texts type, return localized string for key,
// if no key is found, return key
func (t *lTexts) Get(key string) string {
	result := (*t)[key]
	if result != "" {
		return result
	}
	return key
}

var langTexts = map[string]lTexts{
	"pl": {
		"Table type:":                                                       "Typ tabeli:",
		"Table number:":                                                     "Numer tabeli:",
		"Publication date:":                                                 "Data publikacji:",
		"CODE \t NAME \t AVERAGE (PLN)":                                     "KOD \t NAZWA \t ŚREDNI (PLN)",
		"---- \t ---- \t -------------":                                     "--- \t ----- \t ------------",
		"Trading date:":                                                     "Data notowania:",
		"CODE \t NAME \t BUY (PLN) \t SELL (PLN) ":                          "KOD \t NAZWA \t KUPNO (PLN) \t SPRZEDAŻ (PLN) ",
		"---- \t ---- \t --------- \t ---------- ":                          "--- \t ----- \t ----------- \t -------------- ",
		"TABLE,CODE,NAME,AVERAGE (PLN)":                                     "TABELA,KOD,NAZWA,ŚREDNI (PLN)",
		"TABLE,CODE,NAME,BUY (PLN),SELL (PLN)":                              "TABELA,KOD,NAZWA,KUPNO (PLN),SPRZEDAŻ (PLN)",
		"No --output parameter value, output format must be specified":      "Brak wartości parametru --output, należy podać format danych wyjściowych",
		"Invalid --output parameter value, allowed: table, json, csv, xml":  "Nieprawidłowa wartość parametru --output, dozwolone: table, json, csv, xml",
		"Value of one of the parameters should be given: --date or --last":  "Należy podać wartość jednego z parametrów: --date lub --last",
		"Invalid --last parameter value, allowed value > 0":                 "Nieprawidłowa wartość parametru --last, dozwolona wartość > 0",
		"Only one of the parameters must be given: either --date or --last": "Należy podać wartość tylko jednego z parametrów: albo --date albo --last",
		"Invalid --date parameter value, allowed values: 'today', 'current', 'YYYY-MM-DD' or 'YYYY-MM-DD: YYYY-MM-DD'": "Nieprawidłowa wartość parametru --date, dozwolone wartości: 'today', 'current', 'RRRR-MM-DD' lub 'RRRR-MM-DD:RRRR-MM-DD'",
		"The --table parameter value is missing, the type of the exchange table should be specified":                   "Brak wartości parametru --table, należy podać typ tabeli kursów",
		"Invalid parameter --table value, allowed values: A, B or C":                                                   "Nieprawidłowa wartość parametru --table, dozwolone wartości: A, B lub C",
		"No value of parameter --code, currency code should be given":                                                  "Brak wartości parametru --code, należy podać kod waluty",
		"No value of parameter --table, please specify type of exchange rate table":                                    "Brak wartości parametru --table, należy podać typ tabeli kursów",
		"Incorrect parameter value --table, allowed values: A, B or C":                                                 "Nieprawidłowa wartość parametru --table, dozwolone: A, B lub C",
		"Incorrect value of the --code parameter, ":                                                                    "Nieprawidłowa wartość parametru --code, ",
		"valid currency code from those available for Table A is allowed:":                                             "dozwolony poprawny kod waluty z dostępnych dla tabeli A: ",
		"valid currency code from those available for Table B is allowed:":                                             "dozwolony poprawny kod waluty z dostępnych dla tabeli B: ",
		"valid currency code from those available for Table C is allowed:":                                             "dozwolony poprawny kod waluty z dostępnych dla tabeli C: ",
		"The price of 1g of gold (of 1000 millesimal fineness)":                                                        "Cena złota (1 g złota w próbie 1000)",
		"DATE \t PRICE (PLN)":                       "DATA \t CENA (PLN)",
		"---- \t ---------- ":                       "---- \t ---------- ",
		"DATE,PRICE (PLN)":                          "DATA,CENA (PLN)",
		"Currency name:":                            "Nazwa waluty:",
		"Currency code:":                            "Kod waluty:",
		"TABLE \t DATE \t AVERAGE (PLN)":            "TABELA \t DATA \t ŚREDNI (PLN)",
		"----- \t ---- \t -------------":            "------ \t ---- \t ------------",
		"TABLE \t DATE \t BUY (PLN) \t SELL (PLN) ": "TABELA \t DATA \t KUPNO (PLN) \t SPRZEDAŻ (PLN) ",
		"----- \t ---- \t --------- \t ---------- ": "------ \t ---- \t ----------- \t -------------- ",
		"TABLE,DATE,AVERAGE (PLN)":                  "TABELA,DATA,ŚREDNI (PLN)",
		"TABLE,DATE,BUY (PLN),SELL (PLN)":           "TABELA,DATA,KUPNO (PLN),SPRZEDAŻ (PLN)",
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
		"No --output parameter value, output format must be specified":                                                 "No --output parameter value, output format must be specified",
		"Invalid --output parameter value, allowed: table, json, csv":                                                  "Invalid --output parameter value, allowed: table, json, csv",
		"Value of one of the parameters should be given: --date or --last":                                             "Value of one of the parameters should be given: --date or --last",
		"Invalid --last parameter value, allowed value > 0":                                                            "Invalid --last parameter value, allowed value > 0",
		"Only one of the parameters must be given: either --date or --last":                                            "Only one of the parameters must be given: either --date or --last",
		"Invalid --date parameter value, allowed values: 'today', 'current', 'YYYY-MM-DD' or 'YYYY-MM-DD: YYYY-MM-DD'": "Invalid --date parameter value, allowed values: 'today', 'current', 'YYYY-MM-DD' or 'YYYY-MM-DD: YYYY-MM-DD'",
		"The --table parameter value is missing, the type of the exchange table should be specified":                   "The --table parameter value is missing, the type of the exchange table should be specified",
		"Invalid parameter --table value, allowed values: A, B or C":                                                   "Invalid parameter --table value, allowed values: A, B or C",
		"No value of parameter --code, currency code should be given":                                                  "No value of parameter --code, currency code should be given",
		"No value of parameter --table, please specify type of exchange rate table":                                    "No value of parameter --table, please specify type of exchange rate table",
		"Incorrect parameter value --table, allowed values: A, B or C":                                                 "Incorrect parameter value --table, allowed values: A, B or C",
		"Incorrect value of the --code parameter, ":                                                                    "Incorrect value of the --code parameter, ",
		"valid currency code from those available for Table A is allowed:":                                             "valid currency code from those available for Table A is allowed:",
		"valid currency code from those available for Table B is allowed:":                                             "valid currency code from those available for Table B is allowed:",
		"valid currency code from those available for Table C is allowed:":                                             "valid currency code from those available for Table C is allowed:",
		"The price of 1g of gold (of 1000 millesimal fineness)":                                                        "The price of 1g of gold (of 1000 millesimal fineness)",
		"DATE \t PRICE (PLN)":                       "DATE \t PRICE (PLN)",
		"---- \t ---------- ":                       "---- \t ----------- ",
		"DATE,PRICE (PLN)":                          "DATE,PRICE (PLN)",
		"Currency name:":                            "Currency name:",
		"Currency code:":                            "Currency code:",
		"TABLE \t DATE \t AVERAGE (PLN)":            "TABLE \t DATE \t AVERAGE (PLN)",
		"----- \t ---- \t -------------":            "----- \t ---- \t -------------",
		"TABLE \t DATE \t BUY (PLN) \t SELL (PLN) ": "TABLE \t DATE \t BUY (PLN) \t SELL (PLN) ",
		"----- \t ---- \t --------- \t ---------- ": "----- \t ---- \t --------- \t ---------- ",
		"TABLE,DATE,AVERAGE (PLN)":                  "TABLE,DATE,AVERAGE (PLN)",
		"TABLE,DATE,BUY (PLN),SELL (PLN)":           "TABLE,DATE,BUY,SELL (PLN)",
	},
}
