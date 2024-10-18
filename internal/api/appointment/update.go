package appointment

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"strconv"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/validator"
)

// @Summary update appointment
// @Tags appointment
// @Description update appointment
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param appointmentInfo body model.AppointmentToUpdate true "appointment"
// @Router /appointment/{id} [put]
func (a *appointmentApi) NewUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.userApi.GetSessionInfo(r)
		if err != nil {
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, customError.MustBeOwner.Error(), http.StatusForbidden)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, customError.MustBeOwner.Error(), http.StatusForbidden)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if info.Role != "doctor" {
			http.Error(w, customError.MustBeDoctor.Error(), http.StatusForbidden)
			return
		}
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)
			return
		}

		appointment, err := a.service.Get(id)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if appointment == nil {
			http.Error(w, "запись не найдена", http.StatusNotFound)
			return
		}

		if appointment.DoctorId != info.UserId {
			http.Error(w, "нельзя редактировать чужую запись", http.StatusForbidden)
			return
		}

		var req model.AppointmentToUpdate

		err = render.DecodeJSON(r.Body, &req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)
				return
			}
			http.Error(w, "неверный формат запроса", http.StatusBadRequest)
			return
		}

		if err := validator.ValidateAppointmentUpdate(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = a.service.Update(id, req)
		if err != nil {
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if errors.Is(err, customError.OwnerNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.AnimalNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.DiagnosisNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}

		render.JSON(w, r, model.NewResponse("запись обновлена"))
	}
}
