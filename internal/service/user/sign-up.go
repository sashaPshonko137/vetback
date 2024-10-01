package user

import (
	"errors"
	"vetback/internal/converter"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *userService) SignUp(info model.UserInfo) (string, error) {
	hash, err := s.hashPassword(info.Password)
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}

	toStorage := converter.UserToStorage(info, hash)
	id, err := s.Create(toStorage)
	if err != nil {
		if errors.Is(err, customError.UserAlreadyExist) {
			return "", err
		}
		if errors.Is(err, customError.NoSpecialization) {
			return "", err
		}
		s.logger.Error(err)
		return "", customError.InternalServerError
	}

	toSession := converter.UserInfoToSession(info, id)
	token, err := s.GenerateToken(toSession)
	if err != nil {
		return "", customError.InternalServerError
	}

	return token, nil
}
