package maindb

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

type PgRotationStorage struct {
	db *sqlx.DB
}

func NewPgRotationStorage(dsn string) (*PgRotationStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgRotationStorage{db: db}, nil
}

func (pges *PgRotationStorage) SetRotation(ctx context.Context, rotation *models.Rotation) error {
	query := `
		INSERT INTO rotations(banner_id, slot_id, title)
		VALUES (:banner_id, :slot_id, :title)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"banner_id": rotation.BannerID,
		"slot_id":   rotation.SlotID,
		"title":     rotation.Title,
	})
	return err
}

func (pges *PgRotationStorage) UpdateRotation(ctx context.Context, rotation *models.Rotation) (*models.Rotation, error) {
	query := `
		UPDATE rotations
		SET title = :title
		WHERE banner_id = :banner_id AND slot_id = : slot_id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"banner_id": rotation.BannerID,
		"slot_id":   rotation.SlotID,
		"title":     rotation.Title,
	})
	return rotation, err
}

func (pges *PgRotationStorage) GetRotations(ctx context.Context) ([]*models.Rotation, error) {
	query := `
		SELECT banner_id, slot_id, title FROM rotations
	`

	rows, err := pges.db.QueryContext(ctx, query)
	if err != nil {
		logger.ContextLogger.Infof("QueryContext", "err", err.Error())
	}
	defer rows.Close()
	var rotations []*models.Rotation
	for rows.Next() {
		var rotation models.Rotation
		if err := rows.Scan(&rotation.BannerID, &rotation.SlotID, &rotation.Title); err != nil {
			logger.ContextLogger.Infof("rowScan", "err", err.Error())
		}
		rotations = append(rotations, &rotation)
	}
	if err := rows.Err(); err != nil {
		logger.ContextLogger.Infof("rowErr", "err", err.Error())
	}

	return rotations, nil
}

func (pges *PgRotationStorage) DeleteRotation(ctx context.Context, bannerID int64, slotID int64) error {
	query := `
		DELETE FROM rotations
		WHERE banner_id = :banner_id AND slot_id = :slot_id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"banner_id": bannerID,
		"slot_id":   slotID,
	})

	return err
}
