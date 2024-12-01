package main

import (
	"fmt"
	"github.com/LewisT543/simple-go-api/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting Go API service...")

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Error(err)
	}
}
