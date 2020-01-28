package maindb

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/orensimple/otus_go_project/internal/domain/models"
	"github.com/orensimple/otus_go_project/internal/logger"
)

type PgGroupStorage struct {
	db *sqlx.DB
}

func NewPgGroupStorage(dsn string) (*PgGroupStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgGroupStorage{db: db}, nil
}

func (pges *PgGroupStorage) SetGroup(ctx context.Context, group *models.Group) error {
	query := `
		INSERT INTO groups(id)
		VALUES (:id)
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": group.ID,
	})
	return err
}

func (pges *PgGroupStorage) UpdateGroup(ctx context.Context, group *models.Group) (*models.Group, error) {
	query := `
		UPDATE groups
		SET id = :id
		WHERE id = :id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": group.ID,
	})
	return group, err
}

func (pges *PgGroupStorage) GetGroups(ctx context.Context) ([]*models.Group, error) {
	query := `
		SELECT id FROM groups
	`
	rows, err := pges.db.QueryContext(ctx, query)
	if err != nil {
		logger.ContextLogger.Infof("QueryContext", "err", err.Error())
	}
	defer rows.Close()
	var groups []*models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.ID); err != nil {
			logger.ContextLogger.Infof("rowScan", "err", err.Error())
		}
		groups = append(groups, &group)
	}
	if err := rows.Err(); err != nil {
		logger.ContextLogger.Infof("rowErr", "err", err.Error())
	}

	return groups, nil
}

func (pges *PgGroupStorage) DeleteGroup(ctx context.Context, ID int64) error {
	query := `
		DELETE FROM groups
		WHERE id = :id
	`
	_, err := pges.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": ID,
	})
	return err
}
