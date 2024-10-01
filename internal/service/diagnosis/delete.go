package diagnosis

import customError "vetback/internal/error"

func (s *diagnosisService) Delete(id int) error {
	err := s.storage.DeleteDiagnosis(id)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}
	return nil
}
