package diagnosis

import (
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/validator"
)

// @Summary create diagnosis
// @Tags diagnosis
// @Description create diagnosis
// @Accept json
// @Produce json
// @Param diagnosis body model.SwaggerDiagnosis true "diagnosis"
// @Router /diagnosis [post]
func (a *diagnosisApi) NewCreate() http.HandlerFunc {
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

		if info.Role != "doctor" {
			http.Error(w, customError.MustBeDoctor.Error(), http.StatusForbidden)
			return
		}
		var req model.DiagnosisInfo

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

		if err := validator.ValidateDiagnosis(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		if err := a.service.Create(req); err != nil {
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
