package diagnosis

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/repo"
	"vetback/internal/service"
)

type diagnosisService struct {
	storage     repo.Repo
	logger      logrus.FieldLogger
	userService service.UserService
}

func NewDiagnosisService(storage repo.Repo, logger logrus.FieldLogger, userService service.UserService) service.DiagnosisService {
	return &diagnosisService{storage: storage, logger: logger, userService: userService}
}
