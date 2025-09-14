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

	port := os.Getenv("PORT_URL_SHORTENER")
	if port == "" {
		port = "8081"
	}

	http.ListenAndServe(":"+port, corsHandler)
}
