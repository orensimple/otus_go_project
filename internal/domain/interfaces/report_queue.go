package interfaces

import (
	"context"

	"github.com/orensimple/otus_go_project/internal/domain/models"
)

// ReportQueue interface
type ReportQueue interface {
	PublicReport(ctx context.Context, report models.ReportQueueFormat) error
}
