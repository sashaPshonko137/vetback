package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"vetback/internal/api/animal"
	mock_api "vetback/internal/api/mocks"
	"vetback/internal/model"
	mock_service "vetback/internal/service/mocks"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

func TestCreate(t *testing.T) {
	type mockBehavior = func(a *mock_api.MockUserApi, s *mock_service.MockAnimalService, animal model.AnimalInfo)

	testTable := []struct {
		name string
		inputAnimal model.AnimalInfo
		inputBody string
		mockBehavior mockBehavior
		expectedStatusCode int
		expectedResponseBody model.Response
	} {
		{
			name: "ok",
			inputAnimal: model.AnimalInfo{DoctorId: 14, OwnerId: 14, DiagnosisId: 14, Name: "bebrina", Birthdate: "2006-01-02", Breed: "umnaya", Sex: "другой"},
			inputBody: `{"doctorId":14,"ownerId":14,"diagnosisId":14,"name":"bebrina","birthdate":"2006-01-02","breed":"umnaya","sex":"другой"}`,
			mockBehavior: func(a *mock_api.MockUserApi, s *mock_service.MockAnimalService, animal model.AnimalInfo) {
				a.EXPECT().GetSessionInfo(gomock.AssignableToTypeOf(&http.Request{})).Return(&model.Claims{Role: "owner", UserId: 14}, nil)
				s.EXPECT().Create(animal).Return(nil)
			},
			expectedStatusCode: 200,
		},
		{
			name: "not owner",
			inputBody: `{"doctorId":14,"ownerId":14,"diagnosisId":14,"name":"bebrina","birthdate":"2006-01-02","breed":"umnaya","sex":"другой"}`,
			mockBehavior: func(a *mock_api.MockUserApi, s *mock_service.MockAnimalService, animal model.AnimalInfo) {
				a.EXPECT().GetSessionInfo(gomock.AssignableToTypeOf(&http.Request{})).Return(&model.Claims{Role: "doctor", UserId: 14}, nil)
			},
			expectedStatusCode: 403,
		},
		{
			name: "empty body",
			inputBody: "",
			mockBehavior: func(a *mock_api.MockUserApi, s *mock_service.MockAnimalService, animal model.AnimalInfo) {
				a.EXPECT().GetSessionInfo(gomock.AssignableToTypeOf(&http.Request{})).Return(&model.Claims{Role: "owner", UserId: 14}, nil)
			},
			expectedStatusCode: 400,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			api := mock_api.NewMockUserApi(c)
			serv := mock_service.NewMockAnimalService(c)
			logger := logrus.New()

			testCase.mockBehavior(api, serv, testCase.inputAnimal)

			handler := animal.NewAnimalApi(serv, logger, api)

			r := chi.NewRouter()
			r.Post("/animal", handler.NewCreate())

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/animal", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}