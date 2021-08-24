package utils

import (
	"math/rand"
	"strconv"
	"time"
)

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-!?"

func RandomKey() string {
	return string(randomStringBytes(30))
}

func RandomFiveDigitNumber() string {
	rand.Seed((time.Now().UnixNano()))
	min := 10000
	max := 99999

	randomInt := rand.Intn(max-min) + min
	return strconv.Itoa(randomInt)
}

func GenerateVerificationKey() string {
	return string(randomStringBytes(6))
}

func randomStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = CHARSET[rand.Intn(len(CHARSET))]
	}
	return string(b)
}
