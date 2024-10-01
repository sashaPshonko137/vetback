package animal

import (
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *animalService) GetMany() ([]model.Animal, error) {
	res, err := s.storage.GetManyAnimals()
	if err != nil {
		s.logger.Error(err)
		return nil, customError.InternalServerError
	}
	return res, nil
}
