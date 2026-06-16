package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
	if err != nil {
		log.Printf("error while marshalling the reponse: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"something went wrong"}`))
		return
	}

	w.WriteHeader(code)
	w.Write(data)
}
