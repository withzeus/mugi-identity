package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
)

func SendHttpError(rw http.ResponseWriter, code int, message string) {
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(map[string]any{
		"error": message,
	})
}

func SendBadRequestError(rw http.ResponseWriter) {
	SendHttpError(rw, http.StatusBadRequest, "Bad Request")
}

func ClassifyUserHandle(i string) (string, error) {
	phoneRegex := regexp.MustCompile(`^(?:\+?959|09)\d{7,9}$`)
	if phoneRegex.MatchString(i) {
		return "phone", nil
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(i) {
		if strings.HasSuffix(strings.ToLower(i), "@gmail.com") {
			return "gmail", nil
		}
		return "email", nil
	}
	return "", errors.New("invalid input")
}
