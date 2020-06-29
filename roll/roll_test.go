package roll_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tuxago/kata/roll"
)

func TestD6(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := roll.D6()
	if n <= 0 && n > 6 {
		t.Error("dice roll not ok")
	}
}
