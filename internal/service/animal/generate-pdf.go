package animal

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/signintech/gopdf"
	customError "vetback/internal/error"
	"vetback/internal/service"
)

const (
	br   = 20
	br05 = 10
	br1  = 40
)

func (s *animalService) GeneratePDF(animalId int, appointmentService service.AppointmentService, treatmentService service.TreatmentService) (string, error) {
	appointments, err := appointmentService.GetByAnimal(animalId)
	if err != nil {
		return "", customError.InternalServerError
	}

	treatments, err := treatmentService.GetManyByAnimal(animalId)
	if err != nil {
		return "", customError.InternalServerError
	}

	animal, err := s.Get(animalId)
	if err != nil {
		return "", customError.InternalServerError
	}

	owner, err := s.userService.Get(animal.OwnerId)
	if err != nil {
		return "", customError.InternalServerError
	}

	doctor, err := s.userService.Get(animal.DoctorId)
	if err != nil {
		return "", customError.InternalServerError
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err = pdf.AddTTFFont("DejaVu", "internal/generator/fonts/DejaVuSansCondensed.ttf")
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}

	err = pdf.SetFont("DejaVu", "", 12)
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}

	// Заголовок
	title := fmt.Sprint("История болезни ", animal.Name)
	width, err := pdf.MeasureTextWidth(title)
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}
	pdf.SetX((gopdf.PageSizeA4.W - width) / 2)
	pdf.SetFont("DejaVu", "", 16)
	pdf.Cell(nil, title)
	pdf.Br(br1)

	// Информация о владельце и животном
	pdf.SetFont("DejaVu", "", 14)
	pdf.Cell(nil, fmt.Sprintf("Владелец: %s", owner.FullName))
	pdf.Br(br)
	pdf.Cell(nil, fmt.Sprintf("Врач: %s", doctor.FullName))
	pdf.Br(br)
	pdf.Cell(nil, fmt.Sprintf("Специальность врача: %s", doctor.Specialization))
	pdf.Br(br)
	pdf.Cell(nil, fmt.Sprintf("Кличка: %s", animal.Name))
	pdf.Br(br)
	pdf.Cell(nil, fmt.Sprintf("Дата рождения: %s", animal.Birthdate))
	pdf.Br(br)
	pdf.Cell(nil, fmt.Sprintf("Порода: %s", animal.Breed))
	pdf.Br(br)
	pdf.Cell(nil, fmt.Sprintf("Пол: %s", animal.Sex))
	pdf.Br(br)

	// Линия для разделения блоков
	lineStartX := 10.0
	lineEndX := gopdf.PageSizeA4.W - 10.0
	pdf.SetLineWidth(0.5)
	pdf.Line(lineStartX, pdf.GetY(), lineEndX, pdf.GetY())
	pdf.Br(br1)

	// Назначения (табличное представление)
	pdf.SetFont("DejaVu", "", 14)
	pdf.Cell(nil, "Записи на прием:")
	pdf.Br(br)
	pdf.SetFont("DejaVu", "", 10)

	for _, appointment := range appointments {
		pdf.Cell(nil, fmt.Sprintf("Дата: %s", appointment.Date))
		pdf.Br(br05)
		pdf.Cell(nil, fmt.Sprintf("Причина: %s", appointment.Reason))
		pdf.Br(br05)
		pdf.Cell(nil, fmt.Sprintf("Статус: %s", appointment.Status))
		pdf.Br(br05)
		doctor, err := s.userService.Get(appointment.DoctorId)
		if err != nil {
			return "", customError.InternalServerError
		}
		pdf.Cell(nil, fmt.Sprintf("Врач: %s", doctor.FullName))
		pdf.Br(br)

		// Проверка на необходимость добавления новой страницы
		if pdf.GetY() > gopdf.PageSizeA4.H-40 { // Если осталось менее 40 единиц высоты
			pdf.AddPage()
		}
	}

	pdf.SetFont("DejaVu", "", 14)
	pdf.Cell(nil, "Лечения:")
	pdf.Br(br)
	pdf.SetFont("DejaVu", "", 10)

	for _, treatment := range treatments {
		doctor, err := s.userService.Get(treatment.DoctorId)
		if err != nil {
			return "", customError.InternalServerError
		}
		diagnosis, err := s.diagnosisService.Get(treatment.DiagnosisId)
		if err != nil {
			return "", customError.InternalServerError
		}

		pdf.Cell(nil, fmt.Sprintf("Врач: %s", doctor.FullName))
		pdf.Br(br05)
		pdf.Cell(nil, fmt.Sprintf("Диагноз: %s", diagnosis.Name))
		pdf.Br(br05)
		pdf.Cell(nil, fmt.Sprintf("Начало: %s", treatment.Start))
		pdf.Br(br05)
		pdf.Cell(nil, fmt.Sprintf("Конец: %s", treatment.Finish))
		pdf.Br(br)

		if pdf.GetY() > gopdf.PageSizeA4.H-40 { // Если осталось менее 40 единиц высоты
			pdf.AddPage()
		}
	}

	fileUUID := uuid.New().String()
	pdfFileName := fmt.Sprintf("uploads/%s.pdf", fileUUID)

	err = pdf.WritePdf(pdfFileName)
	if err != nil {
		s.logger.Error(err)
		return "", customError.InternalServerError
	}

	return pdfFileName, nil
}
