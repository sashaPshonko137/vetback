package repo

import "vetback/internal/model"

type Repo interface {
	CreateAnimal(req model.AnimalInfo) error
	GetAnimal(id int) (*model.Animal, error)
	GetManyAnimals() ([]model.Animal, error)
	UpdateAnimal(id int, req model.AnimalToUpdate) error
	DeleteAnimal(id int) error
	CreateUser(req model.UserToStorage) (int, error)
	GetUser(id int) (*model.User, error)
	GetUserByPhone(phoneNumber string) (*model.User, error)
	GetManyUsers() ([]model.SafeUser, error)
	UpdateUser(id int, req model.UpdatedUserToStorage) error
	DeleteUser(id int) error
	CreateDiagnosis(req model.DiagnosisInfo) error
	GetDiagnosis(id int) (*model.Diagnosis, error)
	GetManyDiagnosis() ([]model.Diagnosis, error)
	UpdateDiagnosis(id int, req model.DiagnosisToUpdate) error
	DeleteDiagnosis(id int) error
	CreateTreatment(req model.TreatmentInfo) error
	GetTreatment(id int) (*model.Treatment, error)
	GetManyTreatmentsByAnimal(id int) ([]model.Treatment, error)
	GetManyTreatments() ([]model.Treatment, error)
	UpdateTreatment(id int, req model.TreatmentToUpdate) error
	DeleteTreatment(id int) error
	CreateAppointment(req model.AppointmentInfo) error
	GetAppointment(id int) (*model.Appointment, error)
	GetAppointmentsByAnimal(int) ([]model.Appointment, error)
	GetManyAppointments() ([]model.Appointment, error)
	UpdateAppointment(id int, req model.AppointmentToUpdate) error
	DeleteAppointment(id int) error
}
