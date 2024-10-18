package user

import (
	"net/http"
	"time"
	customError "vetback/internal/error"
)

func (a *userApi) setToken(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: false,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)
}

func (a *userApi) getToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", customError.TokenNotFound
	}
	return cookie.Value, nil
}
