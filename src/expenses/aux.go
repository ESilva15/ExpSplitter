package expenses

import (
	"strconv"
)

func ParseID(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 16)
}
