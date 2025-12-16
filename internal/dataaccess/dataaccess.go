package dataaccess

import (
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

func List_User(db *database.Database) ([]models.User, error) {
	rows, err := db.SQL.Query("SELECT id, name FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.User

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		result = append(result, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// Functions to Access db
