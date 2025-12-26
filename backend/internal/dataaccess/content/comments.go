package contentdata

import (
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// List Post in Database
func ListCommentByID(db *database.Database, postid int) ([]models.Comment, error) {
	// List "comments" Table
	rows, err := db.SQL.Query(`SELECT id, post_id, user_id, body FROM comments WHERE post_id = ?`, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Append All Comments Type to "result"
	var result []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.Post_ID, &comment.User_ID, &comment.Body); err != nil {
			return nil, err
		}
		result = append(result, comment)
	}

	// Return SQL Error
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
