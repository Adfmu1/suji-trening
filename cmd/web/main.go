package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	serv := &http.Server{
		Addr:              port,
		Handler:           mux,
		ReadHeaderTimeout: time.Duration(time.Duration.Seconds(5)),
	}

	err := serv.ListenAndServe()
	if err != nil {
		fmt.Println(http.StatusInternalServerError)
	}
}
