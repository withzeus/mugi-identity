package lib

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
)

type Encoder struct{}

func (e *Encoder) GenerateEncodedString(c Crypto, len int) string {
	return c.EncodeString(len)
}

type Crypto interface {
	EncodeString(len int) string
}

type Rand struct{}

func (r *Rand) RandomBytes(b []byte) {
	rand.Read(b)
}

type Hex struct{}

func (h *Hex) EncodeString(len int) string {
	r := Rand{}
	b := make([]byte, len)
	r.RandomBytes(b)
	return hex.EncodeToString(b)
}

type Base64 struct{}

func (b64 *Base64) EncodeString(len int) string {
	r := Rand{}
	b := make([]byte, len)
	r.RandomBytes(b)
	return base64.StdEncoding.EncodeToString(b)
}

func (e *Encoder) ToJSON(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("encoding.ToJSON - %s", err.Error())
	}
	return b
}

func (j *Encoder) FromJSON(b []byte, v any) {
	err := json.Unmarshal(b, v)
	if err != nil {
		log.Printf("encoding.FromJSON - %s", err.Error())
	}
}

func (j *Encoder) FromIoReader(ir io.Reader, v any) {
	b, err := io.ReadAll(ir)

	if err != nil {
		log.Printf("encoding.FromIoReader - %s", err.Error())
	}

	j.FromJSON(b, v)
}
