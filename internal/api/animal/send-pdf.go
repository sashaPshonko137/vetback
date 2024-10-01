package animal

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"strconv"
	customError "vetback/internal/error"
	"vetback/internal/model"
	"vetback/internal/service"
)

// @Summary Generate and send PDF file
// @Tags animal
// @Description Generate and send PDF file
// @Produce application/pdf
// @Param id path int true "id"
// @Router /story/{id} [get]
func (a *animalApi) SendPDF(appointmentService service.AppointmentService, treatmentService service.TreatmentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "неверный формат идентификатора", http.StatusBadRequest)
		}

		filePath, err := a.service.GeneratePDF(id, appointmentService, treatmentService)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		pdfFile, err := os.Open(filePath)
		if err != nil {
			a.logger.Error(err)
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}
		defer pdfFile.Close()

		stat, err := pdfFile.Stat()
		if err != nil {
			http.Error(w, customError.InternalServerError.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename="+stat.Name())
		w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))

		// Передаем содержимое файла клиенту

		http.ServeFile(w, r, filePath)
		render.JSON(w, r, model.NewResponse("PDF файл успешно отправлен"))
	}
}
