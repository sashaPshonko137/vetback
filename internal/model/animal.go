package model

type Animal struct {
	AnimalId    int
	DoctorId    int
	OwnerId     int
	DiagnosisId int
	Name        string
	Birthdate   string
	Breed       string
	Sex         string
}

type AnimalInfo struct {
	DoctorId    int
	OwnerId     int
	DiagnosisId int
	Name        string
	Birthdate   string
	Breed       string
	Sex         string
}

type AnimalToUpdate struct {
	DoctorId    int
	DiagnosisId int
	Name        string
	Birthdate   string
	Breed       string
	Sex         string
}

type SwaggerAnimal struct {
	DoctorId    int
	DiagnosisId int
	Name        string
	Birthdate   string
	Breed       string
	Sex         string
}
