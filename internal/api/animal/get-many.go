package animal

import (
	"github.com/go-chi/render"
	"net/http"
)

// @Summary Get many animals
// @Tags animal
// @Description Get many animals
// @Accept json
// @Produce json
// @Router /animal [get]
func (s *animalApi) NewGetMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		animals, err := s.service.GetMany()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, animals)
	}
}
