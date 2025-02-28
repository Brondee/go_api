package main

import (
	"fmt"
	"net/http"

	"github.com/brondee/go_api/internal/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)


func main(){

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("starting go api")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}