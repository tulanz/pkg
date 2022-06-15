package crypto

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// GenerateHashAndSalt generate for password encrypt
func GenerateHashAndSalt(pwd string) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		zap.L().Error("GenerateHashAndSalt error", zap.Error(err))
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePassword ComparePassword
func ComparePassword(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		zap.L().Error("ComparePassword error", zap.Error(err))
		return false
	}

	return true
}
