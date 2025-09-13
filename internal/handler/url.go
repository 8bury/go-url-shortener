package handler

import (
	"net/http"

	"github.com/8bury/go-url-shortener/internal/service"
)

type URLHandler struct {
	UrlService *service.URLService
}

func NewURLHandler(api *http.ServeMux, urlService *service.URLService) *URLHandler {
	handler := &URLHandler{UrlService: urlService}

	api.HandleFunc("/url", handler.HandleURL)

	return handler
}

func (h *URLHandler) HandleURL(w http.ResponseWriter, r *http.Request) {
	return
}
