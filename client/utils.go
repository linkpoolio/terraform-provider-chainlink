package client

import (
	"math/rand"
	"time"
)

func Delimiter() string {
	return "<clnode>"
}

func RandomString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
