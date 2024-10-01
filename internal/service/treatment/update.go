package treatment

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *treatmentService) Update(id int, model model.TreatmentToUpdate) error {
	if model.AnimalId > 0 {
		animal, err := s.animalService.Get(model.AnimalId)
		if err != nil {
			return customError.InternalServerError
		}
		if animal == nil {
			return customError.AnimalNotFound
		}
	}
	if model.DiagnosisId > 0 {
		diagnosis, err := s.diagnosisService.Get(model.DiagnosisId)
		if err != nil {
			return customError.InternalServerError
		}
		if diagnosis == nil {
			return customError.DiagnosisNotFound
		}
	}
	err := s.storage.UpdateTreatment(id, model)

	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}
