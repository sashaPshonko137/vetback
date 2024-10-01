package treatment

import "vetback/internal/model"

func (s *treatmentService) Get(id int) (*model.Treatment, error) {
	res, err := s.storage.GetTreatment(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
