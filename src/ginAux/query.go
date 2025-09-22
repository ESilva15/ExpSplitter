package ginaux

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QueryIntArray will return an array of ints of a given query paramenter.
func QueryIntArray(c *gin.Context, param string) ([]int32, error) {
	selCategories := c.QueryArray(param)
	var newArr []int32
	if len(selCategories) > 0 {
		newArr = make([]int32, len(selCategories))

		for k, val := range selCategories {
			id, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				log.Printf("error converting id [%s]: %v", val, err)
				return nil, err
			}

			newArr[k] = int32(id)
		}
	}

	return newArr, nil
}
