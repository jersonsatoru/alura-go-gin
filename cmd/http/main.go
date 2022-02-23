package main

import (
	"net/http"

	"github.com/jersonsatoru/alura-go-gin/internal/db"
)

func main() {
	db.Connect()
	srv := &http.Server{
		Handler: Routes(),
		Addr:    ":8080",
	}
	srv.ListenAndServe()
}
