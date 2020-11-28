// complete tables of currency exchange rates published by the NBP.PL service
//
// public func: NewTable
// types: NBPTable, ExchangeTable, ExchangeTableC
// NBPTable methods:
//			TableRaw, TableByDate, TableLast
//          GetTableCurrent, GetTableCCurrent
//			GetPrettyOutput, GetCSVOutput, GetRawOutput

package nbpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"
)

// base addresses of the NBP API service
const (
	baseAddressTable string = "http://api.nbp.pl/api/exchangerates"
)

// NBPTable type
type NBPTable struct {
	tableType string
	result    []byte
	Exchange  []ExchangeTable
	ExchangeC []ExchangeTableC
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

// Public func

// NewTable - function creates new table type
func NewTable(tFlag string) *NBPTable {
	return &NBPTable{
		tableType: tFlag,
	}
}

/*
TableRaw - function downloads data in json or xml form, the data will be
placed as [] byte in the result field

Function returns error or nil

Parameters:

    dFlag - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
	or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
	(rate for today) or 'current' - current table / rate (last published)

	lFlag - as an alternative to date, the last <n> tables/rates
	can be retrieved

	repFormat - 'json' or 'xml'
*/
func (t *NBPTable) TableRaw(dFlag string, lFlag int, repFormat string) error {
	var err error

	address := getTableAddress(t.tableType, dFlag, lFlag)
	t.result, err = getData(address, repFormat)
	if err != nil {
		return err
	}

	return err
}

/*
TableByDate - function downloads and writes data to NBPTable.Exchange
(NBPTable.ExchangeC) slice, raw data (json) still available in
NBPTable.result field

Function returns error or nil

Parameters:

	dFlag - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
	or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
	(rate for today) or 'current' - current table / rate (last published)
*/
func (t *NBPTable) TableByDate(dFlag string) error {
	var err error

	address := getTableAddress(t.tableType, dFlag, 0)
	t.result, err = getData(address, "json")
	if err != nil {
		return err
	}

	if t.tableType != "C" {
		err = json.Unmarshal(t.result, &t.Exchange)
	} else {
		err = json.Unmarshal(t.result, &t.ExchangeC)
	}
	if err != nil {
		return err
	}

	return err
}

/*
TableLast - function downloads and writes data to NBPTable.Exchange
(NBPTable.ExchangeC) slice, raw data (json) still available in
NBPTable.result field

Function returns error or nil

Parameters:

	lFlag - the last <n> tables/rates can be retrieved

*/
func (t *NBPTable) TableLast(lFlag int) error {
	var err error

	address := getTableAddress(t.tableType, "", lFlag)
	t.result, err = getData(address, "json")
	if err != nil {
		return err
	}

	if t.tableType != "C" {
		err = json.Unmarshal(t.result, &t.Exchange)
	} else {
		err = json.Unmarshal(t.result, &t.ExchangeC)
	}
	if err != nil {
		return err
	}

	return err
}

/*
GetTableCurrent - function downloads current table of currency exchange
rates and return slice of struct ExchangeTable (or error), version for
table A, B (mid - average price)
*/
func (t *NBPTable) GetTableCurrent() ([]ExchangeTable, error) {
	var err error

	if t.tableType == "C" {
		return nil, errors.New("Invalid function call context, use GetTableCCurrent")
	}

	address := getTableAddress(t.tableType, "current", 0)
	t.result, err = getData(address, "json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(t.result, &t.Exchange)
	if err != nil {
		return nil, err
	}

	return t.Exchange, err
}

/*
GetTableCCurrent - function downloads current table of currency exchange
rates and return slice of struct ExchangeTableC (or error), version for
table C (bid, ask - buy, sell prices)
*/
func (t *NBPTable) GetTableCCurrent() ([]ExchangeTableC, error) {
	var err error

	if t.tableType != "C" {
		return nil, errors.New("Invalid function call context, use GetTableCurrent")
	}

	address := getTableAddress(t.tableType, "current", 0)
	t.result, err = getData(address, "json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(t.result, &t.ExchangeC)
	if err != nil {
		return nil, err
	}

	return t.ExchangeC, err
}

/*
GetPrettyOutput - function returns tables of exchange rates as
formatted table, depending on the tableType field: for type A and B tables
a column with an average rate is printed, for type C two columns:
buy price and sell price

Parameters:

	lang - 'en' or 'pl'
*/
func (t *NBPTable) GetPrettyOutput(lang string) string {
	const padding = 3
	var builder strings.Builder
	var output string

	// output language
	setLang(lang)

	w := tabwriter.NewWriter(&builder, 0, 0, padding, ' ', tabwriter.Debug)

	if t.tableType != "C" {
		for _, item := range t.Exchange {
			output += fmt.Sprintln()
			output += fmt.Sprintln(l.Get("Table type:")+"\t\t", item.Table)
			output += fmt.Sprintln(l.Get("Table number:")+"\t\t", item.No)
			output += fmt.Sprintln(l.Get("Publication date:")+"\t", item.EffectiveDate)
			output += fmt.Sprintln()

			fmt.Fprintln(w, l.Get("CODE \t NAME \t AVERAGE (PLN)"))
			fmt.Fprintln(w, l.Get("---- \t ---- \t -------------"))
			for _, currencyItem := range item.Rates {
				currencyValue := fmt.Sprintf("%.4f", currencyItem.Mid)
				fmt.Fprintln(w, currencyItem.Code+" \t "+currencyItem.Currency+" \t "+currencyValue)
			}
			w.Flush()
			output += builder.String()
			builder.Reset()
		}
	} else {
		for _, item := range t.ExchangeC {
			output += fmt.Sprintln()
			output += fmt.Sprintln(l.Get("Table type:")+"\t\t", item.Table)
			output += fmt.Sprintln(l.Get("Table number:")+"\t\t", item.No)
			output += fmt.Sprintln(l.Get("Trading date:")+"\t\t", item.TradingDate)
			output += fmt.Sprintln(l.Get("Publication date:")+"\t", item.EffectiveDate)
			output += fmt.Sprintln()

			fmt.Fprintln(w, l.Get("CODE \t NAME \t BUY (PLN) \t SELL (PLN) "))
			fmt.Fprintln(w, l.Get("---- \t ---- \t --------- \t ---------- "))
			for _, currencyItem := range item.Rates {
				currencyValueBid := fmt.Sprintf("%.4f", currencyItem.Bid)
				currencyValueAsk := fmt.Sprintf("%.4f", currencyItem.Ask)
				fmt.Fprintln(w, currencyItem.Code+" \t "+currencyItem.Currency+" \t "+currencyValueBid+" \t "+currencyValueAsk)
			}
			w.Flush()
			output += builder.String()
			builder.Reset()
		}
	}

	return output
}

/*
GetCSVOutput - function prints tables of exchange rates in the console,
in the form of CSV (data separated by a comma), depending on the
tableType field: for type A and B tables a column with an average
rate is printed, for type C two columns: buy price and sell price

Parameters:

	lang - 'en' or 'pl'
*/
func (t *NBPTable) GetCSVOutput(lang string) string {
	var tableNo string
	var output string = ""

	// output language
	setLang(lang)

	if t.tableType != "C" {
		output += fmt.Sprintln(l.Get("TABLE,CODE,NAME,AVERAGE (PLN)"))

		for _, item := range t.Exchange {
			tableNo = item.No
			for _, currencyItem := range item.Rates {
				currencyValue := fmt.Sprintf("%.4f", currencyItem.Mid)
				output += fmt.Sprintln(tableNo + "," + currencyItem.Code + "," + currencyItem.Currency + "," + currencyValue)
			}
		}
	} else {
		output += fmt.Sprintln(l.Get("TABLE,CODE,NAME,BUY (PLN),SELL (PLN)"))

		for _, item := range t.ExchangeC {
			tableNo = item.No
			for _, currencyItem := range item.Rates {
				currencyValueBid := fmt.Sprintf("%.4f", currencyItem.Bid)
				currencyValueAsk := fmt.Sprintf("%.4f", currencyItem.Ask)
				output += fmt.Sprintln(tableNo + "," + currencyItem.Code + "," + currencyItem.Currency + "," + currencyValueBid + "," + currencyValueAsk)
			}
		}
	}

	return output
}

// GetRawOutput - function returns just result of request (json or xml)
func (t *NBPTable) GetRawOutput() string {
	return string(t.result)
}

// Private func

// getTableAddress - build download address depending on previously
// verified input parameters (--table, --date or --last)
func getTableAddress(tableType string, dFlag string, lFlag int) string {
	var address string

	if lFlag != 0 {
		address = queryTableLast(tableType, strconv.Itoa(lFlag))
	} else if dFlag == "today" {
		address = queryTableToday(tableType)
	} else if dFlag == "current" {
		address = queryTableCurrent(tableType)
	} else if len(dFlag) == 10 {
		address = queryTableDay(tableType, dFlag)
	} else if len(dFlag) == 21 {
		address = queryTableRange(tableType, dFlag)
	}

	return address
}

// queryTableToday - returns query: exchange rate table published today
func queryTableToday(tableType string) string {
	return fmt.Sprintf("%s/tables/%s/today/", baseAddressTable, tableType)
}

// queryTableCurrent - returns query: current table of exchange rates
// (last published table)
func queryTableCurrent(tableType string) string {
	return fmt.Sprintf("%s/tables/%s/", baseAddressTable, tableType)
}

// queryTableDay - returns query: table of exchange rates
// on the given date (YYYY-MM-DD)
func queryTableDay(tableType string, day string) string {
	return fmt.Sprintf("%s/tables/%s/%s/", baseAddressTable, tableType, day)
}

// queryTableRange - returns query: table of exchange rates  within
// the given date range (RRRR-MM-DD:RRRR-MM-DD)
func queryTableRange(tableType string, day string) string {
	var startDate string
	var stopDate string

	temp := strings.Split(day, ":")
	startDate = temp[0]
	stopDate = temp[1]

	address := fmt.Sprintf("%s/tables/%s/%s/%s/", baseAddressTable, tableType, startDate, stopDate)
	return address
}

// queryTableLast - returns query: last <number> tables of exchange rates
func queryTableLast(tableType string, last string) string {
	return fmt.Sprintf("%s/tables/%s/last/%s/", baseAddressTable, tableType, last)
}
