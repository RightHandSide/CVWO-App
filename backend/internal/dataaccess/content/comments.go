package contentdata

import (
	"database/sql"

	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// List Post in Database
func ListCommentByID(db *database.Database, postid int) (*models.Post, []models.Comment, error) {
	// Get Post
	row := db.SQL.QueryRow(`SELECT id, thread_id, user_id, title, body FROM posts WHERE id = ?`, postid)

	// Insert into post Variable
	var post models.Post
	if err := row.Scan(&post.ID, &post.Thread_ID, &post.User_ID, &post.Title, &post.Body); err != nil {
		// If No Row with PostID, Return nil, nil
		if err == sql.ErrNoRows {
			return nil, nil, err
		}
		return nil, nil, err
	}

	// List "comments" Table
	rows, err := db.SQL.Query(`SELECT id, post_id, user_id, body FROM comments WHERE post_id = ?`, postid)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// Append All Comments Type to "result"
	var result []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.Post_ID, &comment.User_ID, &comment.Body); err != nil {
			return nil, nil, err
		}
		result = append(result, comment)
	}

	// Return SQL Error
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}
	return &post, result, nil
}
