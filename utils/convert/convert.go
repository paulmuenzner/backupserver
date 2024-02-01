package convert

import (
	"fmt"
	"strconv"
)

// String to int
func ConvertStringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to int: %v", err)
	}
	return i, nil
}

// Error to string
func ErrorAsString(err error) string {
	if err == nil {
		return "No error detected."
	}
	return fmt.Sprintf("%v", err)
}
