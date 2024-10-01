package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UserId         int
	Hash           string
	FullName       string
	Address        string
	PhoneNumber    string
	Role           string
	Specialization string
}

type UserInfo struct {
	Password       string
	FullName       string
	Address        string
	PhoneNumber    string
	Role           string
	Specialization string
}

type UserToStorage struct {
	Hash           string
	FullName       string
	Address        string
	PhoneNumber    string
	Role           string
	Specialization string
}

type UserToLogin struct {
	PhoneNumber string
	Password    string
}

type Claims struct {
	UserId         int
	FullName       string
	Address        string
	PhoneNumber    string
	Role           string
	Specialization string
	exp            int
	iat            int
	jwt.RegisteredClaims
}

type SafeUser struct {
	UserId         int
	FullName       string
	Role           string
	Specialization string
}

type UserToUpdate struct {
	Password       string
	FullName       string
	Address        string
	PhoneNumber    string
	Specialization string
}

type UpdatedUserToStorage struct {
	Hash           string
	FullName       string
	Address        string
	PhoneNumber    string
	Specialization string
}
