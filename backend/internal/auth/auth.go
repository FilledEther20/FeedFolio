package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	v := strings.Split(val, " ")
	if len(v) < 2 {
		return "", errors.New("no authentication info found")
	}
	if v[0] != "ApiKey" {
		return "", errors.New("malformed authentication header found")
	}
	return v[1], nil
}
