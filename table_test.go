package nbpapi

import (
	"encoding/json"
	"strings"
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

	littleDelay()
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

func TestTableCCSVOutput(t *testing.T) {
	want := "TABLE,CODE,NAME,BUY (PLN),SELL (PLN)"

	littleDelay()
	client := NewTable("C")
	err := client.TableByDate("current")
	if err != nil {
		t.Error(err)
	}
	output := client.GetCSVOutput("en")
	header := output[:len(want)]
	if header != want {
		t.Errorf("invalid csv output, expected header: %s, got: %s", want, header)
	}
}

func TestTableRaw(t *testing.T) {
	littleDelay()
	client := NewTable("C")

	err := client.TableRaw("2020-12-02", 0, "json")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}
}

func TestGetTableRawOutput(t *testing.T) {
	littleDelay()
	client := NewTable("C")

	err := client.TableRaw("2020-12-02", 0, "json")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}

	output := client.GetRawOutput()
	if output == "" {
		t.Errorf("invalid (empty) raw output")
	}
}

func TestGetTableCurrent(t *testing.T) {

	littleDelay()
	client := NewTable("A")
	_, err := client.GetTableCurrent()
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableCurrentFailed(t *testing.T) {

	littleDelay()
	client := NewTable("C")
	_, err := client.GetTableCurrent()
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableCCurrent(t *testing.T) {

	littleDelay()
	client := NewTable("C")
	_, err := client.GetTableCCurrent()
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableCCurrentFailedBecauseOfWrongType(t *testing.T) {
	littleDelay()
	client := NewTable("A")

	_, err := client.GetTableCCurrent()
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableToday(t *testing.T) {
	day := time.Now().Format("2006-01-02")

	littleDelay()
	client := NewTable("A")

	_, err := client.GetTableByDate(day) // test if table for today exists
	if err == nil {
		_, err := client.GetTableToday() // if it works for ..ByDay, should works for ..Today
		if err != nil {
			t.Errorf("want: err == nil, got err != nil")
		}
	}
}

func TestGetTableTodayShouldFailedBecauseOfWrongType(t *testing.T) {
	littleDelay()
	client := NewTable("C")

	_, err := client.GetTableToday() // wrong func for C table type
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableTodayShouldFailedBecauseOfWeekend(t *testing.T) {
	day := time.Now().Format("2006-01-02")

	littleDelay()
	client := NewTable("A")

	_, err := client.GetTableByDate(day)
	if err != nil {
		_, err := client.GetTableToday()
		if err == nil {
			t.Errorf("want: err != nil, got err == nil")
		}
	}
}

func TestGetTableCToday(t *testing.T) {
	day := time.Now().Format("2006-01-02")

	littleDelay()
	client := NewTable("C")

	_, err := client.GetTableCByDate(day) // test if table for today exists
	if err == nil {
		_, err := client.GetTableCToday() // if it works for ..ByDay, should works for ..Today
		if err != nil {
			t.Errorf("want: err == nil, got err != nil")
		}
	}
}

func TestGetTableTodayCFailedBecauseOfWrongType(t *testing.T) {
	littleDelay()
	client := NewTable("A")

	_, err := client.GetTableCToday() // wrong func for A table type
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableCTodayShouldFailedBecauseOfWeekend(t *testing.T) {
	day := time.Now().Format("2006-01-02")

	littleDelay()
	client := NewTable("C")

	_, err := client.GetTableCByDate(day)
	if err != nil {
		_, err := client.GetTableCToday()
		if err == nil {
			t.Errorf("want: err != nil, got err == nil")
		}
	}
}

func TestGetTableByDate(t *testing.T) {
	day := "2020-11-12"

	littleDelay()
	client := NewTable("A")

	_, err := client.GetTableByDate(day)
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableByDateFailedBecauseOfWrongType(t *testing.T) {
	day := "2020-11-12"

	littleDelay()
	client := NewTable("C")

	_, err := client.GetTableByDate(day)
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableCByDate(t *testing.T) {
	day := "2020-11-12"

	littleDelay()
	client := NewTable("C")

	_, err := client.GetTableCByDate(day)
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableCByDateFailedBecauseOfWrongType(t *testing.T) {
	day := "2020-11-12"

	littleDelay()
	client := NewTable("A")

	_, err := client.GetTableCByDate(day)
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestTablePrettyOutput(t *testing.T) {

	littleDelay()
	client := NewTable("A")
	err := client.TableByDate("current")
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

	text = "Table number:"
	if !strings.Contains(output, text) {
		t.Errorf("incorrect pretty output")
	}
}

func TestTableCPrettyOutput(t *testing.T) {

	littleDelay()
	client := NewTable("C")
	err := client.TableByDate("current")
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

	text = "Table number:"
	if !strings.Contains(output, text) {
		t.Errorf("incorrect pretty output")
	}
}

func TestTableCByDate(t *testing.T) {
	var table string = "C"
	var day string = "2020-11-17"
	var tableNo string = "224/C/NBP/2020"

	littleDelay()
	client := NewTable(table)

	err := client.TableByDate(day)
	if err != nil {
		t.Errorf("expected: err == nil, received: err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}

	if client.ExchangeC[0].Table != table {
		t.Errorf("invalid table type, expected: %s, received: %s", table, client.ExchangeC[0].Table)
	}
	if client.ExchangeC[0].No != tableNo {
		t.Errorf("invalid table number, expected: %s, received: %s", tableNo, client.ExchangeC[0].No)
	}
	if client.ExchangeC[0].EffectiveDate != day {
		t.Errorf("invalid publication date, expected: %s, received: %s", day, client.ExchangeC[0].EffectiveDate)
	}
}

func TestTableLastC(t *testing.T) {
	var table string = "C"
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

	if len(client.ExchangeC) != 5 {
		t.Errorf("5 exchange rate tables were expected, received: %d", len(client.ExchangeC))
	}
}
