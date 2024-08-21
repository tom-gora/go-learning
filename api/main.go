package api

import (
	"fmt"
	"io"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world from golang!")
}

func otherRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my other route!")
}

func repeaterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Fprintf(w, "Request body: %s", requestBody)
}

func Run() {
	fmt.Println("Hello - this is my first simple api running.")
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/other", otherRouteHandler)

	mux.HandleFunc("/repeater", repeaterHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
