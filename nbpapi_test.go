package nbpapi

import (
	"fmt"
	"os"
)

var useMock bool

func init() {
	// useMock - main switch for all tests, check environment variable USEMOCK,
	// if not set, by default: useMock == true
	result := os.Getenv("USEMOCK")

	if result == "1" {
		useMock = true
	} else if result == "0" {
		useMock = false
	} else {
		useMock = true
	}

	// print the information if tests uses httpmock or real service
	if useMock {
		fmt.Println("USEMOCK == TRUE")
	} else {
		fmt.Println("USEMOCK == FALSE")
	}
}
