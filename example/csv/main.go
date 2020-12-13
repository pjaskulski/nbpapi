// examples of using the nbpapi module

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pjaskulski/nbpapi"
)

func main() {
	/*
		How to get a latest 5 exchange rates of currency CHF,
		table of type A (mid - average exchange rate), as CSV data
	*/
	var err error

	fmt.Println("Save last 5 (CHF) currency rates to file chf.txt...")

	nbpMid := nbpapi.NewCurrency("A")
	err = nbpMid.CurrencyLast("CHF", 5)
	if err != nil {
		fmt.Println(err)
	}

	// write CSV to file, english version
	f, err := os.Create("chf.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(nbpMid.CreateCSVOutput("en"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Print last 5 (CHF) currency rates...")

	// print as CSV, polish version
	fmt.Println(nbpMid.CreateCSVOutput("pl"))
	fmt.Println()
}
