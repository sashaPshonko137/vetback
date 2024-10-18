package user

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
	customError "vetback/internal/error"
)

// @Summary Get session info
// @Tags auth
// @Description Get your session info
// @Accept json
// @Produce json
// @Router /session-info [get]
func (a *userApi) NewGetSessionInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := a.getToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		info, err := a.service.GetSessionInfo(token)
		if err != nil {
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		render.JSON(w, r, info)
	}
}
