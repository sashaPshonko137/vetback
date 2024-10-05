package validator

import (
	"errors"
	"time"
	"vetback/internal/model"
)

func ValidateTreatment(model model.TreatmentInfo) error {
	if model.DoctorId == 0 {
		return errors.New("необходимо указать врача")
	}
	if model.Name == "" {
		return errors.New("необходимо указать кличку")
	}
	if len(model.Name) > 50 {
		return errors.New("длина клички не должна превышать 50 символов")
	}
	if model.AnimalId == 0 {
		return errors.New("необходимо указать животное")
	}
	if model.DiagnosisId == 0 {
		return errors.New("необходимо указать диагноз")
	}
	if model.Finish == "" {
		return errors.New("необходимо указать дату окончания лечения")
	}
	if !IsValidFinish(model.Finish, "2006-01-02") {
		return errors.New("дата окончания лечения не может быть раньше текущей (дата окончания лечения должна быть в формате 2006-01-02)")
	}
	if model.Price == 0 {
		return errors.New("необходимо указать цену лечения")
	}
	return nil
}

func ValidateTreatmentUpdate(model model.TreatmentToUpdate) error {
	if model.Finish != "" {
		if !IsValidFinish(model.Finish, "2006-01-02") {
			return errors.New("дата окончания лечения должна быть в формате 2006-01-02 и не может быть раньше текущей даты")
		}
	}
	return nil
}

func IsValidFinish(dateStr, layout string) bool {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return false
	}
	return date.After(time.Now()) // Дата должна быть после текущей
}
