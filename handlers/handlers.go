package handlers

import (
	"log"
	"net/http"
	"os"

	"twister/app/middlew"
	"twister/app/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Controllers set PORT, Cors and start server listening
func Controllers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDataBase(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// handler configure the cors
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
