package punum

import (
	"fmt"
	"strconv"
)

// parseInt8 parses s into an int8. If the parsed value is outside [min, max], an error is returned.
func parseInt8(s, name string, min, max int8) (int8, error) {
	v, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: %w", name, s, err)
	}
	n := int8(v)
	if n < min {
		return 0, fmt.Errorf("%s is less than min %d", name, min)
	}
	if n > max {
		return 0, fmt.Errorf("%s is greater than max %d", name, max)
	}
	return n, nil
}

func ParseSubsecs(s string) (int8, error) {
	if len(s) == 1 {
		s += "0"
	}
	return parseInt8(s, "value", 0, 99)
}

func ParseSeconds(s string) (int8, error) {
	return parseInt8(s, "seconds", 0, 59)
}

func ParseMinutes(s string) (int8, error) {
	return parseInt8(s, "minutes", 0, 59)
}

func ParseHours(s string) (int8, error) {
	return parseInt8(s, "hours", 0, 127)
}
