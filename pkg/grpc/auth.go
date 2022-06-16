package grpc

import (
	"github.com/google/uuid"
	"github.com/petar-arandjic/virtual_kelner.proto/gen/go/business/v1"
)

type auth struct {
	*business.Auth
}

func (a auth) GetId() *uuid.UUID {
	id, err := uuid.Parse(a.Id)
	if err != nil {

	}
	return &id
}

func (a auth) IsRole(role string) bool {
	for _, value := range a.Roles {
		if role == value {
			return true
		}
	}

	return false
}
