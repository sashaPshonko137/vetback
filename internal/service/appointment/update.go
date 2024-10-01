package appointment

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *appointmentService) Update(id int, model model.AppointmentToUpdate) error {
	if model.OwnerId > 0 {
		owner, err := s.userService.Get(model.OwnerId)
		if err != nil {
			return customError.InternalServerError
		}
		if owner == nil {
			return customError.OwnerNotFound
		}
		if owner.Role != "owner" {
			return customError.OwnerNotFound
		}
	}
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
	err := s.storage.UpdateAppointment(id, model)

	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}

	return nil
}
