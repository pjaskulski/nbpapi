package nbpapi

import (
	"testing"
	"time"
)

func TestLittleDelayShouldBeBetween400And650(t *testing.T) {
	for i := 1; i <= 10; i++ {
		start := time.Now()
		littleDelay()
		duration := time.Since(start)
		diff := duration.Milliseconds()
		if diff < 400 || diff > 650 {
			t.Errorf("Delay outside the expected range of 400-650")
		}
	}
}

func TestRandomInteger(t *testing.T) {
	var test int

	test = randomInteger(400, 650)
	if test < 400 || test > 650 {
		t.Errorf("randomInteger returned value outside expected range")
	}

	test = randomInteger(100, 150)
	if test < 100 || test > 150 {
		t.Errorf("randomInteger returned value outside expected range")
	}

	test = randomInteger(500, 900)
	if test < 500 || test > 900 {
		t.Errorf("randomInteger returned value outside expected range")
	}
}
