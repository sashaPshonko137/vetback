package appointment

import (
	"github.com/go-chi/render"
	"net/http"
)

// @Summary Get many appointments
// @Tags appointment
// @Description Get many appointments
// @Accept json
// @Produce json
// @Router /appointment [get]
func (a *appointmentApi) NewGetMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := a.service.GetMany()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, res)
	}
}
