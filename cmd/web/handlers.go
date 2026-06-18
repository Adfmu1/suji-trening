package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Adfmu1/suji-trening/internal/authentication"
	"github.com/Adfmu1/suji-trening/internal/database"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from a root handler!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (app application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Firstname string `json:"firstname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		fmt.Println(err)
		return
	}

	hPass, err := authentication.HashPassword(params.Password)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := app.Database.CreateUser(r.Context(), database.CreateUserParams{
		Firstname:      params.Firstname,
		Email:          params.Email,
		Hashedpassword: hPass,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	respondWithJson(w, http.StatusOK, resp)
}
