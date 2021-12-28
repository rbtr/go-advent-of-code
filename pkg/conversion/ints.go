package conversion

import (
	"strconv"

	"github.com/pkg/errors"
)

func ParseInt(in []byte) (int, error) {
	i, err := strconv.Atoi(string(in))
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse string to int")
	}
	return i, nil
}
