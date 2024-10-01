package user

import (
	"net/http"
	"vetback/internal/model"
)

func (a *userApi) GetSessionInfo(w http.ResponseWriter, r *http.Request) (*model.Claims, error) {
	token, err := a.getToken(w, r)
	if err != nil {
		return nil, err
	}

	info, err := a.service.GetSessionInfo(token)
	if err != nil {
		return nil, err
	}

	return info, nil
}
