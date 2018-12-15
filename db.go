package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UID       int    `json:"uid"`
	Email     string `json:"email"`
	FirstName string `json:"fname"`
	blocked   int    `json:"blocked"`
}

type Answear struct {
	ID      int    `json:"id"`
	UID     int    `json:"uid"`
	Answear string `json:"answear"`
}

type PhoneValied struct {
	ID       int    `json:"id"`
	UID      int    `json:"uid"`
	Code     int    `json:"code"`
	Number   string `json:"number"`
	IsValied int    `json:"isValied"`
}

type ResultAPI struct {
	UID    int    `json:"uid"`
	Status string `json:"status"`
	Code   string `json:"code"`
}

type ResultInternalAPI struct {
	UID   int    `json:"uid"`
	Fname string `json:"fname"`
}

var conString = os.Getenv("DBCON")

func checkIfEmailExist(email string) bool {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("")
	results, err := db.Query("SELECT * FROM userDB WHERE email=?", email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag User
		err = results.Scan(&tag.UID, &tag.Email, &tag.FirstName, &tag.blocked)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("	[CheckPoint] User found with same email :: ", tag.UID)
		db.Close()
		return true
	}

	db.Close()
	return false
}

func getUserByEMmail(email string) User {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("")
	results, err := db.Query("SELECT * FROM userDB WHERE email=?", email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag User
		err = results.Scan(&tag.UID, &tag.Email, &tag.FirstName, &tag.blocked)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("	[CheckPoint] get user with id :: ", tag.UID)
		db.Close()
		return tag
	}

	db.Close()
	return User{UID: 0}
}
func createNewUser(email string, fname string) {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO userDB (email, fname) VALUES (?, ?)", email, fname)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insert.Close()
	db.Close()
	return
}

func createAuthCode(uid int, code string) {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO phoneValied (uid, code) VALUES (?, ?)", uid, code)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insert.Close()
	db.Close()
	return
}

func createAnswear(uid int, answear string) {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO giveaway (uid, answear) VALUES (?, ?)", uid, answear)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insert.Close()
	db.Close()
	return
}

func getAuthByCode(code string) PhoneValied {
	fmt.Println("Debug :: " + code)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("")
	results, err := db.Query("SELECT * FROM phoneValied WHERE code=?", code)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag PhoneValied
		err = results.Scan(&tag.ID, &tag.UID, &tag.Code, &tag.Number, &tag.IsValied)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("	[CheckPoint] get phone valied with code :: ", tag.Code)
		db.Close()
		return tag
	}

	db.Close()
	return PhoneValied{UID: 0}
}

func getUserById(uid int) User {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("")
	results, err := db.Query("SELECT * FROM userDB WHERE id=?", uid)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag User
		err = results.Scan(&tag.UID, &tag.Email, &tag.FirstName, &tag.blocked)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("	[CheckPoint] get user with id :: ", tag.UID)
		db.Close()
		return tag
	}

	db.Close()
	return User{UID: 0}
}

func updateVerifyr(uid int) {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("UPDATE phoneValied SET isValied = '1' WHERE id = ?;", uid)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insert.Close()
	db.Close()
	return
}
