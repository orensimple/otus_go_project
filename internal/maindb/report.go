package maindb

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

type PgReportStorage struct {
	db *sqlx.DB
}

func NewPgReportStorage(dsn string) (*PgReportStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgReportStorage{db: db}, nil
}

func (pges *PgReportStorage) SetReport(ctx context.Context, report *models.Report) error {
	query := `
		INSERT INTO reports(banner_id, group_id, slot_id, show, conversion)
		VALUES (:banner_id, :group_id, slot_id, :show, :conversion)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"banner_id":  report.BannerID,
		"group_id":   report.GroupID,
		"slot_id":    report.SlotID,
		"show":       report.Show,
		"conversion": report.Conversion,
	})
	return err
}

func (pges *PgReportStorage) UpdateReport(ctx context.Context, report *models.Report) (*models.Report, error) {
	query := `
		UPDATE reports
		SET show = :show, conversion = :conversion
		WHERE banner_id = :banner_ID AND group_id = : group_id AND slot_id = : slot_id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"banner_id":  report.BannerID,
		"group_id":   report.GroupID,
		"slot_id":    report.SlotID,
		"show":       report.Show,
		"conversion": report.Conversion,
	})
	return report, err
}

func (pges *PgReportStorage) GetReports(ctx context.Context) ([]*models.Report, error) {
	query := `
		SELECT banner_id, group_id, slot_id, show, conversion FROM reports
	`

	rows, err := pges.db.QueryContext(ctx, query)
	if err != nil {
		logger.ContextLogger.Infof("QueryContext", "err", err.Error())
	}
	defer rows.Close()
	var reports []*models.Report
	for rows.Next() {
		var report models.Report
		if err := rows.Scan(&report.BannerID); err != nil {
			logger.ContextLogger.Infof("rowScan", "err", err.Error())
		}
		reports = append(reports, &report)
	}
	if err := rows.Err(); err != nil {
		logger.ContextLogger.Infof("rowErr", "err", err.Error())
	}

	return reports, nil
}

func (pges *PgReportStorage) DeleteReport(ctx context.Context, BannerID int64, GroupID int64, SlotID int64) error {
	query := `
		DELETE FROM reports
		WHERE banner_id = :banner_id AND group_id = : group_id AND slot_id = : slot_id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"banner_id": BannerID,
		"group_id":  GroupID,
		"slot_id":   SlotID,
	})
	return err
}
