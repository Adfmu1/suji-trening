package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from a root handler!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
