package treatment

import (
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/validator"
)

// @Summary create treatment
// @Tags treatment
// @Description create treatment
// @Accept json
// @Produce json
// @Param treatmentInfo body model.SwaggerTreatment true "treatment"
// @Router /treatment [post]
func (a *treatmentApi) NewCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := a.userApi.GetSessionInfo(r)
		if err != nil {
			if errors.Is(err, customError.InvalidToken) {
				http.Error(w, customError.MustBeDoctor.Error(), http.StatusUnauthorized)
				return
			}
			if errors.Is(err, customError.TokenNotFound) {
				http.Error(w, customError.MustBeDoctor.Error(), http.StatusUnauthorized)
				return
			}
		}

		if info.Role != "doctor" {
			http.Error(w, customError.MustBeDoctor.Error(), http.StatusForbidden)
			return
		}
		var req model.TreatmentInfo

		err = render.DecodeJSON(r.Body, &req)
		req.DoctorId = info.UserId
		if err != nil {
			if errors.Is(err, io.EOF) {
				http.Error(w, "пустой запрос", http.StatusBadRequest)

				return
			}
			http.Error(w, "ошибка парсинга запроса", http.StatusBadRequest)

			return
		}
		if err := validator.ValidateTreatment(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		if err := a.service.Create(req); err != nil {
			if errors.Is(err, customError.AnimalNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, customError.DiagnosisNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
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
