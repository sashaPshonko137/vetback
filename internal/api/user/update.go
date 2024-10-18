package user

import (
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"vetback/internal/converter"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

// @Summary Update your account
// @Tags user
// @Description Update your account
// @Accept json
// @Produce json
// @Param user body model.UserToUpdate true "user"
// @Router /user [put]
func (a *userApi) NewUpdate() http.HandlerFunc {
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

		var req model.UserToUpdate
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)

				return
			}
			http.Error(w, "некорректный формат запроса", http.StatusBadRequest)

			return
		}

		if req.Specialization != "" && info.Role != "doctor" {
			http.Error(w, "специальность может быть только у врача", http.StatusForbidden)
		}

		err = a.service.Update(info.UserId, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		claims := converter.UpdatedUserToClaims(req, info.UserId)
		token, err := a.service.GenerateToken(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		a.setToken(w, token)

		render.JSON(w, r, model.NewResponse("вы обновили профиль"))
	}
}
