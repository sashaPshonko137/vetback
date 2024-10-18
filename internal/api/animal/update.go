package animal

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"strconv"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/validator"
)

// @Summary update
// @Tags animal
// @Description update animal
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param animalInfo body model.AnimalToUpdate true "animal"
// @Router /animal/{id} [put]
func (a *animalApi) NewUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.userApi.GetSessionInfo(r)
		if err != nil {
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, customError.MustBeOwner.Error(), http.StatusForbidden)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, customError.MustBeOwner.Error(), http.StatusForbidden)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)
			return
		}

		animal, err := a.service.Get(id)
		if err != nil {
			if errors.Is(err, customError.AnimalNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			a.logger.Error(err)
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		if animal == nil {
			http.Error(w, customError.AnimalNotFound.Error(), http.StatusNotFound)
			return
		}

		if animal.OwnerId != info.UserId {
			http.Error(w, "нельзя редактировать информацию чужого животного", http.StatusForbidden)
			return
		}
		var req model.AnimalToUpdate
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)
				return
			}
			http.Error(w, "неверный формат запроса", http.StatusBadRequest)
			return
		}

		if err := validator.ValidateAnimalUpdate(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = a.service.Update(id, req)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if errors.Is(err, customError.DoctorNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.DiagnosisNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}

		render.JSON(w, r, model.NewResponse("информация обновлена"))
	}
}
