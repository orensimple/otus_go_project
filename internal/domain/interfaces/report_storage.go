package interfaces

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/models"
)

// ReportStorage interface
type ReportStorage interface {
	SetReport(ctx context.Context, report *models.Report) error
	UpdateReport(ctx context.Context, report *models.Report) (*models.Report, error)
	GetReports(ctx context.Context) ([]*models.Report, error)
	DeleteReport(ctx context.Context, BannerID int64, GroupID int64, SlotID int64) error
}
