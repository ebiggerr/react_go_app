package main

import (
	"fmt"
	"net/http"
)

func main() {
	handleWebRequests()
}

func handleWebRequests() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		return
	}
}

func homePage(rw http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(rw, "Hey! Welcome.")
	if err != nil {
		return
	}
}
