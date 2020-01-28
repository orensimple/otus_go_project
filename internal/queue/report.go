package maindb

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/orensimple/otus_go_project/internal/domain/models"
)

// implements domain.interfaces.ReportQueue
type PgReportQueue struct {
	db *sqlx.DB
}

func NewPgReportQueue(dsn string) (*PgReportQueue, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgReportQueue{db: db}, nil
}

func (pges *PgReportQueue) PublicReport(ctx context.Context, report *models.Report) error {
	query := `
		INSERT INTO reports(id, owner, title, text, start_time, end_time)
		VALUES (:id, :owner, :title, :text, :start_time, :end_time)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"bannerID":   report.BannerID,
		"slotID":     report.GroupID,
		"show":       report.Show,
		"conversion": report.Conversion,
	})
	return err
}
