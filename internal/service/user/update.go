package user

import (
	"vetback/internal/converter"
	"vetback/internal/model"
)

func (s *userService) Update(id int, model model.UserToUpdate) error {
	hash, err := s.hashPassword(model.Password)
	if err != nil {
		return err
	}
	toStorage := converter.UserToUpdate(model, hash)
	err = s.storage.UpdateUser(id, toStorage)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}
