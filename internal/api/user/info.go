package user

import (
	"net/http"
	"vetback/internal/model"
)

func (a *userApi) GetSessionInfo(r *http.Request) (*model.Claims, error) {
	token, err := a.getToken(r)
	if err != nil {
		return nil, err
	}

	info, err := a.service.GetSessionInfo(token)
	if err != nil {
		return nil, err
	}

	return info, nil
}
