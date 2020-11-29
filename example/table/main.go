// examples of using the nbpapi module

package main

import (
	"fmt"

	"github.com/pjaskulski/nbpapi"
)

func main() {
	/*
		How to get table A of currency exchange rates published
		on 12 Nov 2020: function TableByDate("2020-11-12") fetch data.
		The data can be read from nbpTable.Exchange structures (or
		nbpTable.ExchangeC structures, depending of the type of
		table of exchange rates)
	*/
	var err error
	var tableNo string

	nbpTable := nbpapi.NewTable("A")
	err = nbpTable.TableByDate("2020-11-12")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, item := range nbpTable.Exchange {
			tableNo = item.No
			for _, currencyItem := range item.Rates {
				fmt.Println(tableNo, currencyItem.Code, currencyItem.Currency, currencyItem.Mid)
			}
		}
	}
}
