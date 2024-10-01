package user

import (
	"golang.org/x/crypto/bcrypt"
	customError "vetback/internal/error"
)

func (s *userService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), s.config.Cost)
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}

	return string(hash), nil
}
