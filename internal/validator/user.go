package validator

import (
	"errors"
	"vetback/internal/model"
)

func ValidateUser(model model.UserInfo) error {
	if model.Password == "" {
		return errors.New("необходимо указать пароль")
	}
	if len(model.Password) > 50 {
		return errors.New("длина пароля не должна превышать 50 символов")
	}
	if model.FullName == "" {
		return errors.New("необходимо указать ФИО")
	}
	if len(model.FullName) > 50 {
		return errors.New("длина ФИО не должна превышать 50 символов")
	}
	if model.Address == "" {
		return errors.New("необходимо указать адрес")
	}
	if len(model.Address) > 50 {
		return errors.New("длина адреса не должна превышать 50 символов")
	}
	if model.PhoneNumber == "" {
		return errors.New("необходимо указать телефон")
	}
	if len(model.PhoneNumber) > 50 {
		return errors.New("длина телефона не должна превышать 50 символов")
	}
	if model.Role == "" {
		return errors.New("необходимо указать роль")
	}
	if len(model.Role) > 50 {
		return errors.New("длина роли не должна превышать 50 символов")
	}
	if len(model.Specialization) > 50 {
		return errors.New("длина специализации не должна превышать 50 символов")
	}
	return nil
}

func ValidateUserToUpdate(model model.UserToUpdate) error {
	if len(model.Password) > 50 {
		return errors.New("длина пароля не должна превышать 50 символов")
	}
	if len(model.FullName) > 50 {
		return errors.New("длина ФИО не должна превышать 50 символов")
	}
	if len(model.Address) > 50 {
		return errors.New("длина адреса не должна превышать 50 символов")
	}
	if len(model.PhoneNumber) > 50 {
		return errors.New("длина телефона не должна превышать 50 символов")
	}
	if len(model.Specialization) > 50 {
		return errors.New("длина специализации не должна превышать 50 символов")
	}
	return nil
}
