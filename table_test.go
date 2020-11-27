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
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}
}

func TestGetTableDay(t *testing.T) {
	var table string = "A"
	var day string = "2020-11-17"
	var tableNo string = "224/A/NBP/2020"

	littleDelay()
	address := queryTableDay(table, day)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}

	var nbpTables []ExchangeTable
	err = json.Unmarshal(result, &nbpTables)
	if err != nil {
		log.Fatal(err)
	}

	if nbpTables[0].Table != table {
		t.Errorf("niepoprawny typ tabeli, oczekiwano %s, otrzymano %s", table, nbpTables[0].Table)
	}
	if nbpTables[0].No != tableNo {
		t.Errorf("niepoprawny numer tabeli, oczekiwano %s, otrzymano %s", tableNo, nbpTables[0].No)
	}
	if nbpTables[0].EffectiveDate != day {
		t.Errorf("niepoprawna data publikacji, oczekiwano %s, otrzymano %s", day, nbpTables[0].EffectiveDate)
	}
}

func TestGetTableRange(t *testing.T) {
	var table string = "A"
	var day string = "2020-11-16:2020-11-17"

	littleDelay()
	address := queryTableRange(table, day)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}

	var nbpTables []ExchangeTable
	err = json.Unmarshal(result, &nbpTables)
	if err != nil {
		log.Fatal(err)
	}

	if len(nbpTables) != 2 {
		t.Errorf("oczekiwano 2 tabel kursów, otrzymano %d", len(nbpTables))
	}

	if nbpTables[0].Table != table {
		t.Errorf("niepoprawny typ tabeli, oczekiwano %s, otrzymano %s", table, nbpTables[0].Table)
	}

	if nbpTables[1].Table != table {
		t.Errorf("niepoprawny typ tabeli, oczekiwano %s, otrzymano %s", table, nbpTables[1].Table)
	}
}

func TestGetTableLast(t *testing.T) {
	var table string = "A"
	var lastNo string = "5"

	littleDelay()
	address := queryTableLast(table, lastNo)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}

	var nbpTables []ExchangeTable
	err = json.Unmarshal(result, &nbpTables)
	if err != nil {
		log.Fatal(err)
	}

	if len(nbpTables) != 5 {
		t.Errorf("oczekiwano 5 tabel kursów, otrzymano %d", len(nbpTables))
	}
}

func TestGetTableToday(t *testing.T) {
	var table string = "A"
	var address string
	today := time.Now()
	var day string = today.Format("2006-01-02")

	littleDelay()
	address = queryTableDay(table, day)
	_, err := getData(address, "json")
	if err == nil {
		address = queryTableToday(table)
		_, err := getData(address, "json")
		if err != nil {
			t.Errorf("oczekiwano err == nil, otrzymano err != nil")
		}
	}
}

func TestGetTableTodayFailed(t *testing.T) {
	var table string = "D"

	littleDelay()
	address := queryTableToday(table)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("oczekiwano err != nil, otrzymano err != nil")
	}
}
