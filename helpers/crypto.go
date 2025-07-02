package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func createRandomBytes(b []byte) error {
	_, err := rand.Read(b)
	return err
}

func RandomHexEncoded(len int) (string, error) {
	b := make([]byte, len)
	if err := createRandomBytes(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func RandomB64StdEncoded(len int) (string, error) {
	b := make([]byte, len)
	if err := createRandomBytes(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
