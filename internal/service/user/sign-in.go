package user

import (
	"golang.org/x/crypto/bcrypt"
	"vetback/internal/converter"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *userService) SignIn(info model.UserToLogin) (string, error) {
	user, err := s.GetUserByPhone(info.PhoneNumber)
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}
	if user == nil {
		return "", customError.FailedLogin
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(info.Password))
	if err != nil {
		return "", customError.FailedLogin
	}

	toSession := converter.UserToSession(*user)

	token, err := s.GenerateToken(toSession)
	if err != nil {
		return "", err
	}

	return token, nil
}
