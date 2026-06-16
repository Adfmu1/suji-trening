package main

import (
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /app", rootHandler)
	mux.HandleFunc("GET /app/healthz", healthHandler)

	return mux
}
