package animal

import (
	"context"
	"sync"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *animalService) Create(model model.AnimalInfo) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)

	go s.checkDoctor(ctx, model.DoctorId, errCh, &wg)
	go s.checkDiagnosis(ctx, model.DiagnosisId, errCh, &wg)

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			cancel()
			return err
		}
	}

	err := s.storage.CreateAnimal(model)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}

	return nil
}

func (s *animalService) checkDoctor(ctx context.Context, id int, errCh chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
		doctor, err := s.userService.Get(id)
		if err != nil {
			s.logger.Error(err)
			errCh <- customError.InternalServerError
			return
		}
		if doctor == nil {
			errCh <- customError.DoctorNotFound
			return
		}
		if doctor.Role != "doctor" {
			errCh <- customError.DoctorNotFound
		}
	}
}

func (s *animalService) checkDiagnosis(ctx context.Context, id int, errCh chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
		diagnosis, err := s.diagnosisService.Get(id)
		if err != nil {
			s.logger.Error(err)
			errCh <- customError.InternalServerError
			return
		}
		if diagnosis == nil {
			errCh <- customError.DiagnosisNotFound
		}
	}
}
