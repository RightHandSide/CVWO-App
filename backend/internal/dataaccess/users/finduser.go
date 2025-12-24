package userdata

import (
	"database/sql"

	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// Check User Existence in Database
func GetUserByName(db *database.Database, name string) (*models.User, error) {
	// Get Row of Same Name
	row := db.SQL.QueryRow(`SELECT id, name FROM users WHERE name = ?`, name)

	// Insert ID & Name into user Variable
	var user models.User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		// If No Row with Name, Return nil, nil
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	// Return User Type
	return &user, nil
}
