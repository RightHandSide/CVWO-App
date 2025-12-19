package dataaccess

import (
	"encoding/json"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// List Users in Database
func ListUser(db *database.Database) ([]models.User, error) {
	rows, err := db.SQL.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		result = append(result, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// Register User in Database
func RegisterUser(db *database.Database, r *http.Request) (string, error) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return "", err
	}

	var name string = req.Name
	_, err := db.SQL.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		return "", err
	}
	return name, nil
}

// List Threads in Database
func ListThread(db *database.Database) ([]models.Thread, error) {
	// Description cannot be NULL
	if _, err := db.SQL.Exec(`UPDATE threads SET desc = '' WHERE desc IS NULL`); err != nil {
		return nil, err
	}

	rows, err := db.SQL.Query("SELECT id, user_id, title, desc FROM threads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.User_ID, &thread.Title, &thread.Description); err != nil {
			return nil, err
		}
		result = append(result, thread)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
