package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	recaptcha "github.com/dpapathanasiou/go-recaptcha"
	"github.com/gorilla/mux"
)

const gameId = "0"

func main() {
	fmt.Println("[Start] API on Port :: ", os.Getenv("PORT"))

	//Google recaptcha
	recaptcha.Init(os.Getenv("recaptchaPrivateKey"))

	//Webserver
	router := mux.NewRouter()
	router.HandleFunc("/api/new", HandleQuestion).Methods("POST")
	router.HandleFunc("/api/internal", VerifyRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
