package stringutils

import (
	"math/rand"
	"time"
)

// RandomCode4 验证码4位
func RandomCode4() string {
	return RandomCaptcha(4)
}
func RandomCode6() string {
	return RandomCaptcha(6)
}

// RandomString copy from https://github.com/moby/moby/blob/master/pkg/stringutils/stringutils.go
func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomStringV2 RandomStringV2
func RandomStringV2(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomCaptcha RandomCaptcha
func RandomCaptcha(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("0123456789")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomCode RandomCode
func RandomCode(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
