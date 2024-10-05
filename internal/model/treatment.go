package model

type Treatment struct {
	Name        string
	TreatmentId int
	DoctorId    int
	AnimalId    int
	DiagnosisId int
	Start       string
	Finish      string
	Price       float64
}

type TreatmentInfo struct {
	Name        string
	DoctorId    int
	AnimalId    int
	DiagnosisId int
	Finish      string
	Price       float64
}

type TreatmentToUpdate struct {
	DoctorId    int
	Name        string
	AnimalId    int
	DiagnosisId int
	Finish      string
	Price       float64
}

type SwaggerTreatment struct {
	DoctorId    int
	Name        string
	AnimalId    int
	DiagnosisId int
	Finish      string
	Price       float64
}
