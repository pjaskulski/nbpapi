// auxiliary program functions

package nbpapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// getData - universal function that returns JSON/XML (or error)
// based on the address provided
func getData(address string, format string) ([]byte, error) {
	if format == "json" {
		format = "application/json"
	} else if format == "xml" {
		format = "application/xml"
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", address, nil)
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

	return data, nil
}

// littleDelay - delay function, so as not to bother the NBP server too much...
func littleDelay() {
	time.Sleep(time.Millisecond * 500)
}
