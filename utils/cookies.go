package utils

import (
	"net/http"
	"strings"
)

func GetUserFromCookie(r *http.Request) (string, string, error) {

	var userId, userName string
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return "", "", err
	}

	if cookie != nil && cookie.Value != "" {
		key := []byte("6091835705053067")
		decrypted, err := Decrypt(key, cookie.Value)
		if err != nil {
			return "", "", err
		}

		userId = strings.Split(decrypted, ":")[0]
		userName = strings.Split(decrypted, ":")[1]
	}

	return userId, userName, nil
}
