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

func (d *URLRepository) DoesURLExist(longURL string) (bool, error) {
	row := d.db.QueryRow("SELECT COUNT(*) FROM urls WHERE long_url = ?", longURL)
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("Error checking if URL exists in database:", err)
		return false, err
	}
	return count > 0, nil
}

func (d *URLRepository) GetShortURL(longURL string) (string, error) {
	row := d.db.QueryRow("SELECT short_url FROM urls WHERE long_url = ?", longURL)
	var shortURL string
	err := row.Scan(&shortURL)
	if err != nil {
		log.Println("Error getting short URL from database:", err)
		return "", err
	}
	return shortURL, nil
}

func (d *URLRepository) GetLongURL(shortURL string) (string, error) {
	row := d.db.QueryRow("SELECT long_url FROM urls WHERE short_url = ?", shortURL)
	var longURL string
	err := row.Scan(&longURL)
	if err != nil {
		log.Println("Error getting URL from database:", err)
		return "", err
	}
	return longURL, nil
}
