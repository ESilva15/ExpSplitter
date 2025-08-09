package expenses

import (
	mod "expenses/expenses/models"
)

func (s *Service) MigGoto(id uint) error {
	return mod.Goto(s.DB, id)
}
