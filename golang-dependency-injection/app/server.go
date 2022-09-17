package app

import (
	"net/http"

	"github.com/Masred/dasar-golang/golang-dependency-injection/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
