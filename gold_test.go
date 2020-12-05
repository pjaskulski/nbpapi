package nbpapi

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

// low level getData tests
func TestGetGoldCurrent(t *testing.T) {
	littleDelay()

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
	today := time.Now()
	var day string = today.Format("2006-01-02")
	var err error

	littleDelay()

	apiClient := NewGold()

	_, err = apiClient.GetPriceByDate(day)
	if err == nil {
		_, err := apiClient.GetPriceToday()
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

func TestGetGoldTodayFailedBecaueOfWeekend(t *testing.T) {
	day := time.Now().Format("2006-01-02")
	var err error

	littleDelay()

	apiClient := NewGold()

	_, err = apiClient.GetPriceByDate(day)
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

	littleDelay()

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

func TestGetGoldDayFailed(t *testing.T) {
	var day string = "2020-11-15" // brak notowań w tym dniu

	littleDelay()

	apiClient := NewGold()
	_, err := apiClient.GetPriceByDate(day)

	if err == nil {
		t.Errorf("expected: err != nil, received: err == nil")
	}
}

func TestGetGoldLast(t *testing.T) {
	var lastNo int = 5

	littleDelay()

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

	littleDelay()

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
		t.Errorf("invalid date, %s expected, %s received", day[0:10], apiClient.GoldRates[0].Data)
	}

	if apiClient.GoldRates[1].Data != day[11:] {
		t.Errorf("invalid date, %s expected, %s received", day[11:], apiClient.GoldRates[0].Data)
	}
}

// NBPGold methods test

func TestShouldGetGoldByDateSuccess(t *testing.T) {
	var result []GoldRate
	date := "2020-12-01"
	expected := 211.7400 // the real price of gold on December 12, 2020 was PLN 211.7400

	littleDelay()

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

func TestShouldGetGoldCurrentSuccess(t *testing.T) {
	var result []GoldRate
	date := "current"

	littleDelay()

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

	today := time.Now()
	var date string = today.Format("2006-01-02")

	littleDelay()

	apiClient := NewGold()
	_, err = apiClient.GetPriceByDate(date)
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

	littleDelay()

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
