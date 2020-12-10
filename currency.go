// particular currency exchange rates
//
// public func: NewCurrency
// types: NBPCurrency, Rate
// NBPCurrency methods:
//            CurrencyRaw, CurrencyByDate, CurrencyLast, CurrencyToday,
//            GetRateCurrent, GetRateToday, GetRateByDate,
//            GetPrettyOutput, GetCSVOutput, GetRawOutput

package nbpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

// NewCurrency - function creates new currency type
func NewCurrency(tableType string) *NBPCurrency {
	cli := &http.Client{
		Timeout: time.Second * 10,
	}
	r := &NBPCurrency{
		tableType: tableType,
		Client:    cli,
	}
	return r
}

/*
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
*/
func (c *NBPCurrency) CurrencyRaw(code, date string, last int, format string) error {
	var err error

	err = checkCurrencyCode(c.tableType, code)
	if err != nil {
		return err
	}

	url := c.getCurrencyAddress(c.tableType, date, last, code)
	c.result, err = c.getData(url, format)
	if err != nil {
		return err
	}

	return err
}

/*
CurrencyByDate - function downloads and writes data to exchange (exchangeC) slice,
raw data (json) still available in result field

Function returns error or nil

Parameters:

    date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
    or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
    (rate for today) or 'current' - current table / rate (last published)

    code - ISO 4217 currency code, depending on the type of the
    table available currencies may vary
*/
func (c *NBPCurrency) CurrencyByDate(code, date string) error {
	var err error

	err = checkCurrencyCode(c.tableType, code)
	if err != nil {
		return err
	}

	url := c.getCurrencyAddress(c.tableType, date, 0, code)
	c.result, err = c.getData(url, "json")
	if err != nil {
		return err
	}

	if c.tableType != "C" {
		err = json.Unmarshal(c.result, &c.Exchange)
	} else {
		err = json.Unmarshal(c.result, &c.ExchangeC)
	}
	if err != nil {
		return err
	}

	return err
}

/*
CurrencyLast - function downloads and writes data to exchange (exchangeC) slice,
raw data (json) still available in result field

Function returns error or nil

Parameters:

    code - ISO 4217 currency code, depending on the type of the
    table available currencies may vary

    last - as an alternative to date, the last <n> tables/rates
    can be retrieved
*/
func (c *NBPCurrency) CurrencyLast(code string, last int) error {
	var err error

	err = checkCurrencyCode(c.tableType, code)
	if err != nil {
		return err
	}

	url := c.getCurrencyAddress(c.tableType, "", last, code)
	c.result, err = c.getData(url, "json")
	if err != nil {
		return err
	}

	if c.tableType != "C" {
		err = json.Unmarshal(c.result, &c.Exchange)
	} else {
		err = json.Unmarshal(c.result, &c.ExchangeC)
	}
	if err != nil {
		return err
	}

	return err
}

/*
CurrencyToday - function downloads and writes data to exchange (exchangeC) slice,
raw data (json) still available in result field

Function returns error or nil

Parameters:

    code - ISO 4217 currency code, depending on the type of the
    table available currencies may vary
*/
func (c *NBPCurrency) CurrencyToday(code string) error {
	var err error

	err = checkCurrencyCode(c.tableType, code)
	if err != nil {
		return err
	}

	url := c.getCurrencyAddress(c.tableType, "today", 0, code)
	c.result, err = c.getData(url, "json")
	if err != nil {
		return err
	}

	if c.tableType != "C" {
		err = json.Unmarshal(c.result, &c.Exchange)
	} else {
		err = json.Unmarshal(c.result, &c.ExchangeC)
	}
	if err != nil {
		return err
	}

	return err
}

/*
GetRateCurrent - function downloads current currency exchange rate
and return Rate struct (or error)

Parameters:

    code - ISO 4217 currency code, depending on the type of the
    table available currencies may vary
*/
func (c *NBPCurrency) GetRateCurrent(code string) (Rate, error) {
	var err error
	var rate Rate

	err = checkCurrencyCode(c.tableType, code)
	if err != nil {
		return rate, err
	}

	url := c.getCurrencyAddress(c.tableType, "current", 0, code)
	c.result, err = c.getData(url, "json")
	if err != nil {
		return rate, err
	}

	if c.tableType != "C" {
		err = json.Unmarshal(c.result, &c.Exchange)
	} else {
		err = json.Unmarshal(c.result, &c.ExchangeC)
	}
	if err != nil {
		return rate, err
	}

	if c.tableType != "C" {
		rate.No = c.Exchange.Rates[0].No
		rate.EffectiveDate = c.Exchange.Rates[0].EffectiveDate
		rate.Mid = c.Exchange.Rates[0].Mid
		rate.Ask = 0
		rate.Bid = 0
	} else {
		rate.No = c.ExchangeC.Rates[0].No
		rate.EffectiveDate = c.ExchangeC.Rates[0].EffectiveDate
		rate.Mid = 0
		rate.Ask = c.ExchangeC.Rates[0].Ask
		rate.Bid = c.ExchangeC.Rates[0].Bid
	}

	return rate, err
}

