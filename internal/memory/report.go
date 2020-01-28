package memory

import (
	"context"
	"sync"

	"github.com/orensimple//otus_go_project/internal/domain/errors"
	"github.com/orensimple//otus_go_project/internal/domain/models"
)

type MemReportStorage struct {
	reports map[int64]*models.Report
	mutex   *sync.Mutex
}

func NewMemReportStorage() *MemReportStorage {
	return &MemReportStorage{
		reports: make(map[int64]*models.Report),
		mutex:   new(sync.Mutex),
	}
}

func (mem *MemReportStorage) SaveReport(ctx context.Context, report *models.Report) error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()
	if _, ok := mem.reports[report.BannerID]; ok {
		return errors.ErrReportExist

	}
	mem.reports[report.BannerID] = report
	return nil
}

func (mem *MemReportStorage) UpdateReport(ctx context.Context, report *models.Report) (*models.Report, error) {
	if _, ok := mem.reports[report.BannerID]; ok {
		mem.mutex.Lock()
		mem.reports[report.BannerID] = report
		mem.mutex.Unlock()

		return report, nil
	}

	return nil, errors.ErrReportNotFound
}

func (mem *MemReportStorage) GetReports(ctx context.Context) ([]*models.Report, error) {
	Reports := make([]*models.Report, 0)
	mem.mutex.Lock()
	for _, bm := range mem.reports {
		Reports = append(Reports, bm)
	}
	mem.mutex.Unlock()

	return Reports, nil
}

func (mem *MemReportStorage) DeleteReport(ctx context.Context, BannerID int64, GroupID int64, SlotID int64) error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()

	_, ex := mem.reports[BannerID]
	if ex {
		delete(mem.reports, BannerID)
		return nil
	}

	return errors.ErrReportNotFound
}
