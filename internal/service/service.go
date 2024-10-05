package service

import (
	"time"
	"vetback/internal/model"
)

type AnimalService interface {
	Create(info model.AnimalInfo) error
	Get(id int) (*model.Animal, error)
	GetMany() ([]model.Animal, error)
	Update(id int, info model.AnimalToUpdate) error
	Delete(id int) error
	GeneratePDF(animalId int, appointmentService AppointmentService, treatmentService TreatmentService) (string, error)
}

type UserService interface {
	GenerateToken(info model.Claims) (string, error)
	SignUp(info model.UserInfo) (string, error)
	SignIn(info model.UserToLogin) (string, error)
	GetSessionInfo(token string) (*model.Claims, error)
	Create(info model.UserToStorage) (int, error)
	Get(id int) (*model.User, error)
	GetUserByPhone(phoneNumber string) (*model.User, error)
	GetMany() ([]model.SafeUser, error)
	Update(id int, info model.UserToUpdate) error
	Delete(id int) error
}

type DiagnosisService interface {
	Create(info model.DiagnosisInfo) error
	Get(id int) (*model.Diagnosis, error)
	GetMany() ([]model.Diagnosis, error)
	Update(id int, info model.DiagnosisToUpdate) error
	Delete(id int) error
}

type TreatmentService interface {
	Create(info model.TreatmentInfo) error
	Get(id int) (*model.Treatment, error)
	GetManyByAnimal(id int) ([]model.Treatment, error)
	GetMany() ([]model.Treatment, error)
	Update(id int, info model.TreatmentToUpdate) error
	Delete(id int) error
}

type AppointmentService interface {
	Create(info model.AppointmentInfo) error
	Get(id int) (*model.Appointment, error)
	GetByAnimal(id int) ([]model.Appointment, error)
	GetMany() ([]model.Appointment, error)
	Update(id int, info model.AppointmentToUpdate) error
	Delete(id int) error
}

type MetricsService interface {
	Listen(address string) error
	ObserveRequest(d time.Duration, status int)
}
