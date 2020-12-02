package nbpapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

/* getData - universal function that retrieves data from the NBP website
   and returns them in the form of JSON / XML (or error), based on
   the arguments provided:

   url - NBP web api address
   format - 'json' or 'xml'
*/
func getData(url string, format string) ([]byte, error) {
	if format == "json" {
		format = "application/json"
	} else if format == "xml" {
		format = "application/xml"
	}

	client := &http.Client{
		Timeout: time.Second * 10,
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

	return data, nil
}

// SetLang function (language setting for output functions)
func setLang(lang string) {
	if lang == "pl" {
		l = langTexts["pl"]
	} else if lang == "en" {
		l = langTexts["en"]
	}
}
