package diagnosis

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *diagnosisService) Update(id int, model model.DiagnosisToUpdate) error {
	err := s.storage.UpdateDiagnosis(id, model)

	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}

	return nil
}
