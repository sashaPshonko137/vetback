package treatment

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *treatmentService) GetMany() ([]model.Treatment, error) {
	res, err := s.storage.GetManyTreatments()
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}

	return res, nil
}
