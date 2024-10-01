package appointment

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/api"
	"vetback/internal/service"
)

type appointmentApi struct {
	service service.AppointmentService
	logger  logrus.FieldLogger
	userApi api.UserApi
}

func NewAppointmentApi(service service.AppointmentService, logger logrus.FieldLogger, userApi api.UserApi) api.AppointmentApi {
	return &appointmentApi{
		service: service,
		logger:  logger,
		userApi: userApi,
	}
}
