package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"gitlab.com/faemproject/backend/faem/services/bootstrap/models"
)

func (p *Pg) GetCurrentVersion(ctx context.Context) (*models.DrvAppVersion, error) {
	// only wait for the reply for 30 seconds
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var ver models.DrvAppVersion
	err := p.Db.ModelContext(ctx, &ver).
		Order("created_at DESC").
		Limit(1).
		Select()
	if err != nil {
		return nil, errors.Wrap(err, "failed to select the latest version from the db")
	}
	return &ver, nil
}
