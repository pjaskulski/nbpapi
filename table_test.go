package nbpapi

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestGetTableCurrent(t *testing.T) {
	var table string = "A"

	littleDelay()
	address := queryTableCurrent(table)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("incorrect json content was received")
	}
}

func TestGetTableDay(t *testing.T) {
	var table string = "A"
	var day string = "2020-11-17"
	var tableNo string = "224/A/NBP/2020"

	littleDelay()
	address := queryTableDate(table, day)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("incorrect json content was received")
	}

	var nbpTables []ExchangeTable
	err = json.Unmarshal(result, &nbpTables)
	if err != nil {
		log.Fatal(err)
	}

	if nbpTables[0].Table != table {
		t.Errorf("invalid table type, expected: %s, received: %s", table, nbpTables[0].Table)
	}
	if nbpTables[0].No != tableNo {
		t.Errorf("invalid table number, expected: %s, received: %s", tableNo, nbpTables[0].No)
	}
	if nbpTables[0].EffectiveDate != day {
		t.Errorf("invalid publication date, expected: %s, received: %s", day, nbpTables[0].EffectiveDate)
	}
}

func TestGetTableRange(t *testing.T) {
	var table string = "A"
	var day string = "2020-11-16:2020-11-17"

	littleDelay()
	address := queryTableRange(table, day)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("incorrect json content was received")
	}

	var nbpTables []ExchangeTable
	err = json.Unmarshal(result, &nbpTables)
	if err != nil {
		log.Fatal(err)
	}

	if len(nbpTables) != 2 {
		t.Errorf("2 exchange rate tables were expected, received: %d", len(nbpTables))
	}

	if nbpTables[0].Table != table {
		t.Errorf("invalid table type, expected: %s, received: %s", table, nbpTables[0].Table)
	}

	if nbpTables[1].Table != table {
		t.Errorf("invalid table type, expected: %s, received: %s", table, nbpTables[1].Table)
	}
}

func TestGetTableLast(t *testing.T) {
	var table string = "A"
	var lastNo string = "5"

	littleDelay()
	address := queryTableLast(table, lastNo)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("incorrect json content was received")
	}

	var nbpTables []ExchangeTable
	err = json.Unmarshal(result, &nbpTables)
	if err != nil {
		log.Fatal(err)
	}

	if len(nbpTables) != 5 {
		t.Errorf("5 exchange rate tables were expected, received: %d", len(nbpTables))
	}
}

func TestGetTableToday(t *testing.T) {
	var table string = "A"
	var address string
	today := time.Now()
	var day string = today.Format("2006-01-02")

	littleDelay()
	address = queryTableDate(table, day)
	_, err := getData(address, "json")
	if err == nil {
		address = queryTableToday(table)
		_, err := getData(address, "json")
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

func TestGetTableTodayFailed(t *testing.T) {
	var table string = "D"

	littleDelay()
	address := queryTableToday(table)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("expected: err != nil, received: err != nil")
	}
}

func TestQueryTableDate(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/tables/A/2020-12-02/"

	got := queryTableDate("A", "2020-12-02")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryTableCurrent(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/tables/A/"

	got := queryTableCurrent("A")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryTableToday(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/tables/A/today/"

	got := queryTableToday("A")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryTableLast(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/tables/A/last/3/"

	got := queryTableLast("A", "3")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestQueryTableRange(t *testing.T) {
	want := "http://api.nbp.pl/api/exchangerates/tables/A/2020-12-01/2020-12-02/"

	got := queryTableRange("A", "2020-12-01:2020-12-02")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
