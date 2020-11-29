// examples of using the nbpapi module

package main

import (
	"fmt"

	"github.com/pjaskulski/nbpapi"
)

func main() {
	/*
		How to get a latest 5 exchange rates of currency CHF,
		table of type A (mid - average exchange rate), as CSV data
	*/
	nbpMid := nbpapi.NewCurrency("A")
	err := nbpMid.CurrencyLast("CHF", 5)
	if err != nil {
		fmt.Println(err)
	}

	// print polish version
	fmt.Println(nbpMid.GetPrettyOutput("pl"))
	fmt.Println()

	/*
		How to get today's rate of CHF, table C (bid, ask - buy and sell
		exchange rate): function GetCurencyToday, returns error or nil and
		populates slice ExchangeC, the currency rate is taken from the first
		element of the ExchangeC, later an example of a more convenient
		function: GetRateToday that returns a struct.
	*/
	nbpC := nbpapi.NewCurrency("C")
	err = nbpC.CurrencyToday("CHF")
	if err != nil {
		fmt.Println("Currently, there is no todays CHF exchange rate")
	} else {
		fmt.Println("Table No: ", nbpC.ExchangeC.Rates[0].No)
		fmt.Println("Date: ", nbpC.ExchangeC.Rates[0].EffectiveDate)
		fmt.Println("Bid: ", nbpC.ExchangeC.Rates[0].Bid)
		fmt.Println("Ask: ", nbpC.ExchangeC.Rates[0].Ask)
	}
	fmt.Println()

	/*
		Alternatively how to get today's rate of CHF, table C (bid, ask -
		buy and sell exchange rates): function GetRateToday return Rate
		struct or error, the ExchangeC slice is also populated, but reading
		the data from the Rate structure is more convenient
	*/
	var today nbpapi.Rate
	today, err = nbpC.GetRateToday("CHF")
	if err != nil {
		fmt.Println("Currently, there is no todays CHF exchange rate")
	} else {
		fmt.Println("Table No: ", today.No)
		fmt.Println("Date: ", today.EffectiveDate)
		fmt.Println("Bid: ", today.Bid)
		fmt.Println("Ask: ", today.Ask)
	}
	fmt.Println()

	/*
		How to get current rate of CHF, table C (bid, ask - buy and sell
		exchange rate): function GetRateCurrent return Rate struct or error,
		the ExchangeC table is also populated, but reading the data from the
		Rate structure is more convenient
	*/
	var result nbpapi.Rate
	result, err = nbpC.GetRateCurrent("CHF")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Table No: ", result.No)
		fmt.Println("Date: ", result.EffectiveDate)
		fmt.Println("Bid: ", result.Bid)
		fmt.Println("Ask: ", result.Ask)
	}
	fmt.Println()

	/*
		Currency exchange rate for date or range of dates:
		function GetRateByDate return slice of Rate struct or error
	*/
	var results []nbpapi.Rate
	results, err = nbpC.GetRateByDate("CHF", "2020-11-12")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, item := range results {
			fmt.Println(item.No, item.EffectiveDate, item.Bid, item.Ask)
		}
	}
	fmt.Println()
}
