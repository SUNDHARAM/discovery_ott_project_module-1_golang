package middlewares

import (
	"regexp"
)

func IsPassword(password string) bool {

	secure := true
	tests := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, password)
		if !t {
			secure = false
			break
		}
	}
	return secure
}
