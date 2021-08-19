package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.com/faemproject/backend/faem/services/bootstrap/handler"
	"gitlab.com/faemproject/backend/faem/services/bootstrap/proto"
)

type mockPublisher struct {
	mock.Mock
}

func (m *mockPublisher) UserCreated(user *proto.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestGreet(t *testing.T) {
	var mpub mockPublisher
	mpub.On("UserCreated", mock.Anything).Return(nil)

	e := echo.New()
	rest := Rest{
		Handler: &handler.Handler{
			Pub: &mpub,
		},
	}

	// Test an empty request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(req, w)

	assert.NoError(t, rest.Hello(c))
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test non-empty request
	q := make(url.Values)
	q.Set("name", "john")
	req = httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
	w = httptest.NewRecorder()
	c = e.NewContext(req, w)

	assert.NoError(t, rest.Hello(c))
	assert.Equal(t, http.StatusOK, w.Code)
}
