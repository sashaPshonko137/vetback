package model

type Appointment struct {
	AppointmentId int
	DoctorId      int
	OwnerId       int
	AnimalId      int
	DiagnosisId   int
	Date          string
	Status        string
	Reason        string
}

type AppointmentInfo struct {
	OwnerId     int
	DoctorId    int
	AnimalId    int
	DiagnosisId int
	Date        string
	Status      string
	Reason      string
}

type AppointmentToUpdate struct {
	DoctorId    int
	OwnerId     int
	AnimalId    int
	DiagnosisId int
	Date        string
	Status      string
	Reason      string
}

type SwaggerAppointment struct {
	DoctorId    int
	OwnerId     int
	AnimalId    int
	DiagnosisId int
	Date        string
	Status      string
	Reason      string
}
