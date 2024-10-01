package model

type Diagnosis struct {
	DiagnosisId int
	DoctorId    int
	Name        string
	Date        string
}

type DiagnosisInfo struct {
	DoctorId int
	Name     string
}

type DiagnosisToUpdate struct {
	Name string
}

type SwaggerDiagnosis struct {
	Name string
}
