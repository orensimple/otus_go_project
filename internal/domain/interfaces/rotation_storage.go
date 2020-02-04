package interfaces

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/models"
)

// RotationStorage interface
type RotationStorage interface {
	SetRotation(ctx context.Context, rotation *models.Rotation) error
	UpdateRotation(ctx context.Context, rotation *models.Rotation) (*models.Rotation, error)
	GetRotations(ctx context.Context) ([]*models.Rotation, error)
	DeleteRotation(ctx context.Context, bannerID int64, slotID int64) error
}

type MemRotationStorage interface {
	SetRotation(ctx context.Context, slotID int64, bannerID int64) error
	GetRotations(ctx context.Context, slotID int64, groupID int64) ([]int64, error)
	DeleteRotation(ctx context.Context, slotID int64, groupID int64) (int64, error)
}
