// auxiliary program functions

package nbpapi

import (
	"encoding/xml"
	"errors"
	"io"
	"math/rand"
	"strings"
	"time"
)

// errors for invalid code, table type
var (
	ErrInvalidCurrencyCode error = errors.New("Invalid currency code")
	ErrInvalidTableType    error = errors.New("Invalid table type, allowed values: A, B or C")
)

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

// checkCurrencyCode - function checks if currency code is allowed for specified table type
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

// checkTableType - function chcek if table type is allowed
func checkTableType(tableType string) error {
	if !inSlice(TableValues, tableType) {
		return ErrInvalidTableType
	}
	return nil
}

// IsValidXML - func from: https://stackoverflow.com/questions/53476012/how-to-validate-a-xml
func IsValidXML(input string) bool {
	decoder := xml.NewDecoder(strings.NewReader(input))
	for {
		err := decoder.Decode(new(interface{}))
		if err != nil {
			return err == io.EOF
		}
	}
}
