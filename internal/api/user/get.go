package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"vetback/internal/converter"
)

// @Summary Get user by ID
// @Tags user
// @Description Get an user by ID
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /user/{id} [get]
func (a *userApi) NewGet() http.HandlerFunc {
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

		user, err := a.service.Get(id)
		if err != nil {
			http.Error(w, "ошибка получения записи из базы данных", http.StatusInternalServerError)
			a.logger.Error(err)

			return
		}
		if user == nil {
			http.Error(w, "запись не найдена", http.StatusNotFound)
			return
		}

		res := converter.UserToSafe(*user)

		render.JSON(w, r, &res)
	}
}
