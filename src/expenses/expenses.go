package expenses

import (
	mod "expenses/expenses/models"
	"log"
	"time"
)

func GetAllExpenses() ([]mod.Expense, error) {
	return mod.GetAllExpenses()
}

func GetExpensesRanged(startDate string, endDate string) ([]mod.Expense, error) {
	startDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", startDate, time.UTC)
	if err != nil {
		log.Printf("error startDate: %v", err)
		return []mod.Expense{}, nil
	}
	start := startDateTime.Unix()

	endDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", endDate, time.UTC)
	if err != nil {
		log.Printf("error endDate: %v", err)
		return []mod.Expense{}, nil
	}
	end := endDateTime.Unix()

	return mod.GetExpensesRange(start, end)
}

func GetExpense(id int64) (mod.Expense, error) {
	return mod.GetExpense(id)
}

func DeleteExpense(id int64) error {
	expense := mod.Expense{
		ExpID: id,
	}

	return expense.Delete()
}
