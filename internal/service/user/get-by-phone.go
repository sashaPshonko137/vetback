package user

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *userService) GetUserByPhone(phoneNumber string) (*model.User, error) {
	user, err := s.storage.GetUserByPhone(phoneNumber)
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}
	return user, nil
}
