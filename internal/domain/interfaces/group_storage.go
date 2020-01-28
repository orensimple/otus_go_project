package interfaces

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/models"
)

// GroupStorage interface
type GroupStorage interface {
	SetGroup(ctx context.Context, group *models.Group) error
	UpdateGroup(ctx context.Context, group *models.Group) (*models.Group, error)
	GetGroups(ctx context.Context) ([]*models.Group, error)
	DeleteGroup(ctx context.Context, ID int64) error
}
