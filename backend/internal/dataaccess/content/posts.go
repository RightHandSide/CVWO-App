package contentdata

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/RightHandSide/CVWO-App/internal/api"
	"github.com/RightHandSide/CVWO-App/internal/database"
	"github.com/RightHandSide/CVWO-App/internal/models"
)

// List Post in Database
func ListPostByID(db *database.Database, threadid int) (*models.Thread, []models.Post, error) {
	// Get Thread
	row := db.SQL.QueryRow(`SELECT id, user_id, title, desc FROM threads WHERE id = ?`, threadid)

	// Insert into thread Variable
	var thread models.Thread
	if err := row.Scan(&thread.ID, &thread.User_ID, &thread.Title, &thread.Description); err != nil {
		// If No Row with ThreadID, Return nil, nil
		if err == sql.ErrNoRows {
			return nil, nil, err
		}
		return nil, nil, err
	}

	// List "posts" Table
	rows, err := db.SQL.Query(`SELECT id, thread_id, user_id, title, body FROM posts WHERE thread_id = ?`, threadid)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// Append All Posts Type to "result"
	var result []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Thread_ID, &post.User_ID, &post.Title, &post.Body); err != nil {
			return nil, nil, err
		}
		result = append(result, post)
	}

	// Return SQL Error
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}
	return &thread, result, nil
}

func MakePost(db *database.Database, user *models.User, r *http.Request, threadid int) error {
	// Decode Request Body as api.CreateRequest
	var req api.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	_, err := db.SQL.Exec("INSERT INTO posts (thread_id, user_id, title, body) VALUES (?, ?, ?, ?)",
		threadid, user.ID, req.Title, req.Body)
	if err != nil {
		return err
	}
	return nil
}
