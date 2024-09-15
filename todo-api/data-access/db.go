package data_access

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"sync"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

func GetInstance() *sql.DB {
	once.Do(func() {
		db, err := sql.Open("sqlite3", "sqlite-storage.db")
		if err != nil {
			panic(err)
		}

		dbInstance = db
	})

	return dbInstance
}

func InitDB() {
	db := GetInstance()

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed BOOLEAN DEFAULT 0,
			completed_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		panic(err)
	}
}
