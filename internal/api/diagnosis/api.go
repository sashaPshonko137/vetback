package diagnosis

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/api"
	"vetback/internal/service"
)

type diagnosisApi struct {
	service service.DiagnosisService
	logger  logrus.FieldLogger
	userApi api.UserApi
}

func NewDiagnosisApi(service service.DiagnosisService, logger logrus.FieldLogger, userApi api.UserApi) api.DiagnosisApi {
	return &diagnosisApi{service: service, logger: logger, userApi: userApi}
}
