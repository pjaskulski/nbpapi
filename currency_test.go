package nbpapi

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestGetCurrencyCurrent(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var currencyName string = "frank szwajcarski"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"A","currency":"frank szwajcarski","code":"CHF","rates":`
		mockResponse += `[{"no":"238/A/NBP/2020","effectiveDate":"2020-12-07","mid":4.1417}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/A/CHF/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `404 NotFound - Not Found - Brak danych`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/A/XXX/",
			httpmock.NewStringResponder(404, mockResponse))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"A","currency":"frank szwajcarski","code":"CHF","rates":`
		mockResponse += `[{"no":"222/A/NBP/2020","effectiveDate":"2020-11-13","mid":4.1605}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/A/CHF/"+day+"/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

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

func TestGetCurrencyDayShouldFailOfInvalidCurrencyCode(t *testing.T) {
	var table string = "A"
	var currency string = "AOA" // currency code valid for B table, not A
	var day string = "2020-11-13"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `404 NotFound - Not Found - Brak danych`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/"+day+"/",
			httpmock.NewStringResponder(404, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency(table)

	err := client.CurrencyByDate(currency, day)
	if err != ErrInvalidCurrencyCode {
		t.Errorf("expected: err == ErrInvalidCurrencyCode")
	}
}

func TestGetCurrencyDayTableC(t *testing.T) {
	var table string = "C"
	var currency string = "CHF"
	var day string = "2020-11-13" // Friday 13 Nov 2020, CHF ask (sell) = 4.1980

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"C","currency":"frank szwajcarski","code":"CHF","rates":`
		mockResponse += `[{"no":"222/C/NBP/2020","effectiveDate":"2020-11-13","bid":4.1148,"ask":4.1980}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/"+day+"/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency(table)

	err := client.CurrencyByDate(currency, day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}

	if client.ExchangeC.Rates[0].EffectiveDate != day {
		t.Errorf("incorect effectiveDate, want: %s, got %s", day, client.ExchangeC.Rates[0].EffectiveDate)
	}

	if client.ExchangeC.Rates[0].Ask != 4.1980 {
		t.Errorf("incorrect data, the CHF exchange rate on 2020-11-13 was 4.1980, not %.4f", client.ExchangeC.Rates[0].Ask)
	}
}

func TestGetCurrencyDaySaturdayFailed(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var day string = "2020-11-14" // Saturday - no table of exchange rates

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `404 NotFound - Not Found - Brak danych`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/"+day+"/",
			httpmock.NewStringResponder(404, mockResponse))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"A","currency":"frank szwajcarski","code":"CHF","rates":`
		mockResponse += `[{"no":"234/A/NBP/2020","effectiveDate":"2020-12-01","mid":4.1226},`
		mockResponse += `{"no":"235/A/NBP/2020","effectiveDate":"2020-12-02","mid":4.1112},`
		mockResponse += `{"no":"236/A/NBP/2020","effectiveDate":"2020-12-03","mid":4.1362},`
		mockResponse += `{"no":"237/A/NBP/2020","effectiveDate":"2020-12-04","mid":4.1241},`
		mockResponse += `{"no":"238/A/NBP/2020","effectiveDate":"2020-12-07","mid":4.1417}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/last/"+strconv.Itoa(last)+"/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency(table)

	err := client.CurrencyLast(currency, last)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if len(client.Exchange.Rates) != 5 {
		t.Errorf("want: %d, got: %d", last, len(client.Exchange.Rates))
	}
}

func TestGetCurrencyLastTableC(t *testing.T) {
	var table string = "C"
	var currency string = "CHF"
	var last int = 5

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"C","currency":"frank szwajcarski","code":"CHF","rates":`
		mockResponse += `[{"no":"234/C/NBP/2020","effectiveDate":"2020-12-01","bid":4.0842,"ask":4.1668},`
		mockResponse += `{"no":"235/C/NBP/2020","effectiveDate":"2020-12-02","bid":4.0724,"ask":4.1546},`
		mockResponse += `{"no":"236/C/NBP/2020","effectiveDate":"2020-12-03","bid":4.1025,"ask":4.1853},`
		mockResponse += `{"no":"237/C/NBP/2020","effectiveDate":"2020-12-04","bid":4.0879,"ask":4.1705},`
		mockResponse += `{"no":"238/C/NBP/2020","effectiveDate":"2020-12-07","bid":4.0937,"ask":4.1765}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/last/"+strconv.Itoa(last)+"/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency(table)

	err := client.CurrencyLast(currency, last)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if len(client.ExchangeC.Rates) != 5 {
		t.Errorf("want: %d, got: %d", last, len(client.ExchangeC.Rates))
	}
}

func TestGetCurrencyLastFailed(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	var last int = 500 // nbp api real max = 255

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `400 BadRequest - Przekroczony limit 255 wyników / Maximum size of 255 data series has been exceeded`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/last/"+strconv.Itoa(last)+"/",
			httpmock.NewStringResponder(400, mockResponse))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"A","currency":"frank szwajcarski","code":"CHF","rates":`
		mockResponse += `[{"no":"221/A/NBP/2020","effectiveDate":"2020-11-12","mid":4.1573},`
		mockResponse += `{"no":"222/A/NBP/2020","effectiveDate":"2020-11-13","mid":4.1605}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/2020-11-12/2020-11-13/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `400 BadRequest - Błędny zakres dat / Invalid date range`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/2020-11-12/2020-11-10/",
			httpmock.NewStringResponder(400, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency(table)

	_, err := client.GetRateByDate(currency, day)
	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetCurrencyToday(t *testing.T) {
	var table string = "A"
	var currency string = "CHF"
	day := time.Now().Format("2006-01-02")

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		var mockStatus int

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			mockResponse = `404 NotFound - Not Found - Brak danych`
			mockStatus = 404

		} else {
			mockResponse = `{"table":"A","currency":"frank szwajcarski","code":"CHF","rates":`
			mockResponse += `[{"no":"238/A/NBP/2020","effectiveDate":"2020-12-07","mid":4.1417}]}`
			mockStatus = 200
		}

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/today/",
			httpmock.NewStringResponder(mockStatus, mockResponse))

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+currency+"/"+day+"/",
			httpmock.NewStringResponder(mockStatus, mockResponse))

	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"A","currency":"euro","code":"EUR","rates":[{"no":"238/A/NBP/2020","effectiveDate":"2020-12-07","mid":4.4745}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/A/EUR/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency("A")
	err := client.CurrencyByDate("EUR", "current")
	if err != nil {
		t.Error(err)
	}
	output := client.CreateCSVOutput("en")
	if output[:24] != want {
		t.Errorf("invalid csv output, expected header: %s, got: %s", want, output[:29])
	}
}

func TestCurrencyPrettyOutput(t *testing.T) {

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"A","currency":"euro","code":"EUR","rates":[{"no":"238/A/NBP/2020","effectiveDate":"2020-12-07","mid":4.4745}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/A/EUR/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency("A")
	err := client.CurrencyByDate("EUR", "current")
	if err != nil {
		t.Error(err)
	}
	output := client.CreatePrettyOutput("en")

	if len(output) == 0 {
		t.Errorf("incorrect (empty) pretty output")
	}

	text := "Table type:"
	if !strings.Contains(output, text) {
		t.Errorf("incorrect pretty output")
	}
}

func TestCurrencyRaw(t *testing.T) {

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"C","currency":"euro","code":"EUR","rates":`
		mockResponse += `[{"no":"235/C/NBP/2020","effectiveDate":"2020-12-02","bid":4.4151,"ask":4.5043}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/C/EUR/2020-12-02/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency("C")

	err := client.CurrencyRaw("EUR", "2020-12-02", 0, "json")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}
}

func TestCurrencyRawXML(t *testing.T) {

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		mockResponse := `
<ExchangeRatesSeries xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<Table>C</Table>
<Currency>euro</Currency>
<Code>EUR</Code>
<Rates>
	<Rate>
		<No>235/C/NBP/2020</No>
		<EffectiveDate>2020-12-02</EffectiveDate>
		<Bid>4.4151</Bid>
		<Ask>4.5043</Ask>
	</Rate>
</Rates>
</ExchangeRatesSeries>`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/C/EUR/2020-12-02/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency("C")

	err := client.CurrencyRaw("EUR", "2020-12-02", 0, "xml")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}
	if !IsValidXML(string(client.result)) {
		t.Errorf("incorrect xml content was received")
	}
}

func TestCurrencyToday(t *testing.T) {
	day := time.Now().Format("2006-01-02")
	var code string = "CHF"
	var table string = "C"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		var mockStatus int

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			mockResponse = `404 NotFound - Not Found - Brak danych`
			mockStatus = 404
		} else {
			mockResponse = `{"table":"C","currency":"euro","code":"EUR","rates":`
			mockResponse += `[{"no":"235/C/NBP/2020","effectiveDate":"` + day + `","bid":4.4151,"ask":4.5043}]}`
			mockStatus = 200
		}

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+code+"/"+day+"/",
			httpmock.NewStringResponder(mockStatus, mockResponse))

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/"+table+"/"+code+"/today/",
			httpmock.NewStringResponder(mockStatus, mockResponse))

	} else {
		littleDelay()
	}

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
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"C","currency":"frank szwajcarski","code":"CHF","rates":`
		mockResponse += `[{"no":"238/C/NBP/2020","effectiveDate":"2020-12-07","bid":4.0937,"ask":4.1765}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/C/CHF/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency("C")

	result, err := client.GetRateCurrent("CHF")
	if err != nil {
		t.Errorf("want: err == nil, got: err != nil")
	}

	if result.Ask == 0 && result.Bid == 0 {
		t.Errorf("want ask and bid != 0, got: %.4f and %.4f", result.Ask, result.Bid)
	}
}

func TestGetRateCurrentB(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `{"table":"B","currency":"birr etiopski","code":"ETB","rates":`
		mockResponse += `[{"no":"048/B/NBP/2020","effectiveDate":"2020-12-02","mid":0.0967}]}`

		httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/B/ETB/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewCurrency("B")

	result, err := client.GetRateCurrent("ETB")
	if err != nil {
		t.Errorf("want: err == nil, got: err != nil")
	}

	if result.Mid == 0 {
		t.Errorf("want average price != 0, got: %.4f", result.Mid)
	}
}

func TestGetRateToday(t *testing.T) {
	day := time.Now().Format("2006-01-02")
	var code string = "CHF"
	var table string = "C"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			mockResponse = `404 NotFound - Not Found - Brak danych`

			httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/C/CHF/today/",
				httpmock.NewStringResponder(404, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/C/CHF/"+day+"/",
				httpmock.NewStringResponder(404, mockResponse))
		} else {
			mockResponse = `{"table":"C","currency":"frank szwajcarski","code":"CHF","rates":`
			mockResponse += `[{"no":"238/C/NBP/2020","effectiveDate":"2020-12-07","bid":4.0937,"ask":4.1765}]}`

			httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/C/CHF/today/",
				httpmock.NewStringResponder(200, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressCurrency+"/rates/C/CHF/"+day+"/",
				httpmock.NewStringResponder(200, mockResponse))
		}
	} else {
		littleDelay()
	}

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

func TestCurrencySetTableType(t *testing.T) {
	client := NewCurrency("B")
	if client.tableType != "B" {
		t.Errorf("want: B, got %s", client.tableType)
	}

	client.SetTableType("A")
	if client.tableType != "A" {
		t.Errorf("want: A, got %s", client.tableType)
	}
}
