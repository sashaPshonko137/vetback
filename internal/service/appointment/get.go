package appointment

import "vetback/internal/model"

func (s *appointmentService) Get(id int) (*model.Appointment, error) {
	res, err := s.storage.GetAppointment(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
