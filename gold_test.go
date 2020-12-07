package nbpapi

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestGetGoldCurrent(t *testing.T) {

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		today := time.Now().Format("2006-01-01")
		httpmock.RegisterResponder("GET", baseAddressGold,
			httpmock.NewStringResponder(200, `[{"data":"`+today+`","cena":717.83}]`))
	} else {
		littleDelay() // delay if use real NBP API service
	}

	apiClient := NewGold()
	result, err := apiClient.GetPriceCurrent()

	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	if !json.Valid(apiClient.result) {
		t.Errorf("incorrect json content was received")
	}

	if result.Price <= 0 {
		t.Errorf("incorrect price of gold was received")
	}
}

func TestGetGoldToday(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	var err error

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			httpmock.RegisterResponder("GET", baseAddressGold+"/today",
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))

			httpmock.RegisterResponder("GET", baseAddressGold+"/"+today,
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))
		} else {
			httpmock.RegisterResponder("GET", baseAddressGold+"/today",
				httpmock.NewStringResponder(200, `[{"data":"`+today+`","cena":717.83}]`))
			httpmock.RegisterResponder("GET", baseAddressGold+"/"+today,
				httpmock.NewStringResponder(200, `[{"data":"`+today+`","cena":717.83}]`))
		}
	} else {
		littleDelay() // delay if use real NBP API service
	}

	apiClient := NewGold()
	_, err = apiClient.GetPriceByDate(today)
	if err == nil {
		_, err := apiClient.GetPriceToday()
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

func TestGetGoldTodayFailedBecaueOfWeekend(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	var err error

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			httpmock.RegisterResponder("GET", baseAddressGold+"/today",
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))

			httpmock.RegisterResponder("GET", baseAddressGold+"/"+today,
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))
		} else {
			httpmock.RegisterResponder("GET", baseAddressGold+"/today",
				httpmock.NewStringResponder(200, `[{"data":"`+today+`","cena":717.83}]`))
			httpmock.RegisterResponder("GET", baseAddressGold+"/"+today,
				httpmock.NewStringResponder(200, `[{"data":"`+today+`","cena":717.83}]`))
		}
	} else {
		littleDelay() // delay if use real NBP API service
	}

	apiClient := NewGold()

	_, err = apiClient.GetPriceByDate(today)
	if err != nil {
		_, err := apiClient.GetPriceToday()
		if err == nil {
			t.Errorf("expected: err != nil, received: err == nil")
		}
	}
}

func TestGetGoldDay(t *testing.T) {
	var day string = "2020-11-17"
	var cena float64 = 229.03

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressGold+"/"+day,
			httpmock.NewStringResponder(200, `[{"data":"`+day+`","cena":229.03}]`))
	} else {
		littleDelay()
	}

	apiClient := NewGold()

	err := apiClient.GoldByDate(day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(apiClient.result) {
		t.Errorf("incorrect json content was received")
	}

	if apiClient.GoldRates[0].Data != day {
		t.Errorf("invalid date, %s expected, %s received", day, apiClient.GoldRates[0].Data)
	}
	if apiClient.GoldRates[0].Price != cena {
		t.Errorf("invalid price, expected %.4f, received %.4f", cena, apiClient.GoldRates[0].Price)
	}
}

