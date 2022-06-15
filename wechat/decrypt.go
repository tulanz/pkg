package wechat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"

	"github.com/tulanz/pkg/util"
)

type WechatWatermark struct {
	AppId     string `json:"appid,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type WechatOAuth struct {
	UnionId    string `json:"unionid,omitempty"`
	OpenId     string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
}
type WechatUser struct {
	UnionId         string          `json:"unionId,omitempty"`
	OpenId          string          `json:"openId,omitempty"`
	NickName        string          `json:"nickName,omitempty"`
	Gender          uint            `json:"gender,omitempty"`
	City            string          `json:"city,omitempty"`
	Province        string          `json:"province,omitempty"`
	Country         string          `json:"country,omitempty"`
	Avatar          string          `json:"avatarUrl,omitempty"`
	Language        string          `json:"language,omitempty"`
	PhoneNumber     string          `json:"phoneNumber,omitempty"`
	PurePhoneNumber string          `json:"purePhoneNumber,omitempty"`
	CountryCode     string          `json:"countryCode,omitempty"`
	Watermark       WechatWatermark `json:"watermark,omitempty"`
}

// Decrypt Decrypt
func Decrypt(sessionKey string, ivText string, cryptoText string) (phone WechatUser, err error) {
	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return phone, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return phone, err
	}
	iv, err := base64.StdEncoding.DecodeString(ivText)
	if err != nil {
		return phone, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return phone, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext, err = util.PKCS7UnPadding(ciphertext, block.BlockSize())
	if err != nil {
		return phone, err
	}
	err = json.Unmarshal(ciphertext, &phone)
	if err != nil {
		return
	}
	return phone, nil
}
