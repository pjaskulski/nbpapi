// examples of using the nbpapi module

package main

import (
	"fmt"
	"log"

	"github.com/pjaskulski/nbpapi"
)

func main() {
	// How to get table A of currency exchange rates published
	// on 12 Nov 2020: function TableByDate("2020-11-12") fetch data, then
	// one can get data from nbpTable.Exchange structures (or
	// nbpTable.ExchangeC structures, depending of the type of
	// table of exchange rates)
	var err error
	var tableNo string

	nbpTable := nbpapi.NewTable("A")
	err = nbpTable.TableByDate("2020-11-12")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range nbpTable.Exchange {
		tableNo = item.No
		for _, currencyItem := range item.Rates {
			fmt.Println(tableNo, currencyItem.Code, currencyItem.Currency, currencyItem.Mid)
		}
	}
	fmt.Println()

	// get current table A
	var tableA []nbpapi.ExchangeTable
	tableA, err = nbpTable.GetTableCurrent()
	if err != nil {
		fmt.Println(err)
	}

	for _, tItem := range tableA {
		for _, item := range tItem.Rates {
			fmt.Println(tItem.No, item.Code, item.Currency, item.Mid)
		}
	}
	fmt.Println()

	// how to get a latest 5 exchange rates of currency CHF,
	// table of type A (mid - average exchange rate), as CSV data
	nbpMid := nbpapi.NewCurrency("A")
	err = nbpMid.CurrencyLast("CHF", 5)
	if err != nil {
		log.Fatal(err)
	}

	output := nbpMid.GetCSVOutput()
	fmt.Println(output)
	fmt.Println()

	// polish version
	nbpapi.SetLang("pl")
	fmt.Println(nbpMid.GetCSVOutput())
	fmt.Println()

	nbpapi.SetLang("en")
	// how to get today's rate of CHF, table C (bid, ask - buy and sell
	// exchange rate), function GetCurencyToday, returns error or nil and
	// populates slice ExchangeC, the currency rate is taken from the first
	// element of the ExchangeC, later an example of a more convenient
	// function: GetRateToday that returns a struct.
	nbpC := nbpapi.NewCurrency("C")
	err = nbpC.CurrencyToday("CHF")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Table No: ", nbpC.ExchangeC.Rates[0].No)
		fmt.Println("Date: ", nbpC.ExchangeC.Rates[0].EffectiveDate)
		fmt.Println("Bid: ", nbpC.ExchangeC.Rates[0].Bid)
		fmt.Println("Ask: ", nbpC.ExchangeC.Rates[0].Ask)
	}
	fmt.Println()

	// alternatively how to get today's rate of CHF, table C (bid, ask -
	// buy and sell exchange rates): function GetRateToday return Rate
	// struct or error, the ExchangeC slice is also populated, but reading
	// the data from the Rate structure is more convenient
	var today nbpapi.Rate
	today, err = nbpC.GetRateToday("CHF")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Table No: ", today.No)
		fmt.Println("Date: ", today.EffectiveDate)
		fmt.Println("Bid: ", today.Bid)
		fmt.Println("Ask: ", today.Ask)
	}
	fmt.Println()

	// how to get current rate of CHF, table C (bid, ask - buy and sell
	// exchange rate): function GetRateCurrent return Rate struct or error,
	// the ExchangeC table is also populated, but reading the data from the
	// Rate structure is more convenient
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

	// currency exchange rate for date or range of dates:
	// function GetRateByDate return slice of Rate struct or error
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

	// today's gold price: function GetPriceToday returns GoldPrice struct
	var price nbpapi.GoldRate
	gold := nbpapi.NewGold()

	price, err = gold.GetPriceToday()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Publication date: ", price.Data)
		fmt.Println("Price of 1g of gold (PLN): ", price.Price)
	}
	fmt.Println()

	// current gold price: function GetPriceCurrent returns GoldPrice struct
	price, err = gold.GetPriceCurrent()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Publication date: ", price.Data)
		fmt.Println("Price of 1g of gold (PLN): ", price.Price)
	}
	fmt.Println()

	// gold price on November 12, 2020: function GetPriceByDate returns
	// slice of GoldPrice struct, in case of date it is always 1 element,
	// in case of range of date is more
	var prices []nbpapi.GoldRate
	prices, err = gold.GetPriceByDate("2020-11-12")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, rate := range prices {
			fmt.Println("Publication date: ", rate.Data)
			fmt.Println("Price of 1g of gold (PLN): ", rate.Price)
		}
	}
	fmt.Println()

	// gold prices between November 12, 2020 and November 19, 2020:
	// function GetPriceByDate return slice of GoldPrice struct
	prices, err = gold.GetPriceByDate("2020-11-12:2020-11-19")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, rate := range prices {
			fmt.Println("Publication date: ", rate.Data)
			fmt.Println("Price of 1g of gold (PLN): ", rate.Price)
		}
	}
}
