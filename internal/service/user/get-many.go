package user

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *userService) GetMany() ([]model.SafeUser, error) {
	users, err := s.storage.GetManyUsers()
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}

	return users, nil
}
