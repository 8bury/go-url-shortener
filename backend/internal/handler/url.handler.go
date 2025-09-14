package handler

import (
	"encoding/json"
	"net/http"

	"github.com/8bury/go-url-shortener/internal/service"
)

type URLHandler struct {
	UrlService *service.URLService
}

func NewURLHandler(api *http.ServeMux, urlService *service.URLService) *URLHandler {
	handler := &URLHandler{UrlService: urlService}

	api.HandleFunc("POST /", handler.CreateURL)
	api.HandleFunc("GET /{shortURL}", handler.RedirectURL)

	return handler
}

func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	longURL := r.FormValue("longURL")

	if longURL == "" {
		http.Error(w, "missing longURL", http.StatusBadRequest)
		return
	}

	shortURL, err := h.UrlService.CreateURL(longURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"shortURL": shortURL,
		"longURL":  longURL,
	})
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.PathValue("shortURL")
	longURL, err := h.UrlService.GetURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusPermanentRedirect)
}
