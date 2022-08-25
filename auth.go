package helper

import (
	"encoding/base64"
	"strings"
)

// Takes input strings username and password and returns base64(username:password)
func BasicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Takes input string b64 (base64 encoded string) and return decoded username and password
func BasicAuthDecode(b64 string) (string, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", "", err
	}
	credentials := strings.Split(string(decoded), ":")
	return credentials[0], strings.Split(credentials[1], "\n")[0], nil
}
