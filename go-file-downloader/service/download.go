package service

import (
	"errors"
	"github.com/imuditg/entity"
	"github.com/imuditg/respository"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// DownloadService is a service that can download files.
type DownloadService struct {
	downloadRepo respository.DownloadRepository
}

// NewDownloadService created a new instance of DownloadService.
func NewDownloadService(downloadRepo *respository.DownloadRepository) *DownloadService {
	return &DownloadService{downloadRepo: *downloadRepo}
}

// Download downloads the files at the specified URLs.
func (s *DownloadService) Download(urls []string) ([]*entity.DownloadResult, error) {
	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	results := make(chan *entity.DownloadResult, len(urls))
	for _, url := range urls {
		go func(u string) {
			defer wg.Done()

			filename := u[strings.LastIndex(u, "/")+1:]
			download := &entity.Download{
				URL:      u,
				Filename: filename,
				Retries:  3,
			}
			err := s.downloadRepo.Save(download)
			if err != nil {
				results <- &entity.DownloadResult{
					Download: *download,
					Error:    err,
				}
				return
			}

			res, err := downloadFile(download, 3)
			if err != nil {
				results <- &entity.DownloadResult{
					Download: *download,
					Error:    err,
				}
			}
			results <- res
		}(url)
	}
	wg.Wait()
	close(results)

	var downloadResults []*entity.DownloadResult
	for r := range results {
		downloadResults = append(downloadResults, r)
	}
	return downloadResults, nil
}

func downloadFile(d *entity.Download, maxRetries int) (*entity.DownloadResult, error) {
	retryCount := 0
	for {
		err := downloadSingleFile(d)
		if err == nil {
			return &entity.DownloadResult{
				Download: *d,
			}, nil
		}
		if !isRetryable(err) {
			return &entity.DownloadResult{
				Download: *d,
				Error:    err,
			}, nil
		}
		if retryCount >= maxRetries {
			return &entity.DownloadResult{
				Download: *d,
				Error:    err,
			}, nil
		}
		retryCount++
	}
}

func downloadSingleFile(d *entity.Download) error {
	resp, err := http.Get(d.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(d.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func isRetryable(err error) bool {
	if err == nil {
		return false
	}

	var retryableErrors = []error{
		errors.New("network error"),
		errors.New("server error"),
	}
	for _, e := range retryableErrors {
		if e.Error() == err.Error() {
			return true
		}
	}
	return false
}
