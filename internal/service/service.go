package service

import (
	"github.com/8bury/go-url-shortener/internal/repo"
)

type URLService struct {
	urlRepo *repo.URLRepository
}

func NewURLService(urlRepo *repo.URLRepository) *URLService {
	return &URLService{urlRepo: urlRepo}
}
