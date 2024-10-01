package treatment

import (
	"context"
	"sync"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *treatmentService) Create(info model.TreatmentInfo) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go s.checkAnimal(ctx, info.AnimalId, errCh, &wg)
	go s.checkDiagnosis(ctx, info.DiagnosisId, errCh, &wg)

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

	err := s.storage.CreateTreatment(info)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}

	return nil
}

func (s *treatmentService) checkAnimal(ctx context.Context, id int, errCh chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
		animal, err := s.animalService.Get(id)
		if err != nil {
			s.logger.Error(err)
			errCh <- customError.InternalServerError
			return
		}
		if animal == nil {
			errCh <- customError.AnimalNotFound
			return
		}
	}
}

func (s *treatmentService) checkDiagnosis(ctx context.Context, id int, errCh chan error, wg *sync.WaitGroup) {
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
