package userdata

import (
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// List Users in Database
func ListUser(db *database.Database) ([]models.User, error) {
	// List "users" Table
	rows, err := db.SQL.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Append All User Type to "result"
	var result []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		result = append(result, user)
	}

	// Return SQL Error
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
