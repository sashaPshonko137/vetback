package user

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *userService) Create(info model.UserToStorage) (int, error) {
	user, err := s.storage.GetUserByPhone(info.PhoneNumber)
	if err != nil {
		s.logger.Error(err)
		return 0, customError.InternalServerError
	}
	if user != nil {
		return 0, customError.UserAlreadyExist
	}

	if !checkSpecialization(info) {
		return 0, customError.NoSpecialization
	}

	id, err := s.storage.CreateUser(info)
	if err != nil {
		return 0, customError.InternalServerError
	}
	return id, nil
}

func checkSpecialization(info model.UserToStorage) bool {
	if info.Role == "doctor" && info.Specialization == "" {
		return false
	}
	return true
}
