package user

import (
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/validator"
)

// @Summary Sign up
// @Tags auth
// @Description registration
// @Accept json
// @Produce json
// @Param user body model.UserInfo true "user info"
// @Success 200 {object} model.UserInfo
// @Router /sign-up [post]
func (a *userApi) NewSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.UserInfo
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)

				return
			}
			http.Error(w, "некорректный формат запроса", http.StatusBadRequest)

			return
		}

		if err := validator.ValidateUser(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}
		token, err := a.service.SignUp(req)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if errors.Is(err, customError.UserAlreadyExist) {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			if errors.Is(err, customError.NoSpecialization) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		//a.logger.Info(token.Raw)
		a.setToken(w, token)

		render.JSON(w, r, model.NewResponse("вы зарегистрировались"))
	}
}
