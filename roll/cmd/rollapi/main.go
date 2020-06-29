package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/tuxago/kata/roll"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h := roll.D6()
		fmt.Fprintf(w, "%d", h)
	})
	http.ListenAndServe(":8080", nil)
}
