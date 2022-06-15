package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// MD5 GetMD5
func MD5(str string) string {
	md5 := md5.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}

// EncodePasswd EncodePasswd
func EncodePasswd(password, salt string) string {
	newPassword := pbkdf2.Key([]byte(password), []byte(salt), 10000, 50, sha256.New)
	return base64.StdEncoding.EncodeToString(newPassword)
}

// EncryptPassword 使用SHA256加密
func EncryptPassword(password, salt string) string {
	newPassword := pbkdf2.Key([]byte(password), []byte(salt), 10000, 50, sha256.New)
	return fmt.Sprintf("%x", newPassword)
}
