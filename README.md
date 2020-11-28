# nbpapi

[![GitHub](https://img.shields.io/github/license/pjaskulski/nbpapi)](https://opensource.org/licenses/MIT) 
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/pjaskulski/nbpapi?include_prereleases) 
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pjaskulski/nbpapi) 
[![go report](https://goreportcard.com/badge/github.com/pjaskulski/kursnbp)](https://goreportcard.com/report/github.com/pjaskulski/nbpapi) 


Go library for NBP (National Bank of Poland) api 

Library used in the kursNBP project: [https://github.com/pjaskulski/kursnbp](https://github.com/pjaskulski/kursnbp)


To install and use:

```
go get https://github.com/pjaskulski/nbpapi
```

To clone and invoke the tests:

```bash
git clone https://github.com/pjaskulski/nbpapi

cd nbpapi

make test

make cover
```

## Examples:
    
```go
// How to get table A of currency exchange rates published
// on 12 Nov 2020: function TableByDate("2020-11-12") fetch data, then
// one can get data from nbpTable.Exchange structures (or
// nbpTable.ExchangeC structures, depending of the type of
// table of exchange rates)
var tableNo string

nbpTable := nbpapi.NewTable("A")
err: = nbpTable.TableByDate("2020-11-12")
if err != nil {
	log.Fatal(err)
}

for _, item := range nbpTable.Exchange {
	tableNo = item.No
	for _, currencyItem := range item.Rates {
		fmt.Println(tableNo, currencyItem.Code, currencyItem.Currency, currencyItem.Mid)
	}
}
```

Output:

    221/A/NBP/2020 THB bat (Tajlandia) 0.1256
    221/A/NBP/2020 USD dolar amerykański 3.7995
    221/A/NBP/2020 AUD dolar australijski 2.7636
    221/A/NBP/2020 HKD dolar Hongkongu 0.4901
    221/A/NBP/2020 CAD dolar kanadyjski 2.9074
    221/A/NBP/2020 NZD dolar nowozelandzki 2.6119
    221/A/NBP/2020 SGD dolar singapurski 2.8171
    221/A/NBP/2020 EUR euro 4.4868
    221/A/NBP/2020 HUF forint (Węgry) 0.012648
    221/A/NBP/2020 CHF frank szwajcarski 4.1573
    [...]
 
    
```go
// today's gold price: function GetPriceToday returns GoldPrice struct
var price nbpapi.GoldRate
var err error

gold := nbpapi.NewGold()

price, err = gold.GetPriceToday()
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Publication date: ", price.Data)
	fmt.Println("Price of 1g of gold (PLN): ", price.Price)
}
```

Output: 

    Publication date:  2020-11-27
    Price of 1g of gold (PLN):  218.41

More examples in the \ example folder.


## Documentation:

```
package nbpapi // import "github.com/pjaskulski/nbpapi"


FUNCTIONS

func setLang(lang string)
    setLang function (language for output functions)


TYPES

type ExchangeTable struct {
	Table         string      `json:"table"`
	No            string      `json:"no"`
	EffectiveDate string      `json:"effectiveDate"`
	Rates         []rateTable `json:"rates"`
}
    ExchangeTable type

type ExchangeTableC struct {
	Table         string       `json:"table"`
	No            string       `json:"no"`
	TradingDate   string       `json:"tradingDate"`
	EffectiveDate string       `json:"effectiveDate"`
	Rates         []rateTableC `json:"rates"`
}
    ExchangeTableC type

type GoldRate struct {
	Data  string  `json:"data"`
	Price float64 `json:"cena"`
}
    GoldRate type

type NBPCurrency struct {
	Exchange  exchangeCurrency
	ExchangeC exchangeCurrencyC
	// Has unexported fields.
}
    NBPCurrency type

func NewCurrency(tFlag string) *NBPCurrency
    NewCurrency - function creates new currency type

func (c *NBPCurrency) CurrencyByDate(dFlag string, cFlag string) error
    CurrencyByDate - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field

func (c *NBPCurrency) CurrencyLast(cFlag string, lFlag int) error
    CurrencyLast - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field

func (c *NBPCurrency) CurrencyRaw(dFlag string, lFlag int, cFlag string, repFormat string) error
    CurrencyRaw - function downloads data in json or xml form

func (c *NBPCurrency) CurrencyToday(cFlag string) error
    CurrencyToday - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field

func (c *NBPCurrency) GetCSVOutput(lang string) string
    GetCSVOutput - function returns currency rates, in the form of CSV (data
    separated by a comma), depending on the tableType field: for type A and B
    tables a column with an average rate is printed, for type C two columns: buy
    price and sell price

func (c *NBPCurrency) GetPrettyOutput(lang string) string
    GetPrettyOutput - function returns exchange rates as formatted table
    depending on the tableType field: for type A and B tables a column with an
    average rate is printed, for type C two columns: buy price and sell price

func (c *NBPCurrency) GetRateByDate(code string, date string) ([]Rate, error)
    GetRateByDate - function downloads today's currency exchange rate and
    returns Rate struct (or error)

func (c *NBPCurrency) GetRateCurrent(cFlag string) (Rate, error)
    GetRateCurrent - function downloads current currency exchange rate and
    return Rate struct (or error)

func (c *NBPCurrency) GetRateToday(cFlag string) (Rate, error)
    GetRateToday - function downloads today's currency exchange rate and returns
    Rate struct (or error)

func (c *NBPCurrency) GetRawOutput() string
    GetRawOutput - function print just result of request (json or xml)

type NBPGold struct {
	GoldRates []GoldRate
	// Has unexported fields.
}
    NBPGold type

func NewGold() *NBPGold
    NewGold - function creates new gold type

func (g *NBPGold) GetCSVOutput(lang string) string
    GetCSVOutput - function returns prices of gold in CSV format (comma
    separated data)

func (g *NBPGold) GetPrettyOutput(lang string) string
    GetPrettyOutput - function returns a formatted table of gold prices

func (g *NBPGold) GetPriceByDate(date string) ([]GoldRate, error)
    GetPriceByDate - function returns gold prices (as slice od struct), by date
    ("YYYY-MM-DD") or range of dates ("YYYY-MM-DD:YYYY-MM-DD")

func (g *NBPGold) GetPriceCurrent() (GoldRate, error)
    GetPriceCurrent - function downloads and returns gold price,

func (g *NBPGold) GetPriceToday() (GoldRate, error)
    GetPriceToday - function downloads and returns gold price,

func (g *NBPGold) GetRawOutput() string
    GetRawOutput - function returns just result of request (json or xml)

func (g *NBPGold) GoldByDate(dFlag string) error
    GoldByDate - function downloads and writes data to goldRates slice, raw data
    (json) still available in result field

func (g *NBPGold) GoldLast(lFlag int) error
    GoldLast - function downloads and writes data to goldRates slice, raw data
    (json) still available in result field

func (g *NBPGold) GoldRaw(dFlag string, lFlag int, repFormat string) error
    GoldRaw - function downloads data in json or xml form

type NBPTable struct {
	Exchange  []ExchangeTable
	ExchangeC []ExchangeTableC
	// Has unexported fields.
}
    NBPTable type

func NewTable(tFlag string) *NBPTable
    NewTable - function creates new table type

func (t *NBPTable) GetCSVOutput(lang string) string
    GetCSVOutput - function prints tables of exchange rates in the console, in
    the form of CSV (data separated by a comma), depending on the tableType
    field: for type A and B tables a column with an average rate is printed, for
    type C two columns: buy price and sell price

func (t *NBPTable) GetPrettyOutput(lang string) string
    GetPrettyOutput - function returns tables of exchange rates as formatted
    table, depending on the tableType field: for type A and B tables a column
    with an average rate is printed, for type C two columns: buy price and sell
    price

func (t *NBPTable) GetRawOutput() string
    GetRawOutput - function returns just result of request (json or xml)

func (t *NBPTable) GetTableCCurrent() ([]ExchangeTableC, error)
    GetTableCCurrent - function downloads current table of currency exchange
    rates and return slice of struct (or error), version for table C (bid, ask)

func (t *NBPTable) GetTableCurrent() ([]ExchangeTable, error)
    GetTableCurrent - function downloads current table of currency exchange
    rates and return slice of struct (or error), version for table A, B (mid)

func (t *NBPTable) TableByDate(dFlag string) error
    TableByDate - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field

func (t *NBPTable) TableLast(lFlag int) error
    TableLast - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field

func (t *NBPTable) TableRaw(dFlag string, lFlag int, repFormat string) error
    TableRaw - function downloads data in json or xml form

type Rate struct {
	No            string
	EffectiveDate string
	Mid           float64
	Bid           float64
	Ask           float64
}
    Rate type

```

## TODO

- more tests