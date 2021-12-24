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

	// Endpoints of Users
	router.HandleFunc("/register", middlew.CheckDataBase(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDataBase(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckDataBase(middlew.ValidateJWT(routers.ShowProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckDataBase(middlew.ValidateJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/uploadAvatar", middlew.CheckDataBase(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadBanner", middlew.CheckDataBase(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDataBase(routers.UploadBanner)).Methods("GET")
	router.HandleFunc("/getBanner", middlew.CheckDataBase(routers.UploadBanner)).Methods("GET")

	// Endpoints of Tweets
	router.HandleFunc("/addTweet", middlew.CheckDataBase(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/userTweets", middlew.CheckDataBase(middlew.ValidateJWT(routers.ReadUserTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDataBase(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// handler configure the cors
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
