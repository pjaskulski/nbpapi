// auxiliary program functions

package nbpapi

import (
	"math/rand"
	"time"
)

// littleDelay - delay function, so as not to bother the NBP server too much...
func littleDelay() {
	interval := randomInteger(400, 650)
	time.Sleep(time.Millisecond * time.Duration(interval))
}

// randomInteger func
func randomInteger(minValue int, maxValue int) int {
	var result int

	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(maxValue-minValue+1) + minValue

	return result
}
