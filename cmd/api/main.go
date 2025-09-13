package main

import (
	"net/http"

	"github.com/8bury/go-url-shortener/internal"
)

func main() {
	api := http.NewServeMux()

	internal.ConfigureDependencies(api)

	http.ListenAndServe(":8080", api)
}
