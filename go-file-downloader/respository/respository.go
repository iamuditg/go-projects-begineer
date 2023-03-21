package respository

import (
	"github.com/imuditg/entity"
	"sync"
)

// DownloadRepository is an in-memory repository for downloads.
type DownloadRepository struct {
	sync.RWMutex
	download []*entity.Download
}

// NewDownloadRepository creates a new instance of DownloadRepository
func NewDownloadRepository() *DownloadRepository {
	return &DownloadRepository{}
}

// Save saves a download to the repository.
func (r *DownloadRepository) Save(d *entity.Download) error {
	r.Lock()
	defer r.Unlock()
	r.download = append(r.download, d)
	return nil
}

// GetAll returns all downloads in the repository.
func (r *DownloadRepository) GetAll() []*entity.Download {
	r.RLock()
	defer r.RUnlock()
	return r.download
}
