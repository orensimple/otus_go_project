package interfaces

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/models"
)

// BannerStorage interface
type BannerStorage interface {
	SetBanner(ctx context.Context, banner *models.Banner) error
	UpdateBanner(ctx context.Context, banner *models.Banner) (*models.Banner, error)
	GetBanners(ctx context.Context) ([]*models.Banner, error)
	DeleteBanner(ctx context.Context, ID int64) error
}
