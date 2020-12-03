package nbpapi

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTableByDateCurrent(t *testing.T) {
	var table string = "A"

	littleDelay()

	client := NewTable(table)

	err := client.TableByDate("current")
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}
}

func TestTableByDate(t *testing.T) {
	var table string = "A"
	var day string = "2020-11-17"
	var tableNo string = "224/A/NBP/2020"

	littleDelay()

	client := NewTable(table)

	err := client.TableByDate(day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}

	if client.Exchange[0].Table != table {
		t.Errorf("invalid table type, expected: %s, received: %s", table, client.Exchange[0].Table)
	}
	if client.Exchange[0].No != tableNo {
		t.Errorf("invalid table number, expected: %s, received: %s", tableNo, client.Exchange[0].No)
	}
	if client.Exchange[0].EffectiveDate != day {
		t.Errorf("invalid publication date, expected: %s, received: %s", day, client.Exchange[0].EffectiveDate)
	}
}

func TestTableByDateRange(t *testing.T) {
	var table string = "A"
	var day string = "2020-11-16:2020-11-17"

	littleDelay()

	client := NewTable(table)

	err := client.TableByDate(day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}

	if len(client.Exchange) != 2 {
		t.Errorf("2 exchange rate tables were expected, received: %d", len(client.Exchange))
	}

	if client.Exchange[0].Table != table {
		t.Errorf("invalid table type, expected: %s, received: %s", table, client.Exchange[0].Table)
	}

	if client.Exchange[1].Table != table {
		t.Errorf("invalid table type, expected: %s, received: %s", table, client.Exchange[1].Table)
	}
}

func TestTableLast(t *testing.T) {
	var table string = "A"
	var lastNo int = 5

	littleDelay()
	client := NewTable(table)
	err := client.TableLast(lastNo)

	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}

	if len(client.Exchange) != 5 {
		t.Errorf("5 exchange rate tables were expected, received: %d", len(client.Exchange))
	}
}

func TestTableByDateToday(t *testing.T) {
	var table string = "A"
	today := time.Now()
	var day string = today.Format("2006-01-02")

	littleDelay()
	client := NewTable(table)

	err := client.TableByDate(day)
	if err == nil {
		err := client.TableByDate("today")
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

func TestTableByDateFailed(t *testing.T) {
	var table string = "D"

	littleDelay()

	client := NewTable(table)

	err := client.TableByDate("today")
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

func TestTableCSVOutput(t *testing.T) {
	want := "TABLE,CODE,NAME,AVERAGE (PLN)"

	client := NewTable("A")
	err := client.TableByDate("current")
	if err != nil {
		t.Error(err)
	}
	output := client.GetCSVOutput("en")
	if output[:29] != want {
		t.Errorf("invalid csv output, expected header: %s, got: %s", want, output[:29])
	}
}

func TestGetTableCurrent(t *testing.T) {
	client := NewTable("A")
	_, err := client.GetTableCurrent()
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableCCurrent(t *testing.T) {
	client := NewTable("C")
	_, err := client.GetTableCCurrent()
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}
