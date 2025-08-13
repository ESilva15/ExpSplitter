package expenses

import (
	mod "expenses/expenses/models"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

func (s *Service) GetAllExpenses() ([]mod.Expense, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return []mod.Expense{}, err
	}
	defer tx.Rollback()

	expenses, err := mod.GetAllExpenses(tx)

	return expenses, tx.Commit()
}

func (s *Service) GetExpensesRanged(startDate string, endDate string) ([]mod.Expense, error) {
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

	tx, err := s.DB.Begin()
	if err != nil {
		return []mod.Expense{}, err
	}
	defer tx.Rollback()

	expenses, err := mod.GetExpensesRange(tx, start, end)

	return expenses, tx.Commit()
}

func (s *Service) GetExpense(id int64) (mod.Expense, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return mod.Expense{}, err
	}
	defer tx.Rollback()

	expense, err := mod.GetExpense(tx, id)
	if err != nil {
		return mod.Expense{}, err
	}

	return expense, tx.Commit()
}

func (s *Service) LoadExpenseShares(e *mod.Expense) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = e.GetShares(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) LoadExpensePayments(e *mod.Expense) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = e.GetPayments(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) DeleteExpense(id int64) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	expense := mod.Expense{
		ExpID: id,
	}

	err = expense.Delete()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func ExpenseTotalPayed(exp *mod.Expense) decimal.Decimal {
	total := decimal.NewFromFloat(0.0)
	for _, p := range exp.Payments {
		log.Println(total)
		total = total.Add(p.PayedAmount)
	}

	return total
}

func (s *Service) NewExpense(exp mod.Expense) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Figure out if its paid off or not by adding the existing payments
	exp.PaidOff = exp.Value.Equal(ExpenseTotalPayed(&exp))

	// Figure out if its evenly shared by the people associated to it

	err = exp.Insert(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) UpdateExpense(exp mod.Expense) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = exp.Update(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
