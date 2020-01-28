package services

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/interfaces"
	"github.com/orensimple/otus_go_project/internal/domain/models"
)

type GroupService struct {
	GroupStorage interfaces.GroupStorage
}

func (es *GroupService) SetGroup(ctx context.Context, ID int64) (*models.Group, error) {
	group := &models.Group{
		ID: ID,
	}

	err := es.GroupStorage.SetGroup(ctx, group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (es *GroupService) UpdateGroup(ctx context.Context, ID int64) (*models.Group, error) {
	group := &models.Group{
		ID: ID,
	}

	group, err := es.GroupStorage.UpdateGroup(ctx, group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (es *GroupService) GetGroups(ctx context.Context) ([]*models.Group, error) {
	groups, err := es.GroupStorage.GetGroups(ctx)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (es *GroupService) DeleteGroup(ctx context.Context, ID int64) error {
	err := es.GroupStorage.DeleteGroup(ctx, ID)

	return err
}
