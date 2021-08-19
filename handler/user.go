package handler

import (
	"context"

	"gitlab.com/faemproject/backend/faem/services/bootstrap/proto"
)

func (h *Handler) CreateUser(ctx context.Context, user *proto.User) error {
	// some business logic to create new user, you can use publisher here to notify other services about the event
	return nil
}
