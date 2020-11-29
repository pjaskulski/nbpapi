// examples of using the nbpapi module

package main

import (
	"fmt"

	"github.com/pjaskulski/nbpapi"
)

func main() {
	// current gold price: function GetPriceCurrent returns GoldPrice struct
	var price nbpapi.GoldRate
	var err error

	gold := nbpapi.NewGold()

	price, err = gold.GetPriceCurrent()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Publication date: ", price.Data)
		fmt.Println("Price of 1g of gold (PLN): ", price.Price)
	}
	fmt.Println()

	/*
		Gold price on November 12, 2020: function GetPriceByDate returns
		slice of GoldPrice struct, in case of date it is always 1 element,
		in case of range of date is more
	*/
	gold.GoldByDate("2020-11-12")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gold.GetPrettyOutput("en"))
	}
	fmt.Println()

	/*
		Gold prices between November 12, 2020 and November 19, 2020:
		function GetPriceByDate return slice of GoldPrice struct
	*/
	var prices []nbpapi.GoldRate
	prices, err = gold.GetPriceByDate("2020-11-12:2020-11-19")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Date", "\t\t", "Price")
		for _, rate := range prices {
			fmt.Println(rate.Data, "\t", rate.Price)
		}
	}
}
