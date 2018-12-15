package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func VerifyRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("[->] phoneVerify")

	r.ParseForm()
	spp := r.PostFormValue("secert-pp")
	code := r.PostFormValue("code")

	if spp == "test123" {
		aut := getAuthByCode(code)
		codeInt, _ := strconv.Atoi(code)
		if aut.Code != codeInt {
			return
		}
		user := getUserById(aut.UID)

		if aut.UID != 0 {
			updateVerifyr(aut.UID)
		}
		gR := ResultInternalAPI{UID: aut.UID, Fname: user.FirstName}
		gRJ, _ := json.Marshal(gR)

		w.Write(gRJ)
		return
	}
	w.WriteHeader(402)

}
