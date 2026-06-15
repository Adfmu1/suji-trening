package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	serv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	serv.ListenAndServe()
}
