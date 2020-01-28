package services

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/interfaces"
	"github.com/orensimple/otus_go_project/internal/domain/models"
)

type RotationService struct {
	RotationStorage interfaces.RotationStorage
}

func (es *RotationService) SetRotation(ctx context.Context, bannerID, slotID int64, title string) (*models.Rotation, error) {
	rotation := &models.Rotation{
		BannerID: bannerID,
		SlotID:   slotID,
		Title:    title,
	}

	err := es.RotationStorage.SetRotation(ctx, rotation)
	if err != nil {
		return nil, err
	}
	return rotation, nil
}

func (es *RotationService) UpdateRotation(ctx context.Context, bannerID, slotID int64, title string) (*models.Rotation, error) {
	rotation := &models.Rotation{
		BannerID: bannerID,
		SlotID:   slotID,
		Title:    title,
	}

	rotation, err := es.RotationStorage.UpdateRotation(ctx, rotation)
	if err != nil {
		return nil, err
	}
	return rotation, nil
}

func (es *RotationService) GetRotations(ctx context.Context) ([]*models.Rotation, error) {
	rotations, err := es.RotationStorage.GetRotations(ctx)
	if err != nil {
		return nil, err
	}
	return rotations, nil
}

func (es *RotationService) DeleteRotation(ctx context.Context, bannerID int64, slotID int64) error {
	err := es.RotationStorage.DeleteRotation(ctx, bannerID, slotID)

	return err
}
