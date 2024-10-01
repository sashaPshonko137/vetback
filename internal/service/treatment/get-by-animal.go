package treatment

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *treatmentService) GetManyByAnimal(id int) ([]model.Treatment, error) {
	res, err := s.storage.GetManyTreatmentsByAnimal(id)
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}

	return res, nil
}
