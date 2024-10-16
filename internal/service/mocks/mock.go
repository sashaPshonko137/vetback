// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	time "time"
	model "vetback/internal/model"
	service "vetback/internal/service"

	gomock "github.com/golang/mock/gomock"
)

// MockAnimalService is a mock of AnimalService interface.
type MockAnimalService struct {
	ctrl     *gomock.Controller
	recorder *MockAnimalServiceMockRecorder
}

// MockAnimalServiceMockRecorder is the mock recorder for MockAnimalService.
type MockAnimalServiceMockRecorder struct {
	mock *MockAnimalService
}

// NewMockAnimalService creates a new mock instance.
func NewMockAnimalService(ctrl *gomock.Controller) *MockAnimalService {
	mock := &MockAnimalService{ctrl: ctrl}
	mock.recorder = &MockAnimalServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnimalService) EXPECT() *MockAnimalServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAnimalService) Create(info model.AnimalInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAnimalServiceMockRecorder) Create(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAnimalService)(nil).Create), info)
}

// Delete mocks base method.
func (m *MockAnimalService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAnimalServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAnimalService)(nil).Delete), id)
}

// GeneratePDF mocks base method.
func (m *MockAnimalService) GeneratePDF(animalId int, appointmentService service.AppointmentService, treatmentService service.TreatmentService) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePDF", animalId, appointmentService, treatmentService)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePDF indicates an expected call of GeneratePDF.
func (mr *MockAnimalServiceMockRecorder) GeneratePDF(animalId, appointmentService, treatmentService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePDF", reflect.TypeOf((*MockAnimalService)(nil).GeneratePDF), animalId, appointmentService, treatmentService)
}

// Get mocks base method.
func (m *MockAnimalService) Get(id int) (*model.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAnimalServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAnimalService)(nil).Get), id)
}

// GetMany mocks base method.
func (m *MockAnimalService) GetMany() ([]model.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany")
	ret0, _ := ret[0].([]model.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockAnimalServiceMockRecorder) GetMany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockAnimalService)(nil).GetMany))
}

// Update mocks base method.
func (m *MockAnimalService) Update(id int, info model.AnimalToUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAnimalServiceMockRecorder) Update(id, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAnimalService)(nil).Update), id, info)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserService) Create(info model.UserToStorage) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", info)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserServiceMockRecorder) Create(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserService)(nil).Create), info)
}

// Delete mocks base method.
func (m *MockUserService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserService)(nil).Delete), id)
}

// GenerateToken mocks base method.
func (m *MockUserService) GenerateToken(info model.Claims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", info)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockUserServiceMockRecorder) GenerateToken(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUserService)(nil).GenerateToken), info)
}

// Get mocks base method.
func (m *MockUserService) Get(id int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserService)(nil).Get), id)
}

// GetMany mocks base method.
func (m *MockUserService) GetMany() ([]model.SafeUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany")
	ret0, _ := ret[0].([]model.SafeUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockUserServiceMockRecorder) GetMany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockUserService)(nil).GetMany))
}

// GetSessionInfo mocks base method.
func (m *MockUserService) GetSessionInfo(token string) (*model.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionInfo", token)
	ret0, _ := ret[0].(*model.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionInfo indicates an expected call of GetSessionInfo.
func (mr *MockUserServiceMockRecorder) GetSessionInfo(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionInfo", reflect.TypeOf((*MockUserService)(nil).GetSessionInfo), token)
}

// GetUserByPhone mocks base method.
func (m *MockUserService) GetUserByPhone(phoneNumber string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhone", phoneNumber)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhone indicates an expected call of GetUserByPhone.
func (mr *MockUserServiceMockRecorder) GetUserByPhone(phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhone", reflect.TypeOf((*MockUserService)(nil).GetUserByPhone), phoneNumber)
}

// SignIn mocks base method.
func (m *MockUserService) SignIn(info model.UserToLogin) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", info)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserServiceMockRecorder) SignIn(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUserService)(nil).SignIn), info)
}

// SignUp mocks base method.
func (m *MockUserService) SignUp(info model.UserInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", info)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserServiceMockRecorder) SignUp(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUserService)(nil).SignUp), info)
}

// Update mocks base method.
func (m *MockUserService) Update(id int, info model.UserToUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserServiceMockRecorder) Update(id, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserService)(nil).Update), id, info)
}

// MockDiagnosisService is a mock of DiagnosisService interface.
type MockDiagnosisService struct {
	ctrl     *gomock.Controller
	recorder *MockDiagnosisServiceMockRecorder
}

// MockDiagnosisServiceMockRecorder is the mock recorder for MockDiagnosisService.
type MockDiagnosisServiceMockRecorder struct {
	mock *MockDiagnosisService
}

// NewMockDiagnosisService creates a new mock instance.
func NewMockDiagnosisService(ctrl *gomock.Controller) *MockDiagnosisService {
	mock := &MockDiagnosisService{ctrl: ctrl}
	mock.recorder = &MockDiagnosisServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiagnosisService) EXPECT() *MockDiagnosisServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDiagnosisService) Create(info model.DiagnosisInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDiagnosisServiceMockRecorder) Create(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDiagnosisService)(nil).Create), info)
}

// Delete mocks base method.
func (m *MockDiagnosisService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDiagnosisServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDiagnosisService)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockDiagnosisService) Get(id int) (*model.Diagnosis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Diagnosis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDiagnosisServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDiagnosisService)(nil).Get), id)
}

