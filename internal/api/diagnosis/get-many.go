package diagnosis

import (
	"github.com/go-chi/render"
	"net/http"
)

// @Summary Get much diagnosis
// @Tags diagnosis
// @Description Get much diagnosis
// @Accept json
// @Produce json
// @Router /diagnosis [get]
func (a *diagnosisApi) NewGetMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := a.service.GetMany()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, res)
	}
}
