package interfaces

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/models"
)

// SlotStorage interface
type SlotStorage interface {
	SetSlot(ctx context.Context, slot *models.Slot) error
	UpdateSlot(ctx context.Context, slot *models.Slot) (*models.Slot, error)
	GetSlots(ctx context.Context) ([]*models.Slot, error)
	DeleteSlot(ctx context.Context, ID int64) error
}
