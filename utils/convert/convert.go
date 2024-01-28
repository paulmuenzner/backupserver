package convert

import (
	"fmt"
	"strconv"
)

func ConvertStringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to int: %v", err)
	}
	return i, nil
}
