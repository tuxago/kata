package main

import (
	"math/rand"
	"time"

	"github.com/tuxago/kata/roll"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	println(roll.D6())

}
