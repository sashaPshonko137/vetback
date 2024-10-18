package user

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

// @Summary Delete your account
// @Tags user
// @Description Delete your account
// @Accept json
// @Produce json
// @Router /user [delete]
func (a *userApi) NewDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.GetSessionInfo(r)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
		}
		err = a.service.Delete(info.UserId)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if errors.Is(err, customError.AccountNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}
		render.JSON(w, r, model.NewResponse("вы удалили аккаунт"))
	}
}
