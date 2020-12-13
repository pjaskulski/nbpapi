// examples of using the nbpapi module

package main

import (
	"fmt"

	"github.com/pjaskulski/nbpapi"
)

func main() {
	// get current table A
	nbpTable := nbpapi.NewTable("A")

	err := nbpTable.TableByDate("current")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nbpTable.CreatePrettyOutput("en"))
	}
}
