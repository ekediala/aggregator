package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization")

	if header == "" {
		return "", errors.New("no authentication info found")
	}

	values := strings.Split(header, " ")

	if len(values) != 2 {
		return "", errors.New("incorrect header format")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("incorrect header format")
	}

	apiKey := values[1]

	if apiKey == "" {
		return "", errors.New("no api key found")
	}

	return apiKey, nil

}
