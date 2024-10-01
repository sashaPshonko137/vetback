package animal

import (
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	_ "vetback/docs"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/validator"
)

// @Summary create animal
// @Tags animal
// @Description create animal
// @Accept json
// @Produce json
// @Param animalInfo body model.SwaggerAnimal true "animal"
// @Router /animal [post]
func (a *animalApi) NewCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.userApi.GetSessionInfo(w, r)

		if err != nil {
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, customError.MustBeOwner.Error(), http.StatusUnauthorized)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, customError.MustBeOwner.Error(), http.StatusUnauthorized)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if info.Role != "owner" {
			http.Error(w, customError.MustBeOwner.Error(), http.StatusForbidden)
			return
		}
		var req model.AnimalInfo
		err = render.DecodeJSON(r.Body, &req)
		req.OwnerId = info.UserId

		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)

				return
			}
			http.Error(w, "неверный формат запроса", http.StatusBadRequest)

			return
		}

		if err := validator.ValidateAnimal(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		if err := a.service.Create(req); err != nil {
			if errors.Is(err, customError.DoctorNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.MustBeOwner) {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			if errors.Is(err, customError.DiagnosisNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		render.JSON(w, r, model.NewResponse("животное создано"))
	}
}
