package treatment

import customError "vetback/internal/error"

func (s *treatmentService) Delete(id int) error {
	err := s.storage.DeleteTreatment(id)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}

	return nil
}
