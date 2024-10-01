package appointment

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *appointmentService) GetByAnimal(id int) ([]model.Appointment, error) {
	res, err := s.storage.GetAppointmentsByAnimal(id)
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}

	return res, nil
}
