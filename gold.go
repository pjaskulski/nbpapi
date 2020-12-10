// prices of gold calculated at NBP
//
// types: NBPGold, GoldRate
// public func: NewGold
// NBPGold methods:
//          GoldRaw, GoldByDate, GoldLast,
//          GetPriceToday, GetPriceCurrent, GetPriceByDate
//          GetPrettyOutput, GetCSVOutput, GetRawOutput

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

// NewGold - function creates new gold type
func NewGold() *NBPGold {
	cli := &http.Client{
		Timeout: time.Second * 10,
	}
	r := &NBPGold{
		Client: cli,
	}
	return r
}

/*
GoldRaw - function downloads data in json or xml form

Function returns error or nil

Parameters:

    date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
    or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
    (price for today) or 'current' - current gold price (last published)

    last - as an alternative to date, the last <n> prices of gold
    can be retrieved

    format - 'json' or 'xml'
*/
func (g *NBPGold) GoldRaw(date string, last int, format string) error {
	var err error

	url := g.getGoldAddress(date, last)
	g.result, err = g.getData(url, format)
	if err != nil {
		return err
	}

	return err
}

/*
GoldByDate - function downloads and writes data to goldRates slice,
raw data (json) still available in NBPGold.result field

Function returns error or nil

Parameters:

    date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
    or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
    (price for today) or 'current' - current gold price (last published)

*/
func (g *NBPGold) GoldByDate(date string) error {
	var err error

	url := g.getGoldAddress(date, 0)
	g.result, err = g.getData(url, "json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(g.result, &g.GoldRates)
	if err != nil {
		return err
	}

	return err
}

/*
GoldLast - function downloads and writes data to GoldRates slice,
raw data (json) still available in NBPGold.result field

Function returns error or nil

Parameters:

    last - as an alternative to date, the last <n> prices of gold
    can be retrieved
*/
func (g *NBPGold) GoldLast(last int) error {
	var err error

	url := g.getGoldAddress("", last)
	g.result, err = g.getData(url, "json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(g.result, &g.GoldRates)
	if err != nil {
		return err
	}

	return err
}

/*
GetPriceToday - function downloads and returns today's gold price,
as GoldRate struct
*/
func (g *NBPGold) GetPriceToday() (GoldRate, error) {
	var err error

	url := g.getGoldAddress("today", 0)
	g.result, err = g.getData(url, "json")
	if err != nil {
		return GoldRate{}, err
	}

	err = json.Unmarshal(g.result, &g.GoldRates)
	if err != nil {
		return GoldRate{}, err
	}

	return g.GoldRates[0], err
}

/*
GetPriceCurrent - function downloads and returns current gold price as
GoldRate struct
*/
func (g *NBPGold) GetPriceCurrent() (GoldRate, error) {
	var err error

	url := g.getGoldAddress("current", 0)
	g.result, err = g.getData(url, "json")
	if err != nil {
		return GoldRate{}, err
	}

	err = json.Unmarshal(g.result, &g.GoldRates)
	if err != nil {
		return GoldRate{}, err
	}

	return g.GoldRates[0], err
}

/*
GetPriceByDate - function returns gold prices (as slice of struct),
by date ("YYYY-MM-DD") or range of dates ("YYYY-MM-DD:YYYY-MM-DD")

Parameters:

    date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
    or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
    (price for today) or 'current' - current gold price (last published)
*/
func (g *NBPGold) GetPriceByDate(date string) ([]GoldRate, error) {
	var err error

	url := g.getGoldAddress(date, 0)
	g.result, err = g.getData(url, "json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(g.result, &g.GoldRates)
	if err != nil {
		return nil, err
	}

	return g.GoldRates, err
}

/*
GetPrettyOutput - function returns a formatted table of gold prices

Parameters:

    lang - 'en' or 'pl'
*/
func (g *NBPGold) GetPrettyOutput(lang string) string {
	const padding = 3
	var builder strings.Builder

	// output language
	setLang(lang)

	w := tabwriter.NewWriter(&builder, 0, 0, padding, ' ', tabwriter.Debug)

	fmt.Fprintln(w)
	fmt.Fprintln(w, l.Get("The price of 1g of gold (of 1000 millesimal fineness)"))
	fmt.Fprintln(w)

	fmt.Fprintln(w, l.Get("DATE \t PRICE (PLN)"))
	fmt.Fprintln(w, l.Get("---- \t ----------- "))
	for _, goldItem := range g.GoldRates {
		goldValue := fmt.Sprintf("%.4f", goldItem.Price)
		fmt.Fprintln(w, goldItem.Data+" \t "+goldValue)
	}
	w.Flush()

	return builder.String()
}

/*
GetCSVOutput - function returns prices of gold in CSV format
(comma separated data)

Parameters:

    lang - 'en' or 'pl'
*/
func (g *NBPGold) GetCSVOutput(lang string) string {
	var output string = ""

	// output language
	setLang(lang)

	output += fmt.Sprintln(l.Get("DATE,PRICE (PLN)"))
	for _, goldItem := range g.GoldRates {
		goldValue := fmt.Sprintf("%.4f", goldItem.Price)
		output += fmt.Sprintln(goldItem.Data + "," + goldValue)
	}

	return output
}

// GetRawOutput - function returns just result of request (json or xml)
func (g *NBPGold) GetRawOutput() string {
	return string(g.result)
}

// ------------------- private func -----------------------------------

/* getData - function that retrieves data from the NBP website
   and returns them in the form of JSON / XML (or error), based on
   the arguments provided:

   url - NBP web api address
   format - 'json' or 'xml'
*/
func (g *NBPGold) getData(url, format string) ([]byte, error) {
	g.clearData()
	return fetchData(g.Client, url, format)
}

// clearData - data cleaning
func (g *NBPGold) clearData() {
	g.result = nil
	g.GoldRates = nil
}

// getGoldAddress - build download address depending on previously
// verified input parameters (date or last)
func (g *NBPGold) getGoldAddress(date string, last int) string {
	var address string

	if last != 0 {
		address = queryGoldLast(strconv.Itoa(last))
	} else if date == "today" {
		address = queryGoldToday()
	} else if date == "current" {
		address = queryGoldCurrent()
	} else if len(date) == 10 {
		address = queryGoldDate(date)
	} else if len(date) == 21 {
		address = queryGoldRange(date)
	}

	return address
}

// queryGoldToday - returns query: today's gold price
func queryGoldToday() string {
	return fmt.Sprintf("%s/today", baseAddressGold)
}

// queryGoldCurrent - returns query: current gold price
// (last published price)
func queryGoldCurrent() string {
	return baseAddressGold
}

// queryGoldLast - returns query: last <number> gold prices
func queryGoldLast(last string) string {
	return fmt.Sprintf("%s/last/%s", baseAddressGold, last)
}

// queryGoldDay - function returns gold price on the given date (RRRR-MM-DD)
// in json/xml form, or error
func queryGoldDate(date string) string {
	return fmt.Sprintf("%s/%s", baseAddressGold, date)
}

// queryGoldRange - returns query: gold prices within the given date range
// (RRRR-MM-DD:RRRR-MM-DD) in json/xml form, or error
func queryGoldRange(date string) string {
	var startDate string
	var stopDate string

	temp := strings.Split(date, ":")
	startDate = temp[0]
	stopDate = temp[1]

	address := fmt.Sprintf("%s/%s/%s", baseAddressGold, startDate, stopDate)
	return address
}
