package appointment

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

// @Summary Delete appointment
// @Tags appointment
// @Description Delete an appointment by ID
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Router /appointment/{id} [delete]
func (a *appointmentApi) NewDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.userApi.GetSessionInfo(w, r)
		if err != nil {
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, customError.MustBeDoctor.Error(), http.StatusUnauthorized)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, customError.MustBeDoctor.Error(), http.StatusUnauthorized)
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
		appointment, err := a.service.Get(id)
		if err != nil {
			if errors.Is(err, customError.AppointmentNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			a.logger.Error(err)
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		if appointment == nil {
			http.Error(w, customError.AppointmentNotFound.Error(), http.StatusNotFound)
			return
		}

		if appointment.DoctorId != info.UserId {
			http.Error(w, "нельзя удалять чужую запись", http.StatusForbidden)
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
