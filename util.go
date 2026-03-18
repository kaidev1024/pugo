package pugo

import (
	"fmt"
	"strings"
)

func GetFullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

func TrimScore(score float32) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", score), "0"), ".")
}

func NormalizeString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
