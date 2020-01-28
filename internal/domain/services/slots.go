package services

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/interfaces"
	"github.com/orensimple/otus_go_project/internal/domain/models"
)

type SlotService struct {
	SlotStorage interfaces.SlotStorage
}

func (es *SlotService) SetSlot(ctx context.Context, ID int64) (*models.Slot, error) {
	slot := &models.Slot{
		ID: ID,
	}

	err := es.SlotStorage.SetSlot(ctx, slot)
	if err != nil {
		return nil, err
	}
	return slot, nil
}

func (es *SlotService) UpdateSlot(ctx context.Context, ID int64) (*models.Slot, error) {
	slot := &models.Slot{
		ID: ID,
	}

	slot, err := es.SlotStorage.UpdateSlot(ctx, slot)
	if err != nil {
		return nil, err
	}
	return slot, nil
}

func (es *SlotService) GetSlots(ctx context.Context) ([]*models.Slot, error) {
	slots, err := es.SlotStorage.GetSlots(ctx)
	if err != nil {
		return nil, err
	}
	return slots, nil
}

func (es *SlotService) DeleteSlot(ctx context.Context, ID int64) error {
	err := es.SlotStorage.DeleteSlot(ctx, ID)

	return err
}
