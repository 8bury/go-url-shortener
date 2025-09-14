package main

import (
	"net/http"
	"os"

	"github.com/8bury/go-url-shortener/internal"
	"github.com/8bury/go-url-shortener/internal/middleware"
)

func main() {
	api := http.NewServeMux()

	internal.ConfigureDependencies(api)

	corsHandler := middleware.CorsMiddleware(api)

	http.ListenAndServe(":"+os.Getenv("PORT"), corsHandler)
}
