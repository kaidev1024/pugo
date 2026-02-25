package puvalid

import "strings"

func ValidateEmail(email string) string {
	return strings.ToLower(email)
}
