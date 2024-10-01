package appointment

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *appointmentService) GetMany() ([]model.Appointment, error) {
	res, err := s.storage.GetManyAppointments()
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}

	return res, nil
}
