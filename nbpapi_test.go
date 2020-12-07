package nbpapi

import (
	"fmt"
	"os"
)

var useMock bool

func init() {
	// useMock - main switch for all tests

	result := os.Getenv("USEMOCK")
	if result == "1" {
		useMock = true
	} else if result == "0" {
		useMock = false
	} else {
		useMock = true // default
	}

	if useMock {
		fmt.Println("USEMOCK == TRUE")
	}

}
