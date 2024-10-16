package animal

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

// @Summary Delete animal
// @Tags animal
// @Description Delete an animal by ID
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /animal/{id} [delete]
func (a *animalApi) NewDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.userApi.GetSessionInfo(r)
		if err != nil {
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, customError.MustBeDoctor.Error(), http.StatusUnauthorized)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, customError.MustBeDoctor.Error(), http.StatusUnauthorized)
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
		}

		if animal == nil {
			http.Error(w, customError.AnimalNotFound.Error(), http.StatusNotFound)
			return
		}

		if animal.OwnerId != info.UserId {
			http.Error(w, "нельзя удалять чужое животное", http.StatusForbidden)
			return
		}
		err = a.service.Delete(id)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		render.JSON(w, r, model.NewResponse("Запись удалена"))
	}
}
