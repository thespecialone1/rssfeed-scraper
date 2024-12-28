package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func respondWithErrors(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding with level 5xx error: ", msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg ,
	})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	log.Printf("Failed to Marshal JSON response: %v", payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}