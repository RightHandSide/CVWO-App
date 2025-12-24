package database

import (
	"database/sql"
	"sync"

	_ "modernc.org/sqlite"
)

type Database struct {
	SQL *sql.DB
}

var (
	db   *Database
	once sync.Once
)

func GetDB() (*Database, error) {
	var err error

	once.Do(func() {
		var sqlDB *sql.DB
		sqlDB, err := sql.Open("sqlite", "app.db")
		if err != nil {
			return
		}

		if err := sqlDB.Ping(); err != nil {
			return
		}

		if _, err := sqlDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
			return
		}

		db = &Database{
			SQL: sqlDB,
		}
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Database
// users: id, name
// threads: id, user_id, title, desc
// posts: id, thread_id, user_id, title, body
// comments: id, post_id, user_id, parent_comment_id, body
