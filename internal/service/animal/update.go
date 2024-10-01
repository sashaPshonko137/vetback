package animal

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *animalService) Update(id int, model model.AnimalToUpdate) error {
	if model.DoctorId > 0 {
		doctor, err := s.userService.Get(model.DoctorId)
		if err != nil {
			return customError.DoctorNotFound
		}
		if doctor == nil {
			return customError.DoctorNotFound
		}
		if doctor.Role != "doctor" {
			return customError.DoctorNotFound
		}
	}
	if model.DiagnosisId > 0 {
		diagnosis, err := s.diagnosisService.Get(model.DiagnosisId)
		if err != nil {
			return customError.DiagnosisNotFound
		}
		if diagnosis == nil {
			return customError.DiagnosisNotFound
		}
	}

	err := s.storage.UpdateAnimal(id, model)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}
	return nil
}
