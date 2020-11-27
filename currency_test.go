package nbpapi

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
	"time"
)

func TestGetCurrencyCurrent(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"

	littleDelay()
	address := queryCurrencyCurrent(table, currency)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !strings.Contains(string(result), "\"table\":\"A\",\"currency\":\"frank szwajcarski\"") {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}
}

func TestGetCurrencyCurrentXXX(t *testing.T) {
	var table string = "A"
	var currency string = "XXX" // niepoprawny kod waluty

	littleDelay()
	address := queryCurrencyCurrent(table, currency)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("oczekiwano err != nil, otrzymano err == nil")
	}
}

func TestGetCurrencyDay(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-13" // Friday - ok, kurs CHF = 4.1605

	littleDelay()
	address := queryCurrencyDay(table, day, currency)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	// mało eleganckie ale skuteczne
	if !strings.Contains(string(result), "\"effectiveDate\":\"2020-11-13\",\"mid\":4.1605") {
		t.Errorf("niepoprawna zawartość json, kurs CHF 13.11.2020 wynosił 4.1605")
	}
}

func TestGetCurrencyDaySaturday(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-14" // Saturday - no table of exchange rates

	littleDelay()
	address := queryCurrencyDay(table, day, currency)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("oczekiwano err != nil, otrzymano err == nil")
	}
}

func TestGetCurrencyLast(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var last string = "5"

	littleDelay()
	address := queryCurrencyLast(table, last, currency)
	_, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
}

func TestGetCurrencyLastFailed(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var last string = "500" // za dużo kursów, max = 255

	littleDelay()
	address := queryCurrencyLast(table, last, currency)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("oczekiwano err != nil, otrzymano err == nil")
	}
}

func TestGetCurrencyRange(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-12:2020-11-13" // poprawny zakres dat, spodziewane 2 kursy

	littleDelay()
	address := queryCurrencyRange(table, day, currency)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}

	var nbpCurrency exchangeCurrency
	err = json.Unmarshal(result, &nbpCurrency)
	if err != nil {
		log.Fatal(err)
	}
	var ratesCount int = len(nbpCurrency.Rates)
	if ratesCount != 2 {
		t.Errorf("oczekiwana liczba kursów == 2, otrzymano %d", ratesCount)
	}
}

func TestGetCurrencyRangeFailed(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-12:2020-11-10" // niepoprawny zakres dat

	littleDelay()
	address := queryCurrencyRange(table, day, currency)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("oczekiwano err != nil, otrzymano err == nil")
	}
}

func TestGetCurrencyToday(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var address string
	today := time.Now()
	var day string = today.Format("2006-01-02")

	littleDelay()
	address = queryCurrencyDay(table, day, currency)
	_, err := getData(address, "json")
	if err == nil {
		address = queryCurrencyToday(table, currency)
		_, err := getData(address, "json")
		if err != nil {
			t.Errorf("oczekiwano err == nil, otrzymano err != nil")
		}
	}
}
