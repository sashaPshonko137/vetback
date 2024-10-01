package animal

import "vetback/internal/model"

func (s *animalService) Get(id int) (*model.Animal, error) {
	res, err := s.storage.GetAnimal(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
