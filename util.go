package xomologisou

import (
	"crypto/rand"
	"encoding/base32"
	"log"
)

func RandomId(len int) string {
	code := make([]byte, len)
	_, err := rand.Read(code)
	if err != nil {
		log.Panic(err)
	}
	return base32.HexEncoding.EncodeToString(code)[:len]
}
