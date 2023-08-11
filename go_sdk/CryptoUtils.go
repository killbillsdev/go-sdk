package go_sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func CipherHmacPayload(payload string, hmacKey string) string {
	h := hmac.New(sha256.New, []byte(hmacKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}