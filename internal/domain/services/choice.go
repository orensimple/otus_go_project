package services

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/models"
)

type ChoiceService struct {
}

func (es *ChoiceService) GetChoice(ctx context.Context, bannerID, slotID int64) (*models.Banner, error) {
	banner := &models.Banner{
		ID: 1,
	}

	return banner, nil
}
