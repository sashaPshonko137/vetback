package diagnosis

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *diagnosisService) GetMany() ([]model.Diagnosis, error) {
	diagnoses, err := s.storage.GetManyDiagnosis()
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}

	return diagnoses, nil
}
