package main

import (
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /app", rootHandler)
	mux.HandleFunc("GET /api/healthz", healthHandler)
	mux.HandleFunc("POST /api/register", app.createUserHandler)
	mux.HandleFunc("POST /api/deleteuser", app.deleteUserHandler)

	return app.loggingMiddleware(mux)
}
