package model

type Treatment struct {
	TreatmentId int
	DoctorId    int
	AnimalId    int
	DiagnosisId int
	Start       string
	Finish      string
	Price       float64
}

type TreatmentInfo struct {
	DoctorId    int
	AnimalId    int
	DiagnosisId int
	Finish      string
	Price       float64
}

type TreatmentToUpdate struct {
	AnimalId    int
	DiagnosisId int
	Finish      string
	Price       float64
}

type SwaggerTreatment struct {
	AnimalId    int
	DiagnosisId int
	Finish      string
	Price       float64
}
