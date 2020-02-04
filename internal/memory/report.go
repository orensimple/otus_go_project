package memory

import (
	"context"
	"sync"

	"github.com/orensimple/otus_go_project/internal/domain/errors"
	"github.com/orensimple/otus_go_project/internal/domain/models"
)

type MemReportStorage struct {
	reports map[int64]map[int64]map[int64]*models.Stat
	mutex   *sync.Mutex
}

func NewMemReportStorage() *MemReportStorage {

	return &MemReportStorage{
		reports: make(map[int64]map[int64]map[int64]*models.Stat),
		mutex:   new(sync.Mutex),
	}
}

func (mem *MemReportStorage) UpdateReport(ctx context.Context, slotID int64, groupID int64, bannerID int64) error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()
	if _, ok := mem.reports[slotID][groupID][bannerID]; ok {

		return errors.ErrReportExist
	}
	var stat *models.Stat
	stat.Show = 0
	stat.Conversion = 0
	mem.reports[slotID][groupID][bannerID] = stat

	return nil
}

func (mem *MemReportStorage) AddClickToReport(ctx context.Context, slotID int64, groupID int64, bannerID int64) error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()
	if _, ok := mem.reports[slotID][groupID][bannerID]; ok {
		mem.reports[slotID][groupID][bannerID].Conversion++

		return nil
	}

	return errors.ErrReportNotFound
}

func (mem *MemReportStorage) AddShowToReport(ctx context.Context, slotID int64, groupID int64, bannerID int64) error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()
	if _, ok := mem.reports[slotID][groupID][bannerID]; ok {
		mem.reports[slotID][groupID][bannerID].Show++

		return nil
	}
	if _, ok := mem.reports[slotID]; !ok {
		mem.reports[slotID] = make(map[int64]map[int64]*models.Stat)
	}
	if _, ok := mem.reports[slotID][groupID]; !ok {
		mem.reports[slotID][groupID] = make(map[int64]*models.Stat)
	}

	stat := models.Stat{Show: 1, Conversion: 0}
	mem.reports[slotID][groupID][bannerID] = &stat

	return nil
}

func (mem *MemReportStorage) GetReports(ctx context.Context) (map[int64]map[int64]map[int64]*models.Stat, error) {

	return mem.reports, nil
}
