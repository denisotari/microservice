package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gitlab.com/faemproject/backend/faem/services/bootstrap/models"
)

type mockDB struct {
	mock.Mock
}

func (m *mockDB) GetCurrentVersion(ctx context.Context) (*models.DrvAppVersion, error) {
	args := m.Called(ctx)
	return args.Get(0).(*models.DrvAppVersion), args.Error(1) // also you need to check for nil if the case is possible
}

func TestBootstrap_GetCurrentVersion(t *testing.T) {
	var mdb mockDB
	mdb.On("GetCurrentVersion", mock.Anything).Return(&models.DrvAppVersion{}, nil)

	h := Handler{DB: &mdb}
	_, err := h.GetCurrentVersion(context.Background())
	assert.NoError(t, err)

	// TODO: just a sample, test the real logic here
	mdb.AssertExpectations(t)
}
