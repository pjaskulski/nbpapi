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
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !strings.Contains(string(result), "\"table\":\"A\",\"currency\":\"frank szwajcarski\"") {
		t.Errorf("incorrect json content was received")
	}
}

func TestGetCurrencyCurrentXXX(t *testing.T) {
	var table string = "A"
	var currency string = "XXX" // niepoprawny kod waluty

	littleDelay()
	address := queryCurrencyCurrent(table, currency)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetCurrencyDay(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-13" // Friday 13 Nov 2020, CHF = 4.1605

	littleDelay()
	address := queryCurrencyDate(table, day, currency)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	// mało eleganckie ale skuteczne
	if !strings.Contains(string(result), "\"effectiveDate\":\"2020-11-13\",\"mid\":4.1605") {
		t.Errorf("incorrect json content was received, the CHF exchange rate on November 13, 2020 was 4.1605")
	}
}

func TestGetCurrencyDaySaturday(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-14" // Saturday - no table of exchange rates

	littleDelay()
	address := queryCurrencyDate(table, day, currency)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
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
		t.Errorf("expected: err == nil, received: err != nil")
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
		t.Errorf("expected: err != nil, received: err == nil")
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
		t.Errorf("expected: err == nil, received: err != nil")
	}

	var nbpCurrency exchangeCurrency
	err = json.Unmarshal(result, &nbpCurrency)
	if err != nil {
		log.Fatal(err)
	}
	var ratesCount int = len(nbpCurrency.Rates)
	if ratesCount != 2 {
		t.Errorf("expected number of exchange rates == 2, obtained %d", ratesCount)
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
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetCurrencyToday(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var address string
	today := time.Now()
	var day string = today.Format("2006-01-02")

	littleDelay()
	address = queryCurrencyDate(table, day, currency)
	_, err := getData(address, "json")
	if err == nil {
		address = queryCurrencyToday(table, currency)
		_, err := getData(address, "json")
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

func TestQueryCurrencyRange(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/rates/A/CHF/2020-11-12/2020-11-19/"

	got := queryCurrencyRange("A", "2020-11-12:2020-11-19", "CHF")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryCurrencyLast(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/rates/A/CHF/last/5/"

	got := queryCurrencyLast("A", "5", "CHF")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryCurrencyToday(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/rates/A/CHF/today/"

	got := queryCurrencyToday("A", "CHF")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryCurrencyCurrent(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/rates/A/CHF/"

	got := queryCurrencyCurrent("A", "CHF")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryCurrencyDate(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/rates/A/CHF/2020-11-12/"

	got := queryCurrencyDate("A", "2020-11-12", "CHF")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
