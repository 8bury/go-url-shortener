package database

//sqlite database
import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConfigureDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "urls.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err := createTable(db); err != nil {
		log.Fatal("Error creating table:", err)
	}
	return db

}

func createTable(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_url TEXT UNIQUE NOT NULL,
		long_url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table:", err)
		return err
	}
	return nil
}
