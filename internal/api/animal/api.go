package animal

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/api"
	"vetback/internal/service"
)

type animalApi struct {
	service service.AnimalService
	logger  logrus.FieldLogger
	userApi api.UserApi
}

func NewAnimalApi(service service.AnimalService, logger logrus.FieldLogger, userApi api.UserApi) api.AnimalApi {
	return &animalApi{
		service: service,
		logger:  logger,
		userApi: userApi,
	}
}
