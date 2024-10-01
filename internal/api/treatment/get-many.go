package treatment

import (
	"github.com/go-chi/render"
	"net/http"
)

// @Summary Get much treatment
// @Tags treatment
// @Description Get much treatment
// @Accept json
// @Produce json
// @Router /treatment [get]
func (a *treatmentApi) NewGetMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := a.service.GetMany()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, res)
	}
}
