package nbpapi

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestTableByDateCurrent(t *testing.T) {
	var table string = "A"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/",
			httpmock.NewStringResponder(200, mockTableA))
	} else {
		littleDelay()
	}

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
	var day string = "2020-12-07"
	var tableNo string = "238/A/NBP/2020"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/2020-12-07/",
			httpmock.NewStringResponder(200, mockTableA))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/2020-11-16/2020-11-17/",
			httpmock.NewStringResponder(200, mockRangeOfTablesA))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/last/5/",
			httpmock.NewStringResponder(200, mockTableLast5))
	} else {
		littleDelay()
	}

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
	today := time.Now().Format("2006-01-02")

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			mockResponse = `404 NotFound - Not Found - Brak danych`

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/"+today+"/",
				httpmock.NewStringResponder(404, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/today/",
				httpmock.NewStringResponder(404, mockResponse))

		} else {
			mockResponse = mockTableA // data published on 2020-12-07, but this does not matter in this test

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/"+today+"/",
				httpmock.NewStringResponder(200, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/today/",
				httpmock.NewStringResponder(200, mockResponse))
		}
	} else {
		littleDelay()
	}

	client := NewTable(table)

	err := client.TableByDate(today)
	if err == nil {
		err := client.TableByDate("today")
		if err != nil {
			t.Errorf("expected: err == nil, received: err != nil")
		}
	}
}

func TestTableByDateFailed(t *testing.T) {
	var table string = "D"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		mockResponse := `400 BadRequest - Bad Request - Nieprawidłowa wartość parametru: {table}='D'`
		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/D/today/",
			httpmock.NewStringResponder(400, mockResponse))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/",
			httpmock.NewStringResponder(200, mockTableA))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

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
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/2020-12-07/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

	client := NewTable("C")

	err := client.TableRaw("2020-12-07", 0, "json")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}
	if !json.Valid(client.result) {
		t.Errorf("incorrect json content was received")
	}
}

func TestTableRawXML(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/2020-12-07/",
			httpmock.NewStringResponder(200, mockTableCXML))
	} else {
		littleDelay()
	}

	client := NewTable("C")

	err := client.TableRaw("2020-12-07", 0, "xml")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}
	if !IsValidXML(string(client.result)) {
		t.Errorf("incorrect xml content was received")
	}
}

func TestGetTableRawOutput(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/2020-12-07/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

	client := NewTable("C")

	err := client.TableRaw("2020-12-07", 0, "json")
	if err != nil {
		t.Errorf("want err == nil, got err != nil")
	}

	output := client.GetRawOutput()
	if output == "" {
		t.Errorf("invalid (empty) raw output")
	}
}

func TestGetTableCurrent(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = mockTableA
		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewTable("A")
	_, err := client.GetTableCurrent()
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableCurrentFailed(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = mockTableC
		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewTable("C")
	_, err := client.GetTableCurrent()
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableCCurrent(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = mockTableC
		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewTable("C")
	_, err := client.GetTableCCurrent()
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableCCurrentFailedBecauseOfWrongType(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string
		mockResponse = mockTableA
		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/",
			httpmock.NewStringResponder(200, mockResponse))
	} else {
		littleDelay()
	}

	client := NewTable("A")

	_, err := client.GetTableCCurrent()
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableToday(t *testing.T) {
	today := time.Now().Format("2006-01-02")

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		var mockResponse string

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			mockResponse = `404 NotFound - Not Found - Brak danych`

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/",
				httpmock.NewStringResponder(404, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/today/",
				httpmock.NewStringResponder(404, mockResponse))
		} else {
			mockResponse = mockTableA
			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/",
				httpmock.NewStringResponder(200, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/today/",
				httpmock.NewStringResponder(200, mockResponse))
		}
	} else {
		littleDelay()
	}

	client := NewTable("A")

	_, err := client.GetTableByDate(today) // test if table for today exists
	if err == nil {
		_, err := client.GetTableToday() // if it works for ..ByDay, should works for ..Today
		if err != nil {
			t.Errorf("want: err == nil, got err != nil")
		}
	}
}

