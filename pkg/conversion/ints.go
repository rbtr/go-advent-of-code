package conversion

import (
	"strconv"
)

func ParseInt(in []byte) (int, error) {
	return strconv.Atoi(string(in))
}

func ParseSignedInts(in []byte) (int, error) {
	str := string(in)
	sign := 0
	switch str[0] {
	case '-':
		sign = -1
	default:
		sign = 1
	}
	i, err := strconv.Atoi(str[1:])
	if err != nil {
		return 0, err
	}
	return i * sign, nil
}
