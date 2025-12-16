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

		db = &Database{
			SQL: sqlDB,
		}
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Connect to app.db
