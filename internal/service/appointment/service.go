package appointment

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/repo"
	"vetback/internal/service"
)

type appointmentService struct {
	storage          repo.Repo
	logger           logrus.FieldLogger
	userService      service.UserService
	animalService    service.AnimalService
	diagnosisService service.DiagnosisService
}

func NewAppointmentService(storage repo.Repo, logger logrus.FieldLogger, userService service.UserService, animalService service.AnimalService, diagnosisService service.DiagnosisService) service.AppointmentService {
	return &appointmentService{
		storage:          storage,
		logger:           logger,
		userService:      userService,
		animalService:    animalService,
		diagnosisService: diagnosisService,
	}
}
