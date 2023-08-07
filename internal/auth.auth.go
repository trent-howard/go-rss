package auth

import (
	"fmt"
	"net/http"
	"strings"
)

// GetAPIKey extacts the API key from HTTP request headers
// E.g. Authorization: ApiKey {key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", fmt.Errorf("no authentication found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", fmt.Errorf("malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", fmt.Errorf("malformed Authorization header")
	}
	return vals[1], nil
}
