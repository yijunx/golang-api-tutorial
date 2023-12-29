package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/yijunx/golang-api-tutorial/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	// r points to a struct where we going to set up the api
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`

\/\/\/\/\/\/\/\/\ GO API /\/\/\/\/\/\/\/\/\
\/\/\/\/\/\/\/\/\ YEAHHH /\/\/\/\/\/\/\/\/\`)
	err := http.ListenAndServe("0.0.0.0:8000", r)
	if err != nil {
		log.Error(err)
	}
}

// curl --location 'http://localhost:3721/account/coins?username=tom' --header 'Authorization: tomtom'
