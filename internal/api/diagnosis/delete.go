package diagnosis

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

// @Summary Delete diagnosis
// @Tags diagnosis
// @Description Delete an diagnosis by ID
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /diagnosis/{id} [delete]
func (a *diagnosisApi) NewDelete() http.HandlerFunc {
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

		diagnosis, err := a.service.Get(id)
		if err != nil {
			if errors.Is(err, customError.DiagnosisNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if diagnosis == nil {
			http.Error(w, customError.DiagnosisNotFound.Error(), http.StatusNotFound)
			return
		}

		if info.UserId != diagnosis.DoctorId {
			http.Error(w, "нельзя удалять чужой диагноз", http.StatusForbidden)
			return
		}
		err = a.service.Delete(id)
		if err != nil {
			if errors.Is(err, customError.DiagnosisNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		render.JSON(w, r, model.NewResponse("Запись удалена"))
	}
}
