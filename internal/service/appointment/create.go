package appointment

import (
	"context"
	"sync"
	customError "vetback/internal/error"
	"vetback/internal/model"
)

func (s *appointmentService) Create(info model.AppointmentInfo) error {
	// Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1) // Канал ошибок с буфером для первой ошибки
	var wg sync.WaitGroup
	wg.Add(2)

	go s.checkAnimal(ctx, info.AnimalId, errCh, &wg)
	go s.checkDiagnosis(ctx, info.DiagnosisId, errCh, &wg)

	// Ждем завершения всех горутин и закрываем канал ошибок
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Обработка первой ошибки
	for err := range errCh {
		if err != nil {
			cancel() // Отменяем остальные горутины
			return err
		}
	}

	// Создание новой записи о назначении
	err := s.storage.CreateAppointment(info)
	if err != nil {
		s.logger.Error(err)
		return customError.InternalServerError
	}

	return nil
}

func (s *appointmentService) checkAnimal(ctx context.Context, id int, errCh chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done(): // Если контекст отменён, выходим
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
		}
	}
}

func (s *appointmentService) checkDiagnosis(ctx context.Context, id int, errCh chan error, wg *sync.WaitGroup) {
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

func (s *appointmentService) checkOwner(ctx context.Context, id int, errCh chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
		owner, err := s.userService.Get(id)
		if err != nil {
			s.logger.Error(err)
			errCh <- customError.InternalServerError
			return
		}
		if owner == nil {
			errCh <- customError.OwnerNotFound
			return
		}
		if owner.Role != "owner" {
			errCh <- customError.OwnerNotFound
		}
	}
}
