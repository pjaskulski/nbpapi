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
	"net/http"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

// NewTable - function creates new table type
func NewTable(tableType string) *NBPTable {
	cli := &http.Client{
		Timeout: time.Second * 10,
	}
	r := &NBPTable{
		tableType: tableType,
		Client:    cli,
	}
	return r
}

/*
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
*/
func (t *NBPTable) TableRaw(date string, last int, format string) error {
	var err error

	url := getTableAddress(t.tableType, date, last)
	t.result, err = t.getData(url, format)
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

    date - date in the format: 'YYYY-MM-DD' (ISO 8601 standard),
    or a range of dates in the format: 'YYYY-MM-DD:YYYY-MM-DD' or 'today'
    (rate for today) or 'current' - current table / rate (last published)
*/
func (t *NBPTable) TableByDate(date string) error {
	var err error

	url := getTableAddress(t.tableType, date, 0)
	t.result, err = t.getData(url, "json")
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

    last - the last <n> tables/rates can be retrieved

*/
func (t *NBPTable) TableLast(last int) error {
	var err error

	url := getTableAddress(t.tableType, "", last)
	t.result, err = t.getData(url, "json")
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

	url := getTableAddress(t.tableType, "current", 0)
	t.result, err = t.getData(url, "json")
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

	url := getTableAddress(t.tableType, "current", 0)
	t.result, err = t.getData(url, "json")
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

/* getData - function that retrieves data from the NBP website
   and returns them in the form of JSON / XML (or error), based on
   the arguments provided:

   url - NBP web api address
   format - 'json' or 'xml'
*/
func (t *NBPTable) getData(url string, format string) ([]byte, error) {
	return fetchData(t.Client, url, format)
}

// --------------------- Private func ---------------------------------

// getTableAddress - build download address depending on previously
// verified input parameters (--table, --date or --last)
func getTableAddress(tableType string, date string, last int) string {
	var address string

	if last != 0 {
		address = queryTableLast(tableType, strconv.Itoa(last))
	} else if date == "today" {
		address = queryTableToday(tableType)
	} else if date == "current" {
		address = queryTableCurrent(tableType)
	} else if len(date) == 10 {
		address = queryTableDate(tableType, date)
	} else if len(date) == 21 {
		address = queryTableRange(tableType, date)
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
func queryTableDate(tableType string, date string) string {
	return fmt.Sprintf("%s/tables/%s/%s/", baseAddressTable, tableType, date)
}

// queryTableRange - returns query: table of exchange rates  within
// the given date range (RRRR-MM-DD:RRRR-MM-DD)
func queryTableRange(tableType string, date string) string {
	var startDate string
	var stopDate string

	temp := strings.Split(date, ":")
	startDate = temp[0]
	stopDate = temp[1]

	address := fmt.Sprintf("%s/tables/%s/%s/%s/", baseAddressTable, tableType, startDate, stopDate)
	return address
}

// queryTableLast - returns query: last <number> tables of exchange rates
func queryTableLast(tableType string, last string) string {
	return fmt.Sprintf("%s/tables/%s/last/%s/", baseAddressTable, tableType, last)
}
