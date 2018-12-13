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
	router.HandleFunc("/api/new", myHandler).Methods("OPTIONS")
	router.HandleFunc("/api/internal", myHandler).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "giveaway.support-pp.de")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
