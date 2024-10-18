package appointment

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/validator"
)

// @Summary create appointment
// @Tags appointment
// @Description create appointment
// @Accept json
// @Produce json
// @Param appointmentInfo body model.SwaggerAppointment true "appointment"
// @Router /appointment [post]
func (s *appointmentApi) NewCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := s.userApi.GetSessionInfo(r)
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
		if info.Role != "doctor" {
			http.Error(w, customError.MustBeDoctor.Error(), http.StatusForbidden)
			return
		}
		var req model.AppointmentInfo

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)

				return
			}
			http.Error(w, "ошибка парсинга запроса", http.StatusBadRequest)

			return
		}

		req.DoctorId = info.UserId

		if err := validator.ValidateAppointment(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		if err := s.service.Create(req); err != nil {
			if errors.Is(err, customError.AnimalNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.DiagnosisNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.DoctorNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
			}
			if errors.Is(err, customError.OwnerNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.NotOwner) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if errors.Is(err, customError.InternalServerError) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		render.JSON(w, r, struct {
			Message string `json:"message"`
		}{
			Message: "Запись добавлена",
		})
	}
}
