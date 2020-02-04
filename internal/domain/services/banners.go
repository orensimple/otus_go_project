package services

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/interfaces"
	"github.com/orensimple/otus_go_project/internal/domain/models"
)

type BannerService struct {
	BannerStorage      interfaces.BannerStorage
	MemRotationStorage interfaces.MemRotationStorage
}

func (es *BannerService) SetBanner(ctx context.Context, ID int64, title string) (*models.Banner, error) {
	banner := &models.Banner{
		ID:    ID,
		Title: title,
	}

	err := es.BannerStorage.SetBanner(ctx, banner)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func (es *BannerService) UpdateBanner(ctx context.Context, ID int64, title string) (*models.Banner, error) {
	banner := &models.Banner{
		ID:    ID,
		Title: title,
	}

	banner, err := es.BannerStorage.UpdateBanner(ctx, banner)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func (es *BannerService) GetBanners(ctx context.Context) ([]*models.Banner, error) {
	banners, err := es.BannerStorage.GetBanners(ctx)
	if err != nil {
		return nil, err
	}
	return banners, nil
}

func (es *BannerService) DeleteBanner(ctx context.Context, ID int64) error {
	err := es.BannerStorage.DeleteBanner(ctx, ID)

	return err
}
