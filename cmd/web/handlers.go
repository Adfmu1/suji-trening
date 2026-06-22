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
	defer r.Body.Close()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	if checkIfEmailInDB(r, params.Email) {
		respondWithError(w, http.StatusConflict, "User with this email already exists")
		return
	}

	hPass, err := authentication.HashPassword(params.Password)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Password hashing failed")
		return
	}

	resp, err := app.Database.CreateUser(r.Context(), database.CreateUserParams{
		Firstname:      params.Firstname,
		Email:          params.Email,
		Hashedpassword: hPass,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	respondWithJson(w, http.StatusOK, resp)
}

func (app application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&params)
	defer r.Body.Close()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}
	if !checkIfEmailInDB(r, params.Email) {
		respondWithError(w, http.StatusConflict, "User with this email doesn't exists")
		return
	}

	dbUsr, err := app.Database.GetUserByEmail(r.Context(), params.Email)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	if !authentication.CheckHashPassword(params.Password, dbUsr.Hashedpassword) {
		respondWithError(w, http.StatusUnauthorized, "Wrong password")
		return
	}

	err = app.Database.DeleteUser(r.Context(), database.DeleteUserParams{
		Email:          params.Email,
		Hashedpassword: dbUsr.Hashedpassword,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	respondWithJson(w, http.StatusOK, struct{}{})
}
