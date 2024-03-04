package helper

import "strconv"

func ConvStringToUint(s string) (result uint, err error) {
	parsed, err := strconv.ParseUint(s, 10, 64)
	result = uint(parsed)

	return
}
