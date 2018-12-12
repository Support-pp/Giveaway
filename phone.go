package main

import (
	"math/rand"
)

var letterRunes = []rune("0123456789")

func generateAuthCode() string {
	b := make([]rune, 4)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}
