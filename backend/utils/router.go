package utils

import (
	"math/rand"
)

const (
	CHARSET      = "abcdefghijklmnopqrstuvwxyz-_+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CHARSETLEN   = len(CHARSET)
	DEEPLINKSIZE = 15
)

func GenerateDeepLink() string {

	var result string

	for DEEPLINKSIZE != len(result) {
		index := rand.Intn(CHARSETLEN)
		result += string(CHARSET[index])
	}

	return result
}
