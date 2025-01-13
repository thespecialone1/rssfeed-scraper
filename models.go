package main

import (
	"time"

	"githhub.com/thespecialone1/rssfeed-scraper/internal/database"
	"github.com/google/uuid"
)
type User struct {
	ID        uuid.UUID `json:"id"`	// change uuid.UUID to string
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Name      string	`json:"name"`
}

func databaseUserToUser (dbUser database.User) User{
	return User{
		ID: dbUser.ID,		//convert UUID to string
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
	}
}