package treatment

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

// @Summary Delete treatment
// @Tags treatment
// @Description Delete an treatment by ID
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /treatment/{id} [delete]
func (a *treatmentApi) NewDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		if idStr == "" {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)
			return
		}

		treatment, err := a.service.Get(id)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			a.logger.Error(err)
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}
		if treatment == nil {
			http.Error(w, "запись не найдена", http.StatusNotFound)
			return
		}

		info, err := a.userApi.GetSessionInfo(w, r)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, "нельзя удалить чужую запись", http.StatusForbidden)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, "нельзя удалить чужую запись", http.StatusForbidden)
				return
			}

			a.logger.Error(err)
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		if treatment.DoctorId != info.UserId {
			http.Error(w, "нельзя удалить чужую запись", http.StatusForbidden)
			return
		}
		err = a.service.Delete(id)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			a.logger.Error(err)
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, model.NewResponse("Запись удалена"))
	}
}
