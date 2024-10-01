package treatment

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/repo"
	"vetback/internal/service"
)

type treatmentService struct {
	storage          repo.Repo
	logger           logrus.FieldLogger
	userService      service.UserService
	animalService    service.AnimalService
	diagnosisService service.DiagnosisService
}

func NewTreatmentService(storage repo.Repo, logger logrus.FieldLogger, userService service.UserService, animalService service.AnimalService, diagnosisService service.DiagnosisService) service.TreatmentService {
	return &treatmentService{
		storage:          storage,
		logger:           logger,
		userService:      userService,
		animalService:    animalService,
		diagnosisService: diagnosisService,
	}
}
