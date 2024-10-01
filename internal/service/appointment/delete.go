package appointment

import customError "vetback/internal/error"

func (s *appointmentService) Delete(id int) error {
	err := s.storage.DeleteAppointment(id)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}

	return nil
}
