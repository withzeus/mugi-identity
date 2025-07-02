package helpers

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var GorillaDecoder = schema.NewDecoder()
var Validate = validator.New()

func WriteJSON(w io.Writer, i any) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func ReadJSON(r io.Reader, i any) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func ReadSchemaJSON(src map[string][]string, dst any) error {
	return GorillaDecoder.Decode(dst, src)
}

func ToJSON(dst any) ([]byte, error) {
	return json.Marshal(dst)
}

func FromJSON(j string, dst any) error {
	return json.Unmarshal([]byte(j), dst)
}
