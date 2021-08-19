package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.com/faemproject/backend/faem/services/bootstrap/proto"
)

type mockPublisher struct {
	mock.Mock
}

func (m *mockPublisher) UserCreated(user *proto.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestBootstrap_Hello(t *testing.T) {
	var (
		in   proto.User
		mpub mockPublisher
	)
	mpub.On("UserCreated", mock.Anything).Return(nil)
	h := Handler{Pub: &mpub}

	_, err := h.Hello(context.Background(), &in)
	assert.Error(t, err)

	in.Name = "john"
	out, err := h.Hello(context.Background(), &in)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, John", out.Message)
}
