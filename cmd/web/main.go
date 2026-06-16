package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(http.StatusInternalServerError, err)
	}
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	serv := &http.Server{
		Addr:              port,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}

	mux.HandleFunc("/", rootHandler)

	fmt.Println("Listening on a ", port)

	err = serv.ListenAndServe()
	if err != nil {
		fmt.Println(http.StatusInternalServerError, err)
	}
}
