package user

import (
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

// @Summary Sign in
// @Tags auth
// @Description Sign in
// @Accept json
// @Produce json
// @Param user body model.UserToLogin true "user"
// @Router /sign-in [post]
func (a *userApi) NewSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.UserToLogin

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)
				return
			}
			http.Error(w, "некорректный формат запроса", http.StatusBadRequest)
			return
		}

		token, err := a.service.SignIn(req)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if errors.Is(err, customError.FailedLogin) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
		}

		a.setToken(w, token)

		render.JSON(w, r, model.NewResponse("вы вошли в аккаунт"))
	}
}
