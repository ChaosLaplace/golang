package controllers

import (
	"Heroku/utils/otp"
	"Heroku/utils/totp"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// HTTP Basic Authentication && TOTP
func AuthTotp(c *gin.Context) {
	password := GeneratePassCode("g77654321@gmail.comHENNGECHALLENGE003")
	fmt.Printf("HTTP Basic Authentication password \n")
	fmt.Printf("%s \n\n", password)

	authBasic := "g77654321@gmail.com:" + password
	totpBase64 := base64.StdEncoding.EncodeToString([]byte(authBasic))
	fmt.Printf("Header Authorization \n")
	fmt.Printf("Basic %s \n\n", totpBase64)
}

// Generates Passcode using a UTF-8 (not base32) secret and custom parameters
func GeneratePassCode(utf8string string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      0,
		Digits:    10,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}
