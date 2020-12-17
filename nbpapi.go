package nbpapi

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/patrickmn/go-cache"
)

// base addresses of the NBP API service
const (
	baseAddressTable    string = "http://api.nbp.pl/api/exchangerates"
	baseAddressCurrency string = "http://api.nbp.pl/api/exchangerates"
	baseAddressGold     string = "http://api.nbp.pl/api/cenyzlota"
)

// TableValues - list of table types
var TableValues = []string{"A", "B", "C"}

// CurrencyValuesA - list of supported currencies for table type A
var CurrencyValuesA = []string{"THB", "USD", "AUD", "HKD", "CAD", "NZD", "SGD", "EUR", "HUF", "CHF",
	"GBP", "UAH", "JPY", "CZK", "DKK", "ISK", "NOK", "SEK", "HRK", "RON",
	"BGN", "TRY", "ILS", "CLP", "PHP", "MXN", "ZAR", "BRL", "MYR", "RUB",
	"IDR", "INR", "KRW", "CNY", "XDR"}

// CurrencyValuesB - list of supported currencies for table type B
var CurrencyValuesB = []string{"MGA", "PAB", "ETB", "AFN", "VES", "BOB", "CRC", "SVC", "NIO", "GMD",
	"MKD", "DZD", "BHD", "IQD", "JOD", "KWD", "LYD", "RSD", "TND", "MAD",
	"AED", "STN", "BSD", "BBD", "BZD", "BND", "FJD", "GYD", "JMD", "LRD",
	"NAD", "SRD", "TTD", "XCD", "SBD", "ZWL", "VND", "AMD", "CVE", "AWG",
	"BIF", "XOF", "XAF", "XPF", "DJF", "GNF", "KMF", "CDF", "RWF", "EGP",
	"GIP", "LBP", "SSP", "SDG", "SYP", "GHS", "HTG", "PYG", "ANG", "PGK",
	"LAK", "MWK", "ZMW", "AOA", "MMK", "GEL", "MDL", "ALL", "HNL", "SLL",
	"SZL", "LSL", "AZN", "MZN", "NGN", "ERN", "TWD", "TMT", "MRU", "TOP",
	"MOP", "ARS", "DOP", "COP", "CUP", "UYU", "BWP", "GTQ", "IRR", "YER",
	"QAR", "OMR", "SAR", "KHR", "BYN", "LKR", "MVR", "MUR", "NPR", "PKR",
	"SCR", "PEN", "KGS", "TJS", "UZS", "KES", "SOS", "TZS", "UGX", "BDT",
	"WST", "KZT", "MNT", "VUV", "BAM"}

// CurrencyValuesC - list of supported currencies for table type C
var CurrencyValuesC = []string{"USD", "AUD", "CAD", "EUR", "HUF", "CHF", "GBP", "JPY", "CZK", "DKK", "NOK",
	"SEK", "XDR"}

// SetLang function (language setting for output functions)
func setLang(lang string) {
	if lang == "pl" {
		l = langTexts["pl"]
	} else if lang == "en" {
		l = langTexts["en"]
	}
}

/* fetchData - function that retrieves data from the NBP website
   and returns them in the form of JSON / XML (or error), based on
   the arguments provided:

   url - NBP web api address
   format - 'json' or 'xml'

   Optionally, if the cache is turned on, the function trying to get data
   from in-memory store.
*/
func fetchData(client *http.Client, url, format string) ([]byte, error) {
	if format == "json" {
		format = "application/json"
	} else if format == "xml" {
		format = "application/xml"
	}

	// search in-memory store
	if CacheOn {
		value, found := Memory.Get(url)
		if found {
			return value.([]byte), nil
		}
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", format)

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode >= 400 {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// save data to in-memory store
	if CacheOn {
		Memory.Set(url, data, cache.DefaultExpiration)
	}

	return data, nil
}
