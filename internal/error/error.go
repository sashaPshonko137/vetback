package error

import (
	"errors"
)

var (
	UserAlreadyExist    = errors.New("номер телефона уже зарегистрирован")
	DoctorNotFound      = errors.New("врач не найден")
	OwnerNotFound       = errors.New("владелец не найден")
	FailedLogin         = errors.New("проверьте правильность номера телефона и пароля")
	AnimalNotFound      = errors.New("животное не найдено")
	DiagnosisNotFound   = errors.New("диагноз не найден")
	AppointmentNotFound = errors.New("запись не найдена")
	InternalServerError = errors.New("ошибка сервера")
	NotOwner            = errors.New("пользователь не является владельцем животного")
	NoSpecialization    = errors.New("врач должен иметь специализацию")
	InvalidToken        = errors.New("неверный токен")
	TokenNotFound       = errors.New("токен не найден")
	MustBeOwner         = errors.New("вы должны быть владельцем животного")
	MustBeDoctor        = errors.New("вы должны быть врачом")
	AccountNotFound     = errors.New("аккаунт не найден")
)
