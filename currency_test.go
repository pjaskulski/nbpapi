package nbpapi

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestGetCurrencyCurrent(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var currencyName string = "frank szwajcarski"

	littleDelay()
	client := NewCurrency(table)

	err := client.CurrencyByDate(currency, "current")
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	if client.Exchange.Table != table {
		t.Errorf("Type of table: want: %s, got: %s", table, client.Exchange.Table)
	}
	if client.Exchange.Currency != currencyName {
		t.Errorf("Currency name: want: %s, got: %s", currencyName, client.Exchange.Currency)
	}
	if len(client.Exchange.Rates) != 1 {
		t.Errorf("expected 1 exchange currency rate, received: %d", len(client.Exchange.Rates))
	}
}

func TestGetCurrencyCurrentXXX(t *testing.T) {
	var table string = "A"
	var currency string = "XXX" // invalid currency code

	littleDelay()

	client := NewCurrency(table)

	_, err := client.GetRateCurrent(currency)
	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetCurrencyDay(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-13" // Friday 13 Nov 2020, CHF = 4.1605

	littleDelay()
	client := NewCurrency(table)

	err := client.CurrencyByDate(currency, day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}

	if client.Exchange.Rates[0].EffectiveDate != day {
		t.Errorf("incorect effectiveDate, want: %s, got %s", day, client.Exchange.Rates[0].EffectiveDate)
	}

	if client.Exchange.Rates[0].Mid != 4.1605 {
		t.Errorf("incorrect data, the CHF exchange rate on 2020-11-13 was 4.1605, not %.4f", client.Exchange.Rates[0].Mid)
	}
}

func TestGetCurrencyDaySaturdayFailed(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-14" // Saturday - no table of exchange rates

	littleDelay()

	client := NewCurrency(table)

	err := client.CurrencyByDate(currency, day)
	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetCurrencyLast(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var last int = 5

	littleDelay()
	client := NewCurrency(table)

	err := client.CurrencyLast(currency, last)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if len(client.Exchange.Rates) != 5 {
		t.Errorf("want: %d, got: %d", last, len(client.Exchange.Rates))
	}
}

func TestGetCurrencyLastFailed(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var last int = 500 // nbp api real max = 255

	littleDelay()

	client := NewCurrency(table)

	err := client.CurrencyLast(currency, last)
	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetCurrencyRange(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-12:2020-11-13" // valid range, expected 2 currency exchange rates

	littleDelay()
	client := NewCurrency(table)

	result, err := client.GetRateByDate(currency, day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	var ratesCount int = len(result)

	if ratesCount != 2 {
		t.Errorf("expected number of exchange rates == 2, obtained %d", ratesCount)
	}
}

func TestGetCurrencyRangeFailed(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-12:2020-11-10" // invalid range of dates

	littleDelay()

	client := NewCurrency(table)

	_, err := client.GetRateByDate(currency, day)
	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetCurrencyToday(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	today := time.Now()
	var day string = today.Format("2006-01-02")

	littleDelay()
	client := NewCurrency(table)

	_, err := client.GetRateByDate(currency, day)
	if err == nil {
		_, err := client.GetRateToday(currency)
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

// query tests

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

func TestCurrencyCSVOutput(t *testing.T) {
	want := "TABLE,DATE,AVERAGE (PLN)"

	client := NewCurrency("A")
	err := client.CurrencyByDate("EUR", "current")
	if err != nil {
		t.Error(err)
	}
	output := client.GetCSVOutput("en")
	if output[:24] != want {
		t.Errorf("invalid csv output, expected header: %s, got: %s", want, output[:29])
	}
}

func TestCurrencyPrettyOutput(t *testing.T) {

	client := NewCurrency("A")
	err := client.CurrencyByDate("EUR", "current")
	if err != nil {
		t.Error(err)
	}
	output := client.GetPrettyOutput("en")

	if len(output) == 0 {
		t.Errorf("incorrect (empty) pretty output")
	}

	text := "Table type:"
	if !strings.Contains(output, text) {
		t.Errorf("incorrect pretty output")
	}
}

func TestCurrencyRaw(t *testing.T) {
	client := NewCurrency("C")

	err := client.CurrencyRaw("EUR", "2020-12-02", 0, "json")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}
}

func TestCurrencyToday(t *testing.T) {
	today := time.Now()
	var day string = today.Format("2006-01-02")
	var code string = "CHF"
	var table string = "C"

	client := NewCurrency(table)

	err := client.CurrencyByDate(code, day)
	if err == nil {
		err = client.CurrencyToday(code)
		if err != nil {
			t.Errorf("want: err == nil, got: err != nil")
		}
	}
}

func TestGetRateCurrent(t *testing.T) {
	client := NewCurrency("C")

	result, err := client.GetRateCurrent("CHF")
	if err != nil {
		t.Errorf("want: err == nil, got: err != nil")
	}

	if result.Ask == 0 && result.Bid == 0 {
		t.Errorf("want ask and bid != 0, got: %.4f and %.4f", result.Ask, result.Bid)
	}
}

func TestGetRateToday(t *testing.T) {
	today := time.Now()
	var day string = today.Format("2006-01-02")
	var code string = "CHF"
	var table string = "C"

	client := NewCurrency(table)

	_, err := client.GetRateByDate(code, day)
	if err == nil {
		result, err := client.GetRateToday(code)
		if err != nil {
			t.Errorf("want: err == nil, got: err != nil")
		}

		if result.Ask == 0 && result.Bid == 0 {
			t.Errorf("want ask and bid != 0, got: %.4f and %.4f", result.Ask, result.Bid)
		}
	}
}
