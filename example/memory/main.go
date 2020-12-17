// examples of using the nbpapi module

package main

import (
	"fmt"
	"time"

	"github.com/pjaskulski/nbpapi"
)

func main() {
	// cache example
	var err error

	gold := nbpapi.NewGold()
	nbpapi.EnableCache() // default cache expiration time: 60 minutes

	/*
		Gold price on November 12, 2020: function GetPriceByDate returns
		slice of GoldPrice struct, in case of date it is always 1 element,
		in case of range of date is more
	*/

	// from NBP
	startTime := time.Now()
	err = gold.GoldByDate("2020-11-12")
	if err != nil {
		fmt.Println(err)
	}
	endTime := time.Now()
	fmt.Println("Data from the NBP server :", endTime.Sub(startTime))

	fmt.Println(gold.CreatePrettyOutput("en"))
	fmt.Println()

	// from cache
	startTime = time.Now()
	err = gold.GoldByDate("2020-11-12")
	if err != nil {
		fmt.Println(err)
	}
	endTime = time.Now()
	fmt.Println("Data from in-memory cache:", endTime.Sub(startTime))

	fmt.Println(gold.CreatePrettyOutput("en"))
	fmt.Println()

	nbpapi.DisableCache()
}
