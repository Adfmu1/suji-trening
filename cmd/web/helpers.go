package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type AdvancedResponseWriter struct {
	http.ResponseWriter
	statusCode int
	bytes      int
}

func (arw *AdvancedResponseWriter) WriteHeader(code int) {
	arw.statusCode = code
	arw.ResponseWriter.WriteHeader(code)
}

func (arw *AdvancedResponseWriter) Write(b []byte) (int, error) {
	n, err := arw.ResponseWriter.Write(b)
	arw.bytes += n
	return n, err
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	errDat := struct {
		ErrMsg string `json:"error"`
	}{
		ErrMsg: msg,
	}
	respondWithJson(w, code, errDat)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	var errAf error
	defer func() {
		if errAf != nil {
			log.Printf("An error has occured during writing response %v", data)
		}
	}()

	if err != nil {
		log.Printf("error while marshalling the reponse: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, errAf = w.Write([]byte(`{"error":"something went wrong"}`))
		return
	}

	w.WriteHeader(code)
	_, errAf = w.Write(data)
}

func checkIfEmailInDB(r *http.Request, email string) bool {
	_, err := app.Database.GetUserByEmail(r.Context(), email)
	return err != sql.ErrNoRows
}
