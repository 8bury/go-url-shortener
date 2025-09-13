package repo

import (
	"database/sql"
	"log"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (d *URLRepository) CreateURL(shortURL string, longURL string) error {
	_, err := d.db.Exec("INSERT INTO urls (short_url, long_url) VALUES (?, ?)", shortURL, longURL)
	if err != nil {
		log.Println("Error creating URL in database:", err)
		return err
	}
	return nil
}

func (d *URLRepository) GetURL(shortURL string) (string, error) {
	row := d.db.QueryRow("SELECT long_url FROM urls WHERE short_url = ?", shortURL)
	var longURL string
	err := row.Scan(&longURL)
	if err != nil {
		log.Println("Error getting URL from database:", err)
		return "", err
	}
	return longURL, nil
}
