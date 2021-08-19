package handler

import (
	"context"

	"github.com/pkg/errors"

	"gitlab.com/faemproject/backend/faem/services/bootstrap/models"
	"gitlab.com/faemproject/backend/faem/services/bootstrap/proto"
)

type VersionerRepository interface {
	GetCurrentVersion(context.Context) (*models.DrvAppVersion, error)
}

func (h *Handler) GetCurrentVersion(ctx context.Context) (*proto.CurrentVersion, error) {
	version, err := h.DB.GetCurrentVersion(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the current version")
	}
	return version.ToProto(), nil
}
