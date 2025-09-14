package internal

import (
	"net/http"

	"github.com/8bury/go-url-shortener/internal/database"
	"github.com/8bury/go-url-shortener/internal/handler"
	"github.com/8bury/go-url-shortener/internal/repo"
	"github.com/8bury/go-url-shortener/internal/service"
)

func ConfigureDependencies(api *http.ServeMux) {
	db := database.ConfigureDatabase()
	urlRepo := repo.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	handler.NewURLHandler(api, urlService)
}