func TestGetTableTodayShouldFailedBecauseOfWrongType(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/today/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

	client := NewTable("C")

	_, err := client.GetTableToday() // wrong func for C table type, should be GetTableCToday
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableTodayShouldFailedBecauseOfWeekend(t *testing.T) {
	today := time.Now().Format("2006-01-02")

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			mockResponse := `404 NotFound - Not Found - Brak danych`

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/"+today+"/",
				httpmock.NewStringResponder(404, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/today/",
				httpmock.NewStringResponder(404, mockResponse))
		} else {
			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/"+today+"/",
				httpmock.NewStringResponder(200, mockTableA))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/today/",
				httpmock.NewStringResponder(200, mockTableA))
		}
	} else {
		littleDelay()
	}

	client := NewTable("A")

	_, err := client.GetTableByDate(today)
	if err != nil {
		_, err := client.GetTableToday()
		if err == nil {
			t.Errorf("want: err != nil, got err == nil")
		}
	}
}

func TestGetTableCToday(t *testing.T) {
	today := time.Now().Format("2006-01-02")

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			mockResponse := `404 NotFound - Not Found - Brak danych`

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/"+today+"/",
				httpmock.NewStringResponder(404, mockResponse))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/today/",
				httpmock.NewStringResponder(404, mockResponse))
		} else {
			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/"+today+"/",
				httpmock.NewStringResponder(200, mockTableC))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/today/",
				httpmock.NewStringResponder(200, mockTableC))
		}
	} else {
		littleDelay()
	}

	client := NewTable("C")

	_, err := client.GetTableCByDate(today) // test if table for today exists
	if err == nil {
		_, err := client.GetTableCToday() // if it works for ..ByDay, should works for ..Today
		if err != nil {
			t.Errorf("want: err == nil, got err != nil")
		}
	}
}

func TestGetTableTodayCFailedBecauseOfWrongType(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/today/",
			httpmock.NewStringResponder(200, mockTableA))
	} else {
		littleDelay()
	}

	client := NewTable("A")

	_, err := client.GetTableCToday() // wrong func for A table type
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableCTodayShouldFailedBecauseOfWeekend(t *testing.T) {
	today := time.Now().Format("2006-01-02")

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		weekday := time.Now().Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/today/",
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/"+today+"/",
				httpmock.NewStringResponder(404, `404 NotFound - Not Found - Brak danych`))
		} else {
			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/today/",
				httpmock.NewStringResponder(200, mockTableC))

			httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/"+today+"/",
				httpmock.NewStringResponder(200, mockTableC))
		}

	} else {
		littleDelay()
	}

	client := NewTable("C")

	_, err := client.GetTableCByDate(today)
	if err != nil {
		_, err := client.GetTableCToday()
		if err == nil {
			t.Errorf("want: err != nil, got err == nil")
		}
	}
}

func TestGetTableByDate(t *testing.T) {
	day := "2020-12-07"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/"+day+"/",
			httpmock.NewStringResponder(200, mockTableA))
	} else {
		littleDelay()
	}

	client := NewTable("A")

	_, err := client.GetTableByDate(day)
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableByDateFailedBecauseOfWrongType(t *testing.T) {
	day := "2020-12-07"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/"+day+"/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

	client := NewTable("C")

	_, err := client.GetTableByDate(day)
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestGetTableCByDate(t *testing.T) {
	day := "2020-12-07"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/"+day+"/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

	client := NewTable("C")

	_, err := client.GetTableCByDate(day)
	if err != nil {
		t.Errorf("want: err == nil, got err != nil")
	}
}

func TestGetTableCByDateFailedBecauseOfWrongType(t *testing.T) {
	day := "2020-12-07"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/"+day+"/",
			httpmock.NewStringResponder(200, mockTableA))
	} else {
		littleDelay()
	}

	client := NewTable("A")

	_, err := client.GetTableCByDate(day)
	if err == nil {
		t.Errorf("want: err != nil, got err == nil")
	}
}

func TestTablePrettyOutput(t *testing.T) {
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/A/",
			httpmock.NewStringResponder(200, mockTableA))
	} else {
		littleDelay()
	}

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
	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

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
	var day string = "2020-12-07"
	var tableNo string = "238/C/NBP/2020"

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/2020-12-07/",
			httpmock.NewStringResponder(200, mockTableC))
	} else {
		littleDelay()
	}

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

	if useMock {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", baseAddressTable+"/tables/C/last/5/",
			httpmock.NewStringResponder(200, mockTableCLast5))
	} else {
		littleDelay()
	}

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

func TestTableSetTableType(t *testing.T) {
	client := NewTable("B")
	if client.tableType != "B" {
		t.Errorf("want: B, got %s", client.tableType)
	}

	client.SetTableType("A")
	if client.tableType != "A" {
		t.Errorf("want: A, got %s", client.tableType)
	}
}
