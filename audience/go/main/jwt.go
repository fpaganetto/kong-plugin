package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// Removes the Bearer prefix from the Authorization header
func GetToken(authHeader string) string {
	return authHeader[7:]
}

func GetAudience(jwt string) []string {
	base64Payload := strings.Split(jwt, ".")[1]
	decodedPayload, err := base64.StdEncoding.DecodeString(base64Payload)

	if err != nil {
		// handle error
	}

	pload := payload{}
	json.Unmarshal([]byte(decodedPayload), &pload)
	return pload.Aud
}

type payload struct {
	Aud []string `json:"aud"`
}

func Contains(elements []string, target string) bool {
	for _, elem := range elements {
		if elem == target {
			return true
		}
	}
	return false
}
