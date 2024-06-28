package utils

import (
	"math/rand"
)

const (
	CHARSET      = "abcdefghijklmnopqrstuvwxyz-_+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CHARSETLEN   = len(CHARSET)
	DEEPLINKSIZE = 15
)

type Deeplink string

func GenerateDeepLink() Deeplink {

	var result Deeplink

	for DEEPLINKSIZE != len(result) {
		index := rand.Intn(CHARSETLEN)
		result += Deeplink(CHARSET[index])
	}

	return result
}
