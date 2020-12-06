// auxiliary program functions

package nbpapi

import (
	"errors"
	"math/rand"
	"time"
)

var ErrInvalidCurrencyCode error = errors.New("Invalid currency code")

// littleDelay - delay function, so as not to bother the NBP server too much...
func littleDelay() {
	interval := randomInteger(400, 650)
	time.Sleep(time.Millisecond * time.Duration(interval))
}

// randomInteger func
func randomInteger(minValue, maxValue int) int {
	var result int

	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(maxValue-minValue+1) + minValue

	return result
}

// inSlice - function checks if the specified string is present in the specified slice
func inSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func checkCurrencyCode(tableType, code string) error {
	switch tableType {
	case "A":
		if !inSlice(CurrencyValuesA, code) {
			return ErrInvalidCurrencyCode
		}
	case "B":
		if !inSlice(CurrencyValuesB, code) {
			return ErrInvalidCurrencyCode
		}
	case "C":
		if !inSlice(CurrencyValuesC, code) {
			return ErrInvalidCurrencyCode
		}
	}
	return nil
}
