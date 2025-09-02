package expenses

import (
	"strconv"
)

func ParseID(s string) (int32, error) {
	res, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return -1, err
	}

	return int32(res), nil
}
