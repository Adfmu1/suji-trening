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

	mux := routes()

	serv := &http.Server{
		Addr:              port,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}

	fmt.Println("Listening on a ", port)

	err = serv.ListenAndServe()
	if err != nil {
		fmt.Println(http.StatusInternalServerError, err)
	}
}