/*
GetRateToday - function downloads today's currency exchange rate
and returns Rate struct (or error)

Parameters:

    code - ISO 4217 currency code, depending on the type of the
    table available currencies may vary
*/
func (c *NBPCurrency) GetRateToday(code string) (Rate, error) {
	var err error
	var rate Rate

	err = checkCurrencyCode(c.tableType, code)
	if err != nil {
		return rate, err
	}

	url := c.getCurrencyAddress(c.tableType, "today", 0, code)
	c.result, err = c.getData(url, "json")
	if err != nil {
		return rate, err
	}

	if c.tableType != "C" {
		err = json.Unmarshal(c.result, &c.Exchange)
	} else {
		err = json.Unmarshal(c.result, &c.ExchangeC)
	}
	if err != nil {
		return rate, err
	}

	if c.tableType != "C" {
		rate.No = c.Exchange.Rates[0].No
		rate.EffectiveDate = c.Exchange.Rates[0].EffectiveDate
		rate.Mid = c.Exchange.Rates[0].Mid
		rate.Ask = 0
		rate.Bid = 0
	} else {
		rate.No = c.ExchangeC.Rates[0].No
		rate.EffectiveDate = c.ExchangeC.Rates[0].EffectiveDate
		rate.Mid = 0
		rate.Ask = c.ExchangeC.Rates[0].Ask
		rate.Bid = c.ExchangeC.Rates[0].Bid
	}

	return rate, err
}

/*
GetRateByDate - function downloads today's currency exchange rate
and returns slice of Rate struct (or error)

Parameters:

    code - ISO 4217 currency code, depending on the type of the
    table available currencies may vary

    date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
    or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
    (rate for today) or 'current' - current table / rate (last published)
*/
func (c *NBPCurrency) GetRateByDate(code, date string) ([]Rate, error) {
	var err error
	var rates []Rate
	var rate Rate

	err = checkCurrencyCode(c.tableType, code)
	if err != nil {
		return nil, err
	}

	url := c.getCurrencyAddress(c.tableType, date, 0, code)
	c.result, err = c.getData(url, "json")
	if err != nil {
		return nil, err
	}

	if c.tableType != "C" {
		err = json.Unmarshal(c.result, &c.Exchange)
	} else {
		err = json.Unmarshal(c.result, &c.ExchangeC)
	}
	if err != nil {
		return nil, err
	}

	if c.tableType != "C" {
		for _, item := range c.Exchange.Rates {
			rate.No = item.No
			rate.EffectiveDate = item.EffectiveDate
			rate.Mid = item.Mid
			rate.Ask = 0
			rate.Bid = 0
			rates = append(rates, rate)
		}
	} else {
		for _, item := range c.ExchangeC.Rates {
			rate.No = item.No
			rate.EffectiveDate = item.EffectiveDate
			rate.Mid = 0
			rate.Ask = item.Ask
			rate.Bid = item.Bid
			rates = append(rates, rate)
		}
	}

	return rates, err
}

/*
GetPrettyOutput - function returns exchange rates as formatted table
depending on the tableType field:
for type A and B tables a column with an average rate is printed,
for type C two columns: buy price and sell price

Parameters:

    lang - 'en' or 'pl'
*/
func (c *NBPCurrency) GetPrettyOutput(lang string) string {
	const padding = 3
	var builder strings.Builder
	var output string

	// output language
	setLang(lang)

	w := tabwriter.NewWriter(&builder, 0, 0, padding, ' ', tabwriter.Debug)

	if c.tableType != "C" {
		output += fmt.Sprintln()
		output += fmt.Sprintln(l.Get("Table type:")+"\t", c.Exchange.Table)
		output += fmt.Sprintln(l.Get("Currency name:")+"\t", c.Exchange.Currency)
		output += fmt.Sprintln(l.Get("Currency code:")+"\t", c.Exchange.Code)
		output += fmt.Sprintln()

		fmt.Fprintln(w, l.Get("TABLE \t DATE \t AVERAGE (PLN)"))
		fmt.Fprintln(w, l.Get("----- \t ---- \t -------------"))
		for _, currencyItem := range c.Exchange.Rates {
			currencyValue := fmt.Sprintf("%.4f", currencyItem.Mid)
			fmt.Fprintln(w, currencyItem.No+" \t "+currencyItem.EffectiveDate+" \t "+currencyValue)
		}
	} else {
		output += fmt.Sprintln()
		output += fmt.Sprintln(l.Get("Table type:")+"\t", c.ExchangeC.Table)
		output += fmt.Sprintln(l.Get("Currency name:")+"\t", c.ExchangeC.Currency)
		output += fmt.Sprintln(l.Get("Currency code:")+"\t", c.ExchangeC.Code)
		output += fmt.Sprintln()

		fmt.Fprintln(w, l.Get("TABLE \t DATE \t BUY (PLN) \t SELL (PLN) "))
		fmt.Fprintln(w, l.Get("----- \t ---- \t --------- \t ---------- "))
		for _, currencyItem := range c.ExchangeC.Rates {
			currencyValueBid := fmt.Sprintf("%.4f", currencyItem.Bid)
			currencyValueAsk := fmt.Sprintf("%.4f", currencyItem.Ask)
			fmt.Fprintln(w, currencyItem.No+" \t "+currencyItem.EffectiveDate+" \t "+currencyValueBid+" \t "+currencyValueAsk)
		}
	}
	w.Flush()

	return output + builder.String()
}

