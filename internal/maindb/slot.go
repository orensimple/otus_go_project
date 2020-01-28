package maindb

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

// implements domain.interfaces.SlotStorage
type PgSlotStorage struct {
	db *sqlx.DB
}

func NewPgSlotStorage(dsn string) (*PgSlotStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgSlotStorage{db: db}, nil
}

func (pges *PgSlotStorage) SetSlot(ctx context.Context, slot *models.Slot) error {
	query := `
		INSERT INTO slots(id)
		VALUES (:id)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": slot.ID,
	})
	return err
}

func (pges *PgSlotStorage) UpdateSlot(ctx context.Context, slot *models.Slot) (*models.Slot, error) {
	query := `
		UPDATE slots
		SET id = :id
		WHERE id = :id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": slot.ID,
	})
	return slot, err
}

func (pges *PgSlotStorage) GetSlots(ctx context.Context) ([]*models.Slot, error) {
	query := `
	SELECT id FROM slots
`
	rows, err := pges.db.QueryContext(ctx, query)
	if err != nil {
		logger.ContextLogger.Infof("QueryContext", "err", err.Error())
	}
	defer rows.Close()
	var slots []*models.Slot
	for rows.Next() {
		var slot models.Slot
		if err := rows.Scan(&slot.ID); err != nil {
			logger.ContextLogger.Infof("rowScan", "err", err.Error())
		}
		slots = append(slots, &slot)
	}
	if err := rows.Err(); err != nil {
		logger.ContextLogger.Infof("rowErr", "err", err.Error())
	}

	return slots, nil
}

func (pges *PgSlotStorage) DeleteSlot(ctx context.Context, ID int64) error {
	query := `
		DELETE FROM slots
		WHERE id = :id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": ID,
	})
	return err
}
