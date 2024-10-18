package api

import (
	"net/http"
	"vetback/internal/model"
	"vetback/internal/service"
)

//go:generate ../../bin/mockgen -source=api.go -destination=mocks/mock.go

type AnimalApi interface {
	NewCreate() http.HandlerFunc
	NewGet() http.HandlerFunc
	NewGetMany() http.HandlerFunc
	NewUpdate() http.HandlerFunc
	NewDelete() http.HandlerFunc
	SendPDF(appointmentService service.AppointmentService, treatmentService service.TreatmentService) http.HandlerFunc
}

type UserApi interface {
	NewSignUp() http.HandlerFunc
	NewSignIn() http.HandlerFunc
	NewSignOut() http.HandlerFunc
	NewGetSessionInfo() http.HandlerFunc
	GetSessionInfo(*http.Request) (*model.Claims, error)
	NewGet() http.HandlerFunc
	NewGetMany() http.HandlerFunc
	NewUpdate() http.HandlerFunc
	NewDelete() http.HandlerFunc
}

type DiagnosisApi interface {
	NewCreate() http.HandlerFunc
	NewGet() http.HandlerFunc
	NewGetMany() http.HandlerFunc
	NewUpdate() http.HandlerFunc
	NewDelete() http.HandlerFunc
}

type TreatmentApi interface {
	NewCreate() http.HandlerFunc
	NewGet() http.HandlerFunc
	NewGetMany() http.HandlerFunc
	NewUpdate() http.HandlerFunc
	NewDelete() http.HandlerFunc
}

type AppointmentApi interface {
	NewCreate() http.HandlerFunc
	NewGet() http.HandlerFunc
	NewGetMany() http.HandlerFunc
	NewUpdate() http.HandlerFunc
	NewDelete() http.HandlerFunc
}
