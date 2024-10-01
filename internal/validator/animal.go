package validator

import (
	"errors"
	"time"
	"vetback/internal/model"
)

func ValidateAnimal(model model.AnimalInfo) error {
	if model.DoctorId == 0 {
		return errors.New("необходимо указать врача")
	}
	if model.DiagnosisId == 0 {
		return errors.New("необходимо указать диагноз")
	}
	if model.Name == "" {
		return errors.New("необходимо указать кличку")
	}
	if len(model.Name) > 50 {
		return errors.New("длина клички не должна превышать 50 символов")
	}
	if model.Birthdate == "" {
		return errors.New("необходимо указать дату рождения")
	}
	if IsValidDate(model.Birthdate, "2006-01-02") == false {
		return errors.New("дата рождения должна быть в формате 2006-01-02")
	}
	if model.Breed == "" {
		return errors.New("необходимо указать породу")
	}
	if len(model.Breed) > 50 {
		return errors.New("длина породы не должна превышать 50 символов")
	}
	if model.Sex == "" {
		return errors.New("необходимо указать пол")
	}
	if len(model.Sex) > 50 {
		return errors.New("длина пола не должна превышать 50 символов")
	}

	return nil
}

func ValidateAnimalUpdate(model model.AnimalToUpdate) error {
	if len(model.Name) > 50 {
		return errors.New("длина клички не должна превышать 50 символов")
	}
	if len(model.Breed) > 50 {
		return errors.New("длина породы не должна превышать 50 символов")
	}
	if len(model.Sex) > 50 {
		return errors.New("длина пола не должна превышать 50 символов")
	}

	return nil
}

func IsValidDate(dateStr, layout string) bool {
	_, err := time.Parse(layout, dateStr)
	return err == nil
}
