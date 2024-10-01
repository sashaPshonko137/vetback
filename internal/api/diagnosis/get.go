package diagnosis

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

// @Summary Get diagnosis
// @Tags diagnosis
// @Description Get an diagnosis by ID
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /diagnosis/{id} [get]
func (a *diagnosisApi) NewGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "diagnosisId")

		if idStr == "" {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)

			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)

			return
		}

		res, err := a.service.Get(id)
		if err != nil {
			http.Error(w, "ошибка получения записи из базы данных", http.StatusInternalServerError)
			a.logger.Error(err)

			return
		}
		if res == nil {
			http.Error(w, "запись не найдена", http.StatusNotFound)
			return
		}

		render.JSON(w, r, res)
	}
}
