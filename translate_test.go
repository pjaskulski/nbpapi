package nbpapi

import "testing"

func TestPolishTranslationShouldBeOk(t *testing.T) {
	var want, got string

	setLang("pl")

	want = "Typ tabeli:"
	got = l.Get("Table type:")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

	want = "Numer tabeli:"
	got = l.Get("Table number:")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

	want = "Nazwa waluty:"
	got = l.Get("Currency name:")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

	want = "TABELA,DATA,ŚREDNI (PLN)"
	got = l.Get("TABLE,DATE,AVERAGE (PLN)")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

}

func TestPolishTranslationMissingKey(t *testing.T) {
	var want, got string

	setLang("pl")

	want = "Gold value"
	got = l.Get("Gold value")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

	want = "currency exchange rates"
	got = l.Get("currency exchange rates")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

func TestEnglishTranslationShouldBeOk(t *testing.T) {
	var want, got string

	setLang("en")

	want = "Table type:"
	got = l.Get("Table type:")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

	want = "Table number:"
	got = l.Get("Table number:")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

	want = "Currency name:"
	got = l.Get("Currency name:")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

	want = "TABLE,DATE,AVERAGE (PLN)"
	got = l.Get("TABLE,DATE,AVERAGE (PLN)")
	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}

}
