package maindb

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

type PgBannerStorage struct {
	db *sqlx.DB
}

func NewPgBannerStorage(dsn string) (*PgBannerStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgBannerStorage{db: db}, nil
}

func (pges *PgBannerStorage) SetBanner(ctx context.Context, banner *models.Banner) error {
	query := `
		INSERT INTO banners(id, title)
		VALUES (:id, :title)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":    banner.ID,
		"title": banner.Title,
	})
	return err
}

func (pges *PgBannerStorage) GetBanners(ctx context.Context) ([]*models.Banner, error) {
	query := `
		SELECT id, title FROM banners
	`
	rows, err := pges.db.QueryContext(ctx, query)
	if err != nil {
		logger.ContextLogger.Infof("QueryContext", "err", err.Error())
	}
	defer rows.Close()
	var banners []*models.Banner
	for rows.Next() {
		var banner models.Banner
		if err := rows.Scan(&banner.ID, &banner.Title); err != nil {
			logger.ContextLogger.Infof("rowScan", "err", err.Error())
		}
		banners = append(banners, &banner)
	}
	if err := rows.Err(); err != nil {
		logger.ContextLogger.Infof("rowErr", "err", err.Error())
	}

	return banners, nil
}

func (pges *PgBannerStorage) UpdateBanner(ctx context.Context, banner *models.Banner) (*models.Banner, error) {
	query := `
		UPDATE banners
		SET title = :title
		WHERE id = :id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":    banner.ID,
		"title": banner.Title,
	})
	return banner, err
}

func (pges *PgBannerStorage) DeleteBanner(ctx context.Context, ID int64) error {
	query := `
		DELETE FROM banners
		WHERE id = :id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": ID,
	})
	return err
}