// GetMany mocks base method.
func (m *MockDiagnosisService) GetMany() ([]model.Diagnosis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany")
	ret0, _ := ret[0].([]model.Diagnosis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockDiagnosisServiceMockRecorder) GetMany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockDiagnosisService)(nil).GetMany))
}

// Update mocks base method.
func (m *MockDiagnosisService) Update(id int, info model.DiagnosisToUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDiagnosisServiceMockRecorder) Update(id, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDiagnosisService)(nil).Update), id, info)
}

// MockTreatmentService is a mock of TreatmentService interface.
type MockTreatmentService struct {
	ctrl     *gomock.Controller
	recorder *MockTreatmentServiceMockRecorder
}

// MockTreatmentServiceMockRecorder is the mock recorder for MockTreatmentService.
type MockTreatmentServiceMockRecorder struct {
	mock *MockTreatmentService
}

// NewMockTreatmentService creates a new mock instance.
func NewMockTreatmentService(ctrl *gomock.Controller) *MockTreatmentService {
	mock := &MockTreatmentService{ctrl: ctrl}
	mock.recorder = &MockTreatmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTreatmentService) EXPECT() *MockTreatmentServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTreatmentService) Create(info model.TreatmentInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTreatmentServiceMockRecorder) Create(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTreatmentService)(nil).Create), info)
}

// Delete mocks base method.
func (m *MockTreatmentService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTreatmentServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTreatmentService)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockTreatmentService) Get(id int) (*model.Treatment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Treatment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTreatmentServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTreatmentService)(nil).Get), id)
}

// GetMany mocks base method.
func (m *MockTreatmentService) GetMany() ([]model.Treatment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany")
	ret0, _ := ret[0].([]model.Treatment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockTreatmentServiceMockRecorder) GetMany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockTreatmentService)(nil).GetMany))
}

// GetManyByAnimal mocks base method.
func (m *MockTreatmentService) GetManyByAnimal(id int) ([]model.Treatment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManyByAnimal", id)
	ret0, _ := ret[0].([]model.Treatment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManyByAnimal indicates an expected call of GetManyByAnimal.
func (mr *MockTreatmentServiceMockRecorder) GetManyByAnimal(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManyByAnimal", reflect.TypeOf((*MockTreatmentService)(nil).GetManyByAnimal), id)
}

// Update mocks base method.
func (m *MockTreatmentService) Update(id int, info model.TreatmentToUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTreatmentServiceMockRecorder) Update(id, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTreatmentService)(nil).Update), id, info)
}

// MockAppointmentService is a mock of AppointmentService interface.
type MockAppointmentService struct {
	ctrl     *gomock.Controller
	recorder *MockAppointmentServiceMockRecorder
}

// MockAppointmentServiceMockRecorder is the mock recorder for MockAppointmentService.
type MockAppointmentServiceMockRecorder struct {
	mock *MockAppointmentService
}

// NewMockAppointmentService creates a new mock instance.
func NewMockAppointmentService(ctrl *gomock.Controller) *MockAppointmentService {
	mock := &MockAppointmentService{ctrl: ctrl}
	mock.recorder = &MockAppointmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppointmentService) EXPECT() *MockAppointmentServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAppointmentService) Create(info model.AppointmentInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAppointmentServiceMockRecorder) Create(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAppointmentService)(nil).Create), info)
}

// Delete mocks base method.
func (m *MockAppointmentService) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAppointmentServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppointmentService)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockAppointmentService) Get(id int) (*model.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*model.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAppointmentServiceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAppointmentService)(nil).Get), id)
}

// GetByAnimal mocks base method.
func (m *MockAppointmentService) GetByAnimal(id int) ([]model.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAnimal", id)
	ret0, _ := ret[0].([]model.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAnimal indicates an expected call of GetByAnimal.
func (mr *MockAppointmentServiceMockRecorder) GetByAnimal(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAnimal", reflect.TypeOf((*MockAppointmentService)(nil).GetByAnimal), id)
}

// GetMany mocks base method.
func (m *MockAppointmentService) GetMany() ([]model.Appointment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany")
	ret0, _ := ret[0].([]model.Appointment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockAppointmentServiceMockRecorder) GetMany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockAppointmentService)(nil).GetMany))
}

// Update mocks base method.
func (m *MockAppointmentService) Update(id int, info model.AppointmentToUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAppointmentServiceMockRecorder) Update(id, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAppointmentService)(nil).Update), id, info)
}

// MockMetricsService is a mock of MetricsService interface.
type MockMetricsService struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsServiceMockRecorder
}

// MockMetricsServiceMockRecorder is the mock recorder for MockMetricsService.
type MockMetricsServiceMockRecorder struct {
	mock *MockMetricsService
}

// NewMockMetricsService creates a new mock instance.
func NewMockMetricsService(ctrl *gomock.Controller) *MockMetricsService {
	mock := &MockMetricsService{ctrl: ctrl}
	mock.recorder = &MockMetricsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricsService) EXPECT() *MockMetricsServiceMockRecorder {
	return m.recorder
}

// Listen mocks base method.
func (m *MockMetricsService) Listen(address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listen", address)
	ret0, _ := ret[0].(error)
	return ret0
}

// Listen indicates an expected call of Listen.
func (mr *MockMetricsServiceMockRecorder) Listen(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockMetricsService)(nil).Listen), address)
}

// ObserveRequest mocks base method.
func (m *MockMetricsService) ObserveRequest(d time.Duration, status int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ObserveRequest", d, status)
}

// ObserveRequest indicates an expected call of ObserveRequest.
func (mr *MockMetricsServiceMockRecorder) ObserveRequest(d, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveRequest", reflect.TypeOf((*MockMetricsService)(nil).ObserveRequest), d, status)
}
