package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*Create new API Token*/
func HandleQuestion(w http.ResponseWriter, r *http.Request) {

	fmt.Println("[->] createGiveaway")

	r.ParseForm()
	challenge := r.PostFormValue("g-recaptcha-response")
	//ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	email := r.PostFormValue("email")
	fname := r.PostFormValue("fname")
	answear := r.PostFormValue("answear")
	if challenge == "" {
		w.WriteHeader(400)
		return
	}
	/*
		result, err := recaptcha.Confirm(ip, challenge)
		if err != nil {
			log.Println("recaptcha server error", err)
		}
		if result != true {
			w.WriteHeader(401)
			return
		}
	*/
	// Request is valied with google
	if email == "" {
		w.WriteHeader(400)
		return
	}
	if answear == "" {
		w.WriteHeader(400)
		return
	}

	if fname == "" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("	-> " + email)

	cmail := checkIfEmailExist(email)
	fmt.Println("	-> Email exist status :: ", cmail)
	if cmail {
		w.WriteHeader(409)
		return
	}

	createNewUser(email, fname)
	user := getUserByEMmail(email)
	fmt.Println("	-> user created with uid :: ", user.UID)
	phoneCode := generateAuthCode()
	fmt.Println("	-> phone code :: ", phoneCode)
	createAuthCode(user.UID, phoneCode)
	fmt.Println("	-> answear :: ", answear)
	createAnswear(user.UID, answear)

	sendSubmittMessage(fname, email, phoneCode)

	gR := ResultAPI{UID: user.UID, Status: "ok", Code: phoneCode}
	gRJ, _ := json.Marshal(gR)

	w.Write(gRJ)
}
