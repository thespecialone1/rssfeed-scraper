package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"githhub.com/thespecialone1/rssfeed-scraper/internal/database"
	"github.com/google/uuid"
)
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name" `
	}
	decoder:= json.NewDecoder(r.Body)
	params:= parameters{}
	err:= decoder.Decode(&params)
	if err!=nil{
		respondWithErrors(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err:= apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err!=nil {
		respondWithErrors(w, 400, fmt.Sprintf("Couldn't create user: %s", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
} 