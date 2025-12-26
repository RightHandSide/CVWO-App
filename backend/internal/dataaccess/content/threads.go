package contentdata

import (
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// List Threads in Database
func ListThread(db *database.Database) ([]models.Thread, error) {
	// List "threads" Table
	rows, err := db.SQL.Query("SELECT id, user_id, title, desc FROM threads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Append All Thread Type to "result"
	var result []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.User_ID, &thread.Title, &thread.Description); err != nil {
			return nil, err
		}
		result = append(result, thread)
	}

	// Return SQL Error
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
