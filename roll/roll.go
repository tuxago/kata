package roll

import "math/rand"

func D6() int {
	n := rand.Intn(6) + 1
	return n
}
