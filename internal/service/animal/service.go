package animal

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/repo"
	"vetback/internal/service"
)

type animalService struct {
	storage          repo.Repo
	logger           logrus.FieldLogger
	userService      service.UserService
	diagnosisService service.DiagnosisService
}

func NewAnimalService(storage repo.Repo, logger logrus.FieldLogger, userService service.UserService, diagnosisService service.DiagnosisService) service.AnimalService {
	return &animalService{storage: storage, logger: logger, userService: userService, diagnosisService: diagnosisService}
}
