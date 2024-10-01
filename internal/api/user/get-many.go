package user

import (
	"github.com/go-chi/render"
	"net/http"
)

// @Summary Get many users
// @Tags user
// @Description Get many users
// @Accept json
// @Produce json
// @Router /user [get]
func (s *userApi) NewGetMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.service.GetMany()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, users)
	}
}
