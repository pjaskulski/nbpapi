package nbpapi

import (
	"encoding/json"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestGetGoldCurrent(t *testing.T) {
	littleDelay()
	address := queryGoldCurrent()
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}
}

func TestGetGoldToday(t *testing.T) {
	today := time.Now()
	var address string
	var day string = today.Format("2006-01-02")

	littleDelay()
	address = queryGoldDay(day)
	_, err := getData(address, "json")
	if err == nil {
		address = queryGoldToday()
		_, err := getData(address, "json")
		if err != nil {
			t.Errorf("oczekiwano err == nil, otrzymano err != nil")
		}
	}
}

func TestGetGoldDay(t *testing.T) {
	var day string = "2020-11-17"
	var cena float64 = 229.03

	littleDelay()
	address := queryGoldDay(day)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}

	var nbpGold []GoldRate
	err = json.Unmarshal(result, &nbpGold)
	if err != nil {
		log.Fatal(err)
	}

	if nbpGold[0].Data != day {
		t.Errorf("niepoprawna data, oczekiwano %s, otrzymano %s", day, nbpGold[0].Data)
	}
	if nbpGold[0].Price != cena {
		t.Errorf("niepoprawny cena, oczekiwano %.4f, otrzymano %.4f", cena, nbpGold[0].Price)
	}
}

func TestGetGoldDayFailed(t *testing.T) {
	var day string = "2020-11-15" // brak notowań w tym dniu

	littleDelay()
	address := queryGoldDay(day)
	_, err := getData(address, "json")
	if err == nil {
		t.Errorf("oczekiwano err != nil, otrzymano err == nil")
	}
}

func TestGetGoldLast(t *testing.T) {
	var lastNo int = 5

	littleDelay()
	address := queryGoldLast(strconv.Itoa(lastNo))
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}

	var nbpGold []GoldRate
	err = json.Unmarshal(result, &nbpGold)
	if err != nil {
		log.Fatal(err)
	}

	if len(nbpGold) != lastNo {
		t.Errorf("oczekiwano %d tabel kursów, otrzymano %d", lastNo, len(nbpGold))
	}
}

func TestGetGoldRange(t *testing.T) {
	var day string = "2020-11-16:2020-11-17"

	littleDelay()
	address := queryGoldRange(day)
	result, err := getData(address, "json")
	if err != nil {
		t.Errorf("oczekiwano err == nil, otrzymano err != nil")
	}
	if !json.Valid(result) {
		t.Errorf("otrzymano niepoprawną zawartość json")
	}

	var nbpGold []GoldRate
	err = json.Unmarshal(result, &nbpGold)
	if err != nil {
		log.Fatal(err)
	}

	if len(nbpGold) != 2 {
		t.Errorf("oczekiwano cen złota z 2 notowań, otrzymano z %d", len(nbpGold))
	}

	if nbpGold[0].Data != day[0:10] {
		t.Errorf("niepoprawna data, oczekiwano %s, otrzymano %s", day[0:10], nbpGold[0].Data)
	}

	if nbpGold[1].Data != day[11:] {
		t.Errorf("niepoprawna data, oczekiwano %s, otrzymano %s", day[11:], nbpGold[1].Data)
	}
}
