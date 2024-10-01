package user

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *userService) Get(id int) (*model.User, error) {
	res, err := s.storage.GetUser(id)
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}

	return res, nil
}
