package contentdata

import (
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// List Post in Database
func ListPostByID(db *database.Database, threadid int) ([]models.Post, error) {
	// List "posts" Table
	rows, err := db.SQL.Query(`SELECT id, thread_id, user_id, title, body FROM posts WHERE thread_id = ?`, threadid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Append All Posts Type to "result"
	var result []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Thread_ID, &post.User_ID, &post.Title, &post.Body); err != nil {
			return nil, err
		}
		result = append(result, post)
	}

	// Return SQL Error
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
