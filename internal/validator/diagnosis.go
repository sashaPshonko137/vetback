package validator

import (
	"errors"
	"vetback/internal/model"
)

func ValidateDiagnosis(model model.DiagnosisInfo) error {
	if model.Name == "" {
		return errors.New("необходимо указать наименование")
	}
	if len(model.Name) > 50 {
		return errors.New("длина наименования не должна превышать 50 символов")
	}
	return nil
}

func ValidateDiagnosisUpdate(model model.DiagnosisToUpdate) error {
	if len(model.Name) > 50 {
		return errors.New("длина наименования не должна превышать 50 символов")
	}
	return nil
}
