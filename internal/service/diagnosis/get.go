package diagnosis

import "vetback/internal/model"

func (s *diagnosisService) Get(id int) (*model.Diagnosis, error) {
	res, err := s.storage.GetDiagnosis(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
