package user

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/api"
	"vetback/internal/service"
)

type userApi struct {
	service service.UserService
	logger  logrus.FieldLogger
}

func NewUserApi(service service.UserService, logger logrus.FieldLogger) api.UserApi {
	return &userApi{service: service, logger: logger}
}
