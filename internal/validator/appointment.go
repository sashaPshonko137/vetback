package validator

import (
	"errors"
	"vetback/internal/model"
)

func ValidateAppointment(info model.AppointmentInfo) error {
	if info.OwnerId == 0 {
		return errors.New("необходимо указать хозяина")
	}
	if info.AnimalId == 0 {
		return errors.New("необходимо указать животное")
	}
	if info.Date == "" {
		return errors.New("необходимо указать дату приема")
	}
	if IsValidDate(info.Date, "02-01-2006 15:04") == false {
		return errors.New("дата приема должна быть в формате 02-01-2006 15:04")
	}
	if !IsValidFinish(info.Date, "02-01-2006 15:04") {
		return errors.New("дата приема не может быть раньше текущей (дата приема должна быть в формате 02-01-2006 15:04)")
	}
	if info.Status == "" {
		return errors.New("необходимо указать статус")
	}
	if len(info.Status) > 50 {
		return errors.New("длина статуса не должна превышать 50 символов")
	}
	if info.Reason == "" {
		return errors.New("необходимо указать причину приема")
	}
	if len(info.Reason) > 200 {
		return errors.New("длина причины не должна превышать 50 символов")
	}
	return nil
}

func ValidateAppointmentUpdate(model model.AppointmentToUpdate) error {
	if len(model.Status) > 50 {
		return errors.New("длина статуса не должна превышать 50 символов")
	}
	if len(model.Reason) > 200 {
		return errors.New("длина причины не должна превышать 50 символов")
	}
	if model.Date != "" && IsValidDate(model.Date, "2006-01-02 15:04") == false {
		return errors.New("дата приема должна быть в формате 2006-01-02 15:04")
	}
	return nil
}
