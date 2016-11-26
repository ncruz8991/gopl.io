// Page 23
package main

import (
	"fmt"
	"github.com/ncruz8991/gopl.io/ch1/lissajous/lissajous"
	"log"
	"net/http"
	"strconv"
)

const (
	cycles = "cycles"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if numCycles := r.Form.Get(cycles); numCycles != "" {
		fmt.Printf("Cycles found: %s\n", numCycles)
		cyclesInt, _ := strconv.Atoi(numCycles)
		lissajous.LissajousCycles(w, cyclesInt)
	} else {
		fmt.Println("No cycles found.")
		lissajous.Lissajous(w)
	}
}