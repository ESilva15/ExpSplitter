package ginaux

import (
	"log"
	"time"

	"github.com/ESilva15/expenses/expenses/repo"
	"github.com/gin-gonic/gin"
)

// ExpenseFilterFromQuery returns the ExpFilter with the current gin context
// data.
func ExpenseFilterFromQuery(c *gin.Context) (repo.ExpFilter, error) {
	var err error
	var start, end time.Time
	filter := repo.NewExpFilter()

	startDate := c.Query("range-start")
	if startDate != "" {
		startDate += " 00:00:00"
		start, err = time.ParseInLocation("02-Jan-2006 15:04:05", startDate, time.UTC)
		if err != nil {
			log.Printf("error startDate: %v", err)
			return repo.ExpFilter{}, nil
		}
		filter.Start = &start
	}

	endDate := c.Query("range-end")
	if endDate != "" {
		endDate += " 00:00:00"
		end, err = time.ParseInLocation("02-Jan-2006 15:04:05", endDate, time.UTC)
		if err != nil {
			log.Printf("error endDate: %v", err)
			return repo.ExpFilter{}, nil
		}
		filter.End = &end
	}

	filter.CatIDs, err = QueryIntArray(c, "cat-dropdown[]")
	if err != nil {
		log.Printf("Failed to retrieve IDs from cat-dropdown[] %+v", err)
		return repo.ExpFilter{}, nil
	}

	filter.StoreIDs, err = QueryIntArray(c, "store-dropdown[]")
	if err != nil {
		log.Printf("Failed to retrieve IDs from store-dropdown[] %+v", err)
		return repo.ExpFilter{}, nil
	}

	filter.TypeIDs, err = QueryIntArray(c, "type-dropdown[]")
	if err != nil {
		log.Printf("Failed to retrieve IDs from type-dropdown[] %+v", err)
		return repo.ExpFilter{}, nil
	}

	return filter, nil
}
