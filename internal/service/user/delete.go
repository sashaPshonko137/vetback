package user

import customError "vetback/internal/error"

func (s *userService) Delete(id int) error {
	user, err := s.storage.GetUser(id)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}
	if user == nil {
		return customError.AccountNotFound
	}
	err = s.storage.DeleteUser(id)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}
	return nil
}
