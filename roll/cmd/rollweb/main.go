package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()
	r.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		tmpl := `<html><body>
			<form method="POST">
				<input type="submit" value="ROLL">
			</form>	
		</body></html>`
		fmt.Fprintf(w, tmpl)
	}).Methods("GET")
	r.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		rsp, err := http.Get("http://localhost:8080/")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rsp.Body.Close()

		b, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", b)
	}).Methods("POST")
	http.ListenAndServe(":9999", r)
}
