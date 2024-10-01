package treatment

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/api"
	"vetback/internal/service"
)

type treatmentApi struct {
	service service.TreatmentService
	logger  logrus.FieldLogger
	userApi api.UserApi
}

func NewTreatmentApi(service service.TreatmentService, logger logrus.FieldLogger, userApi api.UserApi) api.TreatmentApi {
	return &treatmentApi{service: service, logger: logger, userApi: userApi}
}
