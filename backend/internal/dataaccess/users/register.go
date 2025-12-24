package userdata

import (
	"encoding/json"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/api"
	"github.com/RightHandSide/CVWO-App/internal/database"
)

// Register User in Database
func RegisterUser(db *database.Database, r *http.Request) (string, error) {
	// Decode Request Body as api.LoginRequest
	var req api.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return "", err
	}

	// Insert Name into Database, Return Name / Error
	var name string = req.Name
	_, err := db.SQL.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		return "", err
	}
	return name, nil
}
