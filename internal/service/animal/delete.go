package animal

import customError "vetback/internal/error"

func (s *animalService) Delete(id int) error {
	err := s.storage.DeleteAnimal(id)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}
	return nil
}
