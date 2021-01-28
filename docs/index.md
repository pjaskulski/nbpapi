# nbpapi

[![GitHub](https://img.shields.io/github/license/pjaskulski/nbpapi)](https://opensource.org/licenses/MIT) 
![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/pjaskulski/nbpapi?include_prereleases) 
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pjaskulski/nbpapi) 
[![go report](https://goreportcard.com/badge/github.com/pjaskulski/kursnbp)](https://goreportcard.com/report/github.com/pjaskulski/nbpapi) 


Go client library (unofficial) for NBP (National Bank of Poland) API: [http://api.nbp.pl/en.html](http://api.nbp.pl/en.html). 

The module's functions allow you to download currency rates and prices of gold from the NBP API service. Three types have been prepared, corresponding to the basic information provided by the NBP. The **NBPTable** type is used to download full currency exchange rate tables. Three types of tables A, B, and C are supported, tables A and B contain average currency prices, table C contains 'buy' and 'sell' prices. The TableRaw, TableByDate, TableLast methods allow to download the exchange rate table (or multiple tables) and save the result in internal structures of the NBPTable type. The GetTableCurrent, GetTableToday, etc. methods return the downloaded data in the form of slices of structs. Create[]...] methods return strings prepared for printing or saving to file, based on data downloaded by Table[...] methods. 
Other types work the same way: **NBPCurrency** - for downloading the rates of a particular currency, and **NBPGold** for downloading gold prices.

**Disclaimer**: nbpapi client library is a private project, in no way related to the National Bank of Poland, except for the use of the NBP public API.

## Module usage

The library is used in the kursNBP project: [https://github.com/pjaskulski/kursnbp](https://github.com/pjaskulski/kursnbp).

## Cache

Optionally you can store the results of the query in the memory of your computer. If a cache is enabled, the results are stored for 60 minutes by default, each time you call an identical query again API service it will return the results from the cache instead of querying the real NBP API server.
Function EnableCache turns on the caching mechanism, DisableCache turns it off (deleting the memory content).


To install and use:

```
go get https://github.com/pjaskulski/nbpapi
```

To clone the repo and invoke the tests:

```bash
git clone https://github.com/pjaskulski/nbpapi

cd nbpapi

make test

make cover
```

Warning: by default some tests use (great) [httpmock](https://github.com/jarcoal/httpmock) library, to use the real NBP API service it is necessary to set the `USEMOCK` environment variable to `'0'`

## Examples:
    
```go
// How to get table A of currency exchange rates published
// on 12 Nov 2020: function TableByDate("2020-11-12") fetch data, then
// one can get data from nbpTable.Exchange structures (or
// nbpTable.ExchangeC structures, depending of the type of
// table of exchange rates)
var tableNo string

client := nbpapi.NewTable("A")
err: = client.TableByDate("2020-11-12")
if err != nil {
	log.Fatal(err)
}

for _, item := range client.Exchange {
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

gold := nbpapi.NewGold()

price, err := gold.GetPriceToday()
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

0. Global functions:
    * EnableCache
    * DisableCache
1. **NBPTable** type - complete tables of currency exchange rates:
    * create new: `client := NewTable("C")`
    * methods that downloading data into NBPTable type structures:
      - TableRaw _(only raw data: json or xml)_ 
      - TableByDate _(raw data and and structures filled with data)_
      - TableLast _(as above)_
    * methods that return structures with tables of currency exchange rates:
      - GetTableCurrent 
      - GetTableCCurrent 
      - GetTableToday 
      - GetTableCToday 
      - GetTableByDay 
      - GetTableCByDay
    * methods that return data (downloaded previously by TableByDate, or TableLast methods) as text, ready to be printed or saved to a file:
	  - CreatePrettyOutput 
      - CreateCSVOutput
      - CreateRawOutput
    * other methods:
      - SetTableType
2. **NBPCurrency** type - particular currency exchange rates:
    * create new: `client := NewCurrency("A")`
    * methods that downloading data into NBPCurrency type structures:
      - CurrencyRaw _(only raw data: json or xml)_
      - CurrencyByDate _(raw data and and structures filled with data)_
      - CurrencyLast _(as above)_
      - CurrencyToday _(as above)_
    * methods that return structures with currency rates:
      - GetRateCurrent 
      - GetRateToday
      - GetRateByDate
    * methods that return data (downloaded previously) as text, ready to be printed or saved to a file:
      - CreatePrettyOutput 
      - CreateCSVOutput 
      - CreateRawOutput
    * other methods:
      - SetTableType
3. **NBPGold** type - gold prices:
    * create new: `client := NewGold()`
    * methods that downloading data into NBPGold type structures:
      - GoldRaw _(only raw data: json or xml)_
      - GoldByDate _(raw data and and structures filled with data)_
      - GoldLast _(as above)_
    * methods that return structures with prices of gold:
      - GetPriceToday
      - GetPriceCurrent 
      - GetPriceByDate
    * methods that return data (downloaded previously) as text, ready to be printed or saved to a file:
      - CreatePrettyOutput 
      - CreateCSVOutput 
      - CreateRawOutput


Detailed documentation:    

```
package nbpapi // import "github.com/pjaskulski/nbpapi"


VARIABLES

var (
	ErrInvalidCurrencyCode error = errors.New("Invalid currency code")
	ErrInvalidTableType    error = errors.New("Invalid table type, allowed values: A, B or C")
)
    errors for invalid code, table type

var CacheOn bool = false
    CacheOn - is in-memory cache on/off

var CurrencyValuesA = []string{"THB", "USD", "AUD", "HKD", "CAD", "NZD", "SGD", "EUR", "HUF", "CHF",
	"GBP", "UAH", "JPY", "CZK", "DKK", "ISK", "NOK", "SEK", "HRK", "RON",
	"BGN", "TRY", "ILS", "CLP", "PHP", "MXN", "ZAR", "BRL", "MYR", "RUB",
	"IDR", "INR", "KRW", "CNY", "XDR"}
    CurrencyValuesA - list of supported currencies for table type A

var CurrencyValuesB = []string{"MGA", "PAB", "ETB", "AFN", "VES", "BOB", "CRC", "SVC", "NIO", "GMD",
	"MKD", "DZD", "BHD", "IQD", "JOD", "KWD", "LYD", "RSD", "TND", "MAD",
	"AED", "STN", "BSD", "BBD", "BZD", "BND", "FJD", "GYD", "JMD", "LRD",
	"NAD", "SRD", "TTD", "XCD", "SBD", "ZWL", "VND", "AMD", "CVE", "AWG",
	"BIF", "XOF", "XAF", "XPF", "DJF", "GNF", "KMF", "CDF", "RWF", "EGP",
	"GIP", "LBP", "SSP", "SDG", "SYP", "GHS", "HTG", "PYG", "ANG", "PGK",
	"LAK", "MWK", "ZMW", "AOA", "MMK", "GEL", "MDL", "ALL", "HNL", "SLL",
	"SZL", "LSL", "AZN", "MZN", "NGN", "ERN", "TWD", "TMT", "MRU", "TOP",
	"MOP", "ARS", "DOP", "COP", "CUP", "UYU", "BWP", "GTQ", "IRR", "YER",
	"QAR", "OMR", "SAR", "KHR", "BYN", "LKR", "MVR", "MUR", "NPR", "PKR",
	"SCR", "PEN", "KGS", "TJS", "UZS", "KES", "SOS", "TZS", "UGX", "BDT",
	"WST", "KZT", "MNT", "VUV", "BAM"}
    CurrencyValuesB - list of supported currencies for table type B

var CurrencyValuesC = []string{"USD", "AUD", "CAD", "EUR", "HUF", "CHF", "GBP", "JPY", "CZK", "DKK", "NOK",
	"SEK", "XDR"}
    CurrencyValuesC - list of supported currencies for table type C

var Memory *cache.Cache
    Memory - cache

var TableValues = []string{"A", "B", "C"}
    TableValues - list of table types


FUNCTIONS

func DisableCache()
    DisableCache func - turn off im-memory cache

func EnableCache()
    EnableCache func - turn on im-memory cache

func IsValidXML(input string) bool
    IsValidXML - func from:
    https://stackoverflow.com/questions/53476012/how-to-validate-a-xml


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
	Client    *http.Client
	// Has unexported fields.
}
    NBPCurrency type

func NewCurrency(tableType string) *NBPCurrency
    NewCurrency - function creates new currency type

func (c *NBPCurrency) CreateCSVOutput(lang string) string
    CreateCSVOutput - function returns currency rates, in the form of CSV (data
    separated by a comma), depending on the tableType field: for type A and B
    tables a column with an average rate is printed, for type C two columns: buy
    price and sell price

    Parameters:

        lang - 'en' or 'pl'

func (c *NBPCurrency) CreatePrettyOutput(lang string) string
    CreatePrettyOutput - function returns exchange rates as formatted table
    depending on the tableType field: for type A and B tables a column with an
    average rate is printed, for type C two columns: buy price and sell price

    Parameters:

        lang - 'en' or 'pl'

func (c *NBPCurrency) CreateRawOutput() string
    CreateRawOutput - function print just result of request (json or xml)

func (c *NBPCurrency) CurrencyByDate(code, date string) error
    CurrencyByDate - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field


    Function returns error or nil

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (rate for today) or 'current' - current table / rate (last published)

        code - ISO 4217 currency code, depending on the type of the
        table available currencies may vary

func (c *NBPCurrency) CurrencyLast(code string, last int) error
    CurrencyLast - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field


    Function returns error or nil

    Parameters:

        code - ISO 4217 currency code, depending on the type of the
        table available currencies may vary

        last - as an alternative to date, the last <n> tables/rates
        can be retrieved

func (c *NBPCurrency) CurrencyRaw(code, date string, last int, format string) error
    CurrencyRaw - function downloads data in json or xml form


    Function returns error or nil

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (rate for today) or 'current' - current table / rate (last published)

        last - as an alternative to date, the last <n> tables/rates
        can be retrieved

        code - ISO 4217 currency code, depending on the type of the
        table available currencies may vary

        format - 'json' or 'xml'

func (c *NBPCurrency) CurrencyToday(code string) error
    CurrencyToday - function downloads and writes data to exchange (exchangeC)
    slice, raw data (json) still available in result field


    Function returns error or nil

    Parameters:

        code - ISO 4217 currency code, depending on the type of the
        table available currencies may vary

func (c *NBPCurrency) GetRateByDate(code, date string) ([]Rate, error)
    GetRateByDate - function downloads today's currency exchange rate and
    returns slice of Rate struct (or error)

    Parameters:

        code - ISO 4217 currency code, depending on the type of the
        table available currencies may vary

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (rate for today) or 'current' - current table / rate (last published)

func (c *NBPCurrency) GetRateCurrent(code string) (Rate, error)
    GetRateCurrent - function downloads current currency exchange rate and
    return Rate struct (or error)

    Parameters:

        code - ISO 4217 currency code, depending on the type of the
        table available currencies may vary

func (c *NBPCurrency) GetRateToday(code string) (Rate, error)
    GetRateToday - function downloads today's currency exchange rate and returns
    Rate struct (or error)

    Parameters:

        code - ISO 4217 currency code, depending on the type of the
        table available currencies may vary

func (c *NBPCurrency) SetTableType(tableType string) error
    SetTableType - the function allows to set the supported type of exchange
    rate table

type NBPGold struct {
	GoldRates []GoldRate

	Client *http.Client
	// Has unexported fields.
}
    NBPGold type

func NewGold() *NBPGold
    NewGold - function creates new gold type

func (g *NBPGold) CreateCSVOutput(lang string) string
    CreateCSVOutput - function returns prices of gold in CSV format (comma
    separated data)

    Parameters:

        lang - 'en' or 'pl'

func (g *NBPGold) CreatePrettyOutput(lang string) string
    CreatePrettyOutput - function returns a formatted table of gold prices

    Parameters:

        lang - 'en' or 'pl'

func (g *NBPGold) CreateRawOutput() string
    CreateRawOutput - function returns just result of request (json or xml)

func (g *NBPGold) GetPriceByDate(date string) ([]GoldRate, error)
    GetPriceByDate - function returns gold prices (as slice of struct), by date
    ("YYYY-MM-DD") or range of dates ("YYYY-MM-DD:YYYY-MM-DD")

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (price for today) or 'current' - current gold price (last published)

func (g *NBPGold) GetPriceCurrent() (GoldRate, error)
    GetPriceCurrent - function downloads and returns current gold price as
    GoldRate struct

func (g *NBPGold) GetPriceToday() (GoldRate, error)
    GetPriceToday - function downloads and returns today's gold price, as
    GoldRate struct

func (g *NBPGold) GoldByDate(date string) error
    GoldByDate - function downloads and writes data to goldRates slice, raw data
    (json) still available in NBPGold.result field


    Function returns error or nil

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (price for today) or 'current' - current gold price (last published)

func (g *NBPGold) GoldLast(last int) error
    GoldLast - function downloads and writes data to GoldRates slice, raw data
    (json) still available in NBPGold.result field


    Function returns error or nil

    Parameters:

        last - as an alternative to date, the last <n> prices of gold
        can be retrieved

func (g *NBPGold) GoldRaw(date string, last int, format string) error
    GoldRaw - function downloads data in json or xml form


    Function returns error or nil

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (price for today) or 'current' - current gold price (last published)

        last - as an alternative to date, the last <n> prices of gold
        can be retrieved

        format - 'json' or 'xml'

type NBPTable struct {
	Exchange  []ExchangeTable
	ExchangeC []ExchangeTableC
	Client    *http.Client
	// Has unexported fields.
}
    NBPTable type

func NewTable(tableType string) *NBPTable
    NewTable - function creates new table type

func (t *NBPTable) CreateCSVOutput(lang string) string
    CreateCSVOutput - function prints tables of exchange rates in the console,
    in the form of CSV (data separated by a comma), depending on the tableType
    field: for type A and B tables a column with an average rate is printed, for
    type C two columns: buy price and sell price

    Parameters:

        lang - 'en' or 'pl'

func (t *NBPTable) CreatePrettyOutput(lang string) string
    CreatePrettyOutput - function returns tables of exchange rates as formatted
    table, depending on the tableType field: for type A and B tables a column
    with an average rate is printed, for type C two columns: buy price and sell
    price

    Parameters:

        lang - 'en' or 'pl'

func (t *NBPTable) CreateRawOutput() string
    CreateRawOutput - function returns just result of request (json or xml)

func (t *NBPTable) GetTableByDate(date string) ([]ExchangeTable, error)
    GetTableByDate - function retrieves a table of exchange rates for a given
    day and return slice of struct ExchangeTable (or error). Version for table
    A, B (mid - average price)

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (rate for today) or 'current' - current table / rate (last published)

func (t *NBPTable) GetTableCByDate(date string) ([]ExchangeTableC, error)
    GetTableCByDate - function retrieves a table of exchange rates for a given
    day and return slice of struct ExchangeTable (or error). Version for table C
    (ask, bid - buy, sell prices)

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (rate for today) or 'current' - current table / rate (last published)

func (t *NBPTable) GetTableCCurrent() ([]ExchangeTableC, error)
    GetTableCCurrent - function downloads current table of currency exchange
    rates and return slice of struct ExchangeTableC (or error), version for
    table C (bid, ask - buy, sell prices)

func (t *NBPTable) GetTableCToday() ([]ExchangeTableC, error)
    GetTableCToday - function downloads today's table of currency exchange rates
    and return slice of struct ExchangeTableC (or error), version for table C
    (bid, ask - buy, sell prices)

func (t *NBPTable) GetTableCurrent() ([]ExchangeTable, error)
    GetTableCurrent - function downloads current table of currency exchange
    rates and return slice of struct ExchangeTable (or error), version for table
    A, B (mid - average price)

func (t *NBPTable) GetTableToday() ([]ExchangeTable, error)
    GetTableToday - function downloads today's table of currency exchange rates
    and return slice of struct ExchangeTable (or error), version for table A, B
    (mid - average price)

func (t *NBPTable) SetTableType(tableType string) error
    SetTableType - the function allows to set the supported type of exchange
    rate table

func (t *NBPTable) TableByDate(date string) error
    TableByDate - function downloads and writes data to NBPTable.Exchange
    (NBPTable.ExchangeC) slice, raw data (json) still available in
    NBPTable.result field


    Function returns error or nil

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (rate for today) or 'current' - current table / rate (last published)

func (t *NBPTable) TableLast(last int) error
    TableLast - function downloads and writes data to NBPTable.Exchange
    (NBPTable.ExchangeC) slice, raw data (json) still available in
    NBPTable.result field


    Function returns error or nil

    Parameters:

        last - the last <n> tables/rates can be retrieved

func (t *NBPTable) TableRaw(date string, last int, format string) error
    TableRaw - function downloads data in json or xml form, the data will be
    placed as [] byte in the result field


    Function returns error or nil

    Parameters:

        date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
        or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
        (rate for today) or 'current' - current table / rate (last published)

        last - as an alternative to date, the last <n> tables/rates
        can be retrieved

        format - 'json' or 'xml'

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

- [ ]  more tests
