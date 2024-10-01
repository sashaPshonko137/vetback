package model

import "time"

type Response struct {
	Message string    `json:"message"`
	Data    time.Time `json:"data"`
}

func NewResponse(message string) Response {
	return Response{Message: message, Data: time.Now()}
}
