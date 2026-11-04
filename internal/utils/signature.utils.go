package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)



func VerifyHaloMailSignature(payload []byte, signatureHeader, secret string) bool{
	if secret == ""{
		return true
	}


	mac:= hmac.New(sha256.New,[]byte(secret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedSignature),[]byte(signatureHeader))
}