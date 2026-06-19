package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Adfmu1/suji-trening/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	Database *database.Queries
	Logger   *slog.Logger
}

var app application

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(http.StatusInternalServerError, err)
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	env := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("an error occured while opening database")
		fmt.Println(err)
		return
	}
	defer dbConn.Close()

	app.Database = database.New(dbConn)
	app.Logger = newLogger(env)

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
