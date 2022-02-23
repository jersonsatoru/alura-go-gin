package main

import (
	"net/http"
)

func main() {
	srv := &http.Server{
		Handler: Routes(),
		Addr:    ":8080",
	}
	srv.ListenAndServe()
}
