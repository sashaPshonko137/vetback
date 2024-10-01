package animal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

// @Summary Get animal
// @Tags animal
// @Description Get an animal by ID
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /animal/{id} [get]
func (a *animalApi) NewGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)
			return
		}

		res, err := a.service.Get(id)
		if err != nil {
			http.Error(w, "ошибка получения записи из базы данных", http.StatusInternalServerError)

			return
		}
		if res == nil {
			http.Error(w, "запись не найдена", http.StatusNotFound)
			return
		}

		render.JSON(w, r, &res)
	}
}
