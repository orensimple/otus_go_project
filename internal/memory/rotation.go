package memory

import (
	"context"
	"sync"

	"github.com/orensimple/otus_go_project/internal/domain/errors"
)

type MemRotationStorage struct {
	rotation map[int64][]int64
	mutex    *sync.Mutex
}

func NewMemRotationStorage() *MemRotationStorage {

	return &MemRotationStorage{
		rotation: make(map[int64][]int64),
		mutex:    new(sync.Mutex),
	}
}

func (mem *MemRotationStorage) SetRotation(ctx context.Context, slotID int64, bannerID int64) error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()
	if !contains(mem.rotation[slotID], bannerID) {
		mem.rotation[slotID] = append(mem.rotation[slotID], bannerID)
	}

	return nil
}

func (mem *MemRotationStorage) GetRotations(ctx context.Context, slotID int64, groupID int64) ([]int64, error) {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()

	bannersID, ok := mem.rotation[slotID]
	if !ok {

		return nil, errors.ErrReportNotFound
	}

	return bannersID, nil
}

func (mem *MemRotationStorage) DeleteRotation(ctx context.Context, slotID int64, bannerID int64) (int64, error) {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()

	bannersID, ok := mem.rotation[slotID]
	if ok {
		for i, v := range bannersID {
			if v == bannerID {
				copy(bannersID[i:], bannersID[i+1:])
				bannersID[len(bannersID)-1] = 0 // обнуляем "хвост"
				mem.rotation[slotID] = bannersID[:len(bannersID)-1]
			}
		}
		return slotID, nil
	}

	return 0, errors.ErrReportNotFound
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
