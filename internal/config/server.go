package config

import (
	"net/http"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + PORT(),
		Handler: handler,
	}
}
