package diagnosis

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *diagnosisService) Create(model model.DiagnosisInfo) error {
	err := s.storage.CreateDiagnosis(model)
	if err != nil {
		return customError.InternalServerError
	}
	return nil
}
