package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	recaptcha "github.com/dpapathanasiou/go-recaptcha"
)

/*Create new API Token*/
func HandleQuestion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	fmt.Println("[->] createGiveaway")

	r.ParseForm()
	challenge := r.PostFormValue("g-recaptcha-response")
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	email := r.PostFormValue("email")
	fname := r.PostFormValue("fname")
	answear := r.PostFormValue("answear")
	if challenge == "" {
		fmt.Println("	-> no google!")
		w.WriteHeader(400)
		return
	}

	result, err := recaptcha.Confirm(ip, challenge)
	if err != nil {
		log.Println("recaptcha server error", err)
	}
	if result != true {
		w.WriteHeader(401)
		return
	}

	// Request is valied with google
	if email == "" {
		w.WriteHeader(400)
		fmt.Println("	-> no mail!")
		return
	}
	if answear == "" {
		fmt.Println("	-> no answear!")
		w.WriteHeader(400)
		return
	}

	if fname == "" {
		fmt.Println("	-> no fname!")
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

	codeRealCode := "";
	sum := 0
	for i := 0; i < 9000; i++ {
		sum += i

		phoneCode := generateAuthCode()
		if (getAuthByCode(phoneCode).UID == 0){
				fmt.Println("		CODE :: " + phoneCode)
				codeRealCode = phoneCode;
				break
		}
		fmt.Println("		CODE EXIST :: " + phoneCode)
	}

	fmt.Println("	Code generate after a try of: " , sum)
	
	
	fmt.Println("	-> phone code :: ", codeRealCode)
	createAuthCode(user.UID, codeRealCode)
	fmt.Println("	-> answear :: ", answear)
	createAnswear(user.UID, answear)

	sendSubmittMessage(fname, email, codeRealCode)

	gR := ResultAPI{UID: user.UID, Status: "ok", Code: codeRealCode}
	gRJ, _ := json.Marshal(gR)

	w.Write(gRJ)
	
}