/*
GetCSVOutput - function returns currency rates,
in the form of CSV (data separated by a comma), depending on the
tableType field: for type A and B tables a column with an average
rate is printed, for type C two columns: buy price and sell price

Parameters:

    lang - 'en' or 'pl'
*/
func (c *NBPCurrency) GetCSVOutput(lang string) string {
	var output string = ""

	// output language
	setLang(lang)

	if c.tableType != "C" {
		output += fmt.Sprintln(l.Get("TABLE,DATE,AVERAGE (PLN)"))
		for _, currencyItem := range c.Exchange.Rates {
			currencyValue := fmt.Sprintf("%.4f", currencyItem.Mid)
			output += fmt.Sprintln(currencyItem.No + "," + currencyItem.EffectiveDate + "," + currencyValue)
		}
	} else {
		output += fmt.Sprintln(l.Get("TABLE,DATE,BUY (PLN),SELL (PLN)"))
		for _, currencyItem := range c.ExchangeC.Rates {
			currencyValueBid := fmt.Sprintf("%.4f", currencyItem.Bid)
			currencyValueAsk := fmt.Sprintf("%.4f", currencyItem.Ask)
			output += fmt.Sprintln(currencyItem.No + "," + currencyItem.EffectiveDate + "," + currencyValueBid + "," + currencyValueAsk)
		}
	}

	return output
}

// GetRawOutput - function print just result of request (json or xml)
func (c *NBPCurrency) GetRawOutput() string {
	return string(c.result)
}

// SetTableType - the function allows to set the supported type of exchange rate table
func (c *NBPCurrency) SetTableType(tableType string) error {
	err := checkTableType(tableType)
	if err != nil {
		return err
	}

	c.tableType = tableType
	return nil
}

/* getData - function that retrieves data from the NBP website
   and returns them in the form of JSON / XML (or error), based on
   the arguments provided:

   url - NBP web api address
   format - 'json' or 'xml'
*/
func (c *NBPCurrency) getData(url, format string) ([]byte, error) {
	return fetchData(c.Client, url, format)
}

// --------------------- Private func ---------------------------------

// getCurrencyAddress - function builds download address depending on previously
// verified input parameters (--table, --date or --last, --code)
func (c *NBPCurrency) getCurrencyAddress(tableType, date string, last int, code string) string {
	var address string

	if last != 0 {
		address = queryCurrencyLast(tableType, strconv.Itoa(last), code)
	} else if date == "today" {
		address = queryCurrencyToday(tableType, code)
	} else if date == "current" {
		address = queryCurrencyCurrent(tableType, code)
	} else if len(date) == 10 {
		address = queryCurrencyDate(tableType, date, code)
	} else if len(date) == 21 {
		address = queryCurrencyRange(tableType, date, code)
	}

	return address
}

// queryCurrencyLast - returns query: last <number> currency exchange
// rates in json/xml form, or error
func queryCurrencyLast(tableType, last, currency string) string {
	return fmt.Sprintf("%s/rates/%s/%s/last/%s/", baseAddressCurrency, tableType, currency, last)

}

// queryCurrencyToday - returns query: today's currency exchange rate
func queryCurrencyToday(tableType, currency string) string {
	return fmt.Sprintf("%s/rates/%s/%s/today/", baseAddressCurrency, tableType, currency)
}

// queryCurrencyCurrent - returns query: current exchange rate for
// particular currency (last published price)
func queryCurrencyCurrent(tableType, currency string) string {
	return fmt.Sprintf("%s/rates/%s/%s/", baseAddressCurrency, tableType, currency)
}

// queryCurrencyDay - returns query: exchange rate for particular currency
// on the given date (YYYY-MM-DD)
func queryCurrencyDate(tableType, date, currency string) string {
	return fmt.Sprintf("%s/rates/%s/%s/%s/", baseAddressCurrency, tableType, currency, date)
}

// queryCurrencyRange - returns query: exchange rate for particular currency
// within the given date range (RRRR-MM-DD:RRRR-MM-DD)
func queryCurrencyRange(tableType, date, currency string) string {
	var startDate string
	var stopDate string

	temp := strings.Split(date, ":")
	startDate = temp[0]
	stopDate = temp[1]

	address := fmt.Sprintf("%s/rates/%s/%s/%s/%s/", baseAddressCurrency, tableType, currency, startDate, stopDate)
	return address
}