func TestGetGoldDayShouldFailedOnSunday(t *testing.T) {
	var day string = "2020-11-15" // no data on this day (Sunday)

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressGold+"/"+day,
			httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))
	} else {
		littleDelay()
	}

	apiClient := NewGold()
	_, err := apiClient.GetPriceByDate(day)

	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetGoldLast(t *testing.T) {
	var lastNo int = 5

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `[{"data":"2020-12-01","cena":211.74},{"data":"2020-12-02","cena":217.55},`
		mockResponse += `{"data":"2020-12-03","cena":217.04},{"data":"2020-12-04","cena":217.86},`
		mockResponse += `{"data":"2020-12-07","cena":217.83}]`

		httpmock.RegisterResponder("GET", baseAddressGold+"/last/"+strconv.Itoa(lastNo),
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	apiClient := NewGold()
	err := apiClient.GoldLast(lastNo)

	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(apiClient.result) {
		t.Errorf("incorrect json content was received")
	}

	if len(apiClient.GoldRates) != lastNo {
		t.Errorf("expected: %d exchange rate tables, received: %d", lastNo, len(apiClient.GoldRates))
	}
}

func TestGetGoldRange(t *testing.T) {
	var day string = "2020-11-16:2020-11-17"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `[{"data":"2020-11-16","cena":231.09},{"data":"2020-11-17","cena":229.03}]`

		httpmock.RegisterResponder("GET", baseAddressGold+"/"+strings.Replace(day, ":", "/", 1),
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	apiClient := NewGold()

	err := apiClient.GoldByDate(day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	if !json.Valid(apiClient.result) {
		t.Errorf("incorrect json content was received")
	}

	if len(apiClient.GoldRates) != 2 {
		t.Errorf("gold prices were expected from 2 quotes, obtained from %d", len(apiClient.GoldRates))
	}

	if apiClient.GoldRates[0].Data != day[0:10] {
		t.Errorf("invalid start date, %s expected, %s received", day[0:10], apiClient.GoldRates[0].Data)
	}

	if apiClient.GoldRates[1].Data != day[11:] {
		t.Errorf("invalid stop date, %s expected, %s received", day[11:], apiClient.GoldRates[1].Data)
	}
}

func TestShouldGetGoldByDateWithSuccess(t *testing.T) {
	var result []GoldRate
	date := "2020-12-01"
	expected := 211.7400 // the real price of gold on December 12, 2020 was PLN 211.7400

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `[{"data":"2020-12-01","cena":211.74}]`

		httpmock.RegisterResponder("GET", baseAddressGold+"/"+date,
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	apiClient := NewGold()
	result, err := apiClient.GetPriceByDate(date)
	if err != nil {
		t.Error(err)
	}

	if result[0].Data != date {
		t.Errorf("invalid date, %s expected, %s recived", date, result[0].Data)
	}
	if result[0].Price != expected {
		t.Errorf("invalid price, %.4f expected, %.4f received", expected, result[0].Price)
	}
}

func TestShouldGetGoldCurrentWithSuccess(t *testing.T) {
	var result []GoldRate
	date := "current"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `[{"data":"2020-12-01","cena":211.74}]`

		httpmock.RegisterResponder("GET", baseAddressGold,
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	apiClient := NewGold()
	result, err := apiClient.GetPriceByDate(date)
	if err != nil {
		t.Error(err)
	}

	if !(result[0].Price > 0) {
		t.Errorf("invalid current price, expected >0, %.4f received", result[0].Price)
	}
}

func TestGetPriceToday(t *testing.T) {
	var err error
	today := time.Now().Format("2006-01-02")

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			httpmock.RegisterResponder("GET", baseAddressGold+"/today",
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))

			httpmock.RegisterResponder("GET", baseAddressGold+"/"+today,
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))
		} else {
			httpmock.RegisterResponder("GET", baseAddressGold+"/today",
				httpmock.NewStringResponder(200, `[{"data":"`+today+`","cena":717.83}]`))
			httpmock.RegisterResponder("GET", baseAddressGold+"/"+today,
				httpmock.NewStringResponder(200, `[{"data":"`+today+`","cena":717.83}]`))
		}
	} else {
		littleDelay()
	}

	apiClient := NewGold()
	_, err = apiClient.GetPriceByDate(today)
	if err == nil {
		_, err := apiClient.GetPriceToday()
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

func TestGetPriceCurrentShouldReturnNonZeroPrice(t *testing.T) {
	var err error
	var result GoldRate

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = `[{"data":"2020-12-01","cena":211.74}]`

		httpmock.RegisterResponder("GET", baseAddressGold,
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	apiClient := NewGold()
	result, err = apiClient.GetPriceCurrent()

	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	if result.Price <= 0 {
		t.Errorf("incorrect price of gold was received")
	}
}

func TestQueryGoldToday(t *testing.T) {
	var want string = "http://api.nbp.pl/api/cenyzlota/today"

	got := queryGoldToday()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryGoldCurrent(t *testing.T) {
	var want string = "http://api.nbp.pl/api/cenyzlota"

	got := queryGoldCurrent()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryGoldLast(t *testing.T) {
	var want string = "http://api.nbp.pl/api/cenyzlota/last/5"

	got := queryGoldLast("5")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryGoldDate(t *testing.T) {
	var want string = "http://api.nbp.pl/api/cenyzlota/2020-11-12"

	got := queryGoldDate("2020-11-12")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryGoldRange(t *testing.T) {
	want := "http://api.nbp.pl/api/cenyzlota/2020-11-12/2020-11-19"

	got := queryGoldRange("2020-11-12:2020-11-19")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetGoldAddress(t *testing.T) {
	apiClient := NewGold()

	want := "http://api.nbp.pl/api/cenyzlota/2020-11-12/2020-11-19"

	got := apiClient.getGoldAddress("2020-11-12:2020-11-19", 0)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}

	want = "http://api.nbp.pl/api/cenyzlota/2020-11-12"
	got = apiClient.getGoldAddress("2020-11-12", 0)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}

	want = "http://api.nbp.pl/api/cenyzlota/last/5"
	got = apiClient.getGoldAddress("", 5)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}

	want = "http://api.nbp.pl/api/cenyzlota"
	got = apiClient.getGoldAddress("current", 0)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}

	want = "http://api.nbp.pl/api/cenyzlota/today"
	got = apiClient.getGoldAddress("today", 0)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGoldGetCSVOutput(t *testing.T) {
	var err error
	var result string

	littleDelay()

	apiClient := NewGold()
	err = apiClient.GoldByDate("current")

	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	result = apiClient.GetCSVOutput("en")

	if len(result) == 0 {
		t.Errorf("incorrect (empty) CSV output")
	}

	if result[:16] != "DATE,PRICE (PLN)" {
		t.Errorf("incorrect CSV output")
	}
}

func TestGoldGetPrettyOutput(t *testing.T) {
	var err error
	var result string

	littleDelay()

	apiClient := NewGold()
	err = apiClient.GoldByDate("current")

	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}

	result = apiClient.GetPrettyOutput("en")

	if len(result) == 0 {
		t.Errorf("incorrect (empty) pretty output")
	}

	text := "The price of 1g of gold (of 1000 millesimal fineness)"
	if !strings.Contains(result, text) {
		t.Errorf("incorrect pretty output")
	}
}

// low level getData tests
func TestGoldGetDataFailed(t *testing.T) {
	type args struct {
		url    string
		format string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid url",
			args: args{
				url:    "http://api.nbp.pl/api/cenyzlotx",
				format: "json",
			},
			wantErr: true,
		},
		{
			name: "Invalid range od dates",
			args: args{
				url:    "http://api.nbp.pl/api/cenyzlota/2020-11-12/2020-11-10",
				format: "json",
			},
			wantErr: true,
		},
		{
			name: "Invalid date",
			args: args{
				url:    "http://api.nbp.pl/api/cenyzlota/2020-11-29",
				format: "json",
			},
			wantErr: true,
		},
	}

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "http://api.nbp.pl/api/cenyzlotx",
			httpmock.NewStringResponder(404, `404 NotFound`))

		httpmock.RegisterResponder("GET", "http://api.nbp.pl/api/cenyzlota/2020-11-12/2020-11-10",
			httpmock.NewStringResponder(400, `400 BadRequest - Błędny zakres dat / Invalid date range`))

		httpmock.RegisterResponder("GET", "http://api.nbp.pl/api/cenyzlota/2020-11-29",
			httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))
	}

	client := NewGold()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.getData(tt.args.url, tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("getData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGoldRaw(t *testing.T) {
	client := NewGold()

	err := client.GoldRaw("2020-12-02", 0, "json")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}
}

func TestGoldRawOutput(t *testing.T) {
	client := NewGold()

	err := client.GoldLast(1)
	if err != nil {
		t.Error(err)
	}

	output := client.GetRawOutput()
	if output == "" {
		t.Errorf("invalid (empty) raw output")
	}
}
