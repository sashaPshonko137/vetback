package user

import (
	"vetback/internal/model"
)

func (s *userService) GetSessionInfo(token string) (*model.Claims, error) {
	info, err := s.getInfoFromToken(token)
	if err != nil {
		return nil, err
	}

	return info, nil
}
