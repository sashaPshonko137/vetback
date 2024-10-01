package diagnosis

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

// @Summary update diagnosis
// @Tags diagnosis
// @Description update diagnosis
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param diagnosisInfo body model.DiagnosisToUpdate true "diagnosis"
// @Router /diagnosis/{id} [put]
func (a *diagnosisApi) NewUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.userApi.GetSessionInfo(w, r)
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

		diagnosis, err := a.service.Get(id)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if diagnosis == nil {
			http.Error(w, customError.DiagnosisNotFound.Error(), http.StatusNotFound)
			return
		}
		if diagnosis.DoctorId != info.UserId {
			http.Error(w, "нельзя редактировать информацию чужого диагноза", http.StatusForbidden)
			return
		}
		var req model.DiagnosisToUpdate
		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)
				return
			}
			http.Error(w, "неверный формат запроса", http.StatusBadRequest)
			return
		}

		if err := validator.ValidateDiagnosisUpdate(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = a.service.Update(id, req)
		if err != nil {
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
