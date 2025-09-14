package service

import (
	"math/rand"

	"github.com/8bury/go-url-shortener/internal/base62"
	"github.com/8bury/go-url-shortener/internal/repo"
	"github.com/google/uuid"
)

type URLService struct {
	urlRepo *repo.URLRepository
}

func NewURLService(urlRepo *repo.URLRepository) *URLService {
	return &URLService{urlRepo: urlRepo}
}

func (s *URLService) CreateURL(longURL string) (string, error) {
	exists, err := s.urlRepo.DoesURLExist(longURL)
	if err != nil {
		return "", err
	}
	if exists {
		shortURL, err := s.urlRepo.GetShortURL(longURL)
		if err != nil {
			return "", err
		}
		return shortURL, nil
	}

	u := uuid.New().String()

	numericID := 1
	for i := 0; i < len(u); i++ {
		ch := u[i]
		if ch >= '0' && ch <= '9' {
			numericID += int(ch - '0')
		} else if ch >= 'A' && ch <= 'Z' {
			numericID += int(ch - 'A' + 11)
		} else if ch >= 'a' && ch <= 'z' {
			numericID += int(ch - 'a' + 73)
		}
	}

	salt := rand.Intn(100) * 23 * 7
	numericID *= salt

	shortURL := base62.ConvertFromInt(int(numericID))

	return shortURL, s.urlRepo.CreateURL(shortURL, longURL)
}

func (s *URLService) GetURL(shortURL string) (string, error) {
	return s.urlRepo.GetLongURL(shortURL)
}
