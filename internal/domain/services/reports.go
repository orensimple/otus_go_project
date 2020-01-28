package services

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/interfaces"
	"github.com/orensimple/otus_go_project/internal/domain/models"
)

type ReportService struct {
	ReportStorage interfaces.ReportStorage
}

func (es *ReportService) SetReport(ctx context.Context, bannerID, groupID, slotID, show, conversion int64) (*models.Report, error) {
	report := &models.Report{
		BannerID:   bannerID,
		GroupID:    groupID,
		SlotID:     slotID,
		Show:       show,
		Conversion: conversion,
	}

	err := es.ReportStorage.SetReport(ctx, report)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (es *ReportService) UpdateReport(ctx context.Context, bannerID, groupID, slotID, show, conversion int64) (*models.Report, error) {
	report := &models.Report{
		BannerID:   bannerID,
		GroupID:    groupID,
		SlotID:     slotID,
		Show:       show,
		Conversion: conversion,
	}

	report, err := es.ReportStorage.UpdateReport(ctx, report)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (es *ReportService) GetReports(ctx context.Context) ([]*models.Report, error) {
	reports, err := es.ReportStorage.GetReports(ctx)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (es *ReportService) DeleteReport(ctx context.Context, BannerID int64, GroupID int64, SlotID int64) error {
	err := es.ReportStorage.DeleteReport(ctx, BannerID, GroupID, SlotID)

	return err
}
