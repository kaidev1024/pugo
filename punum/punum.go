package punum

import (
	"fmt"
	"strconv"
)

// parseInt8 parses s into an int8. If the parsed value is outside [min, max], an error is returned.
func parseInt8(s string, min, max int8) (int8, error) {
	v, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("punum: parseInt8: invalid value %q: %w", s, err)
	}
	n := int8(v)
	if n < min {
		return 0, fmt.Errorf("punum: parseInt8: value %d is less than min %d", n, min)
	}
	if n > max {
		return 0, fmt.Errorf("punum: parseInt8: value %d is greater than max %d", n, max)
	}
	return n, nil
}

func ParseSubsecs(s string) (int8, error) {
	return parseInt8(s, 0, 99)
}

func ParseSeconds(s string) (int8, error) {
	return parseInt8(s, 0, 59)
}

func ParseMinutes(s string) (int8, error) {
	return parseInt8(s, 0, 59)
}

func ParseHours(s string) (int8, error) {
	return parseInt8(s, 0, 127)
}
