package user

import (
	"github.com/sirupsen/logrus"
	"vetback/internal/config"
	"vetback/internal/repo"
	"vetback/internal/service"
)

type userService struct {
	storage repo.Repo
	logger  logrus.FieldLogger
	config  config.Config
}

func NewUserService(storage repo.Repo, logger logrus.FieldLogger, cfg config.Config) service.UserService {
	return &userService{storage: storage, logger: logger, config: cfg}
}
