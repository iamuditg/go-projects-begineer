package handler

import (
	"encoding/json"
	"github.com/imuditg/service"
	"net/http"
)

// DownloadHandler is an HTTP handler that downloads files
type DownloadHandler struct {
	downloadService service.DownloadService
}

// NewDownloadHandler creates a new instance of DownloadHandler.
func NewDownloadHandler(downloadService *service.DownloadService) *DownloadHandler {
	return &DownloadHandler{downloadService: *downloadService}
}

// ServeHTTP handles HTTP requests to download files.
func (h *DownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var urls []string
	err := decoder.Decode(&urls)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	results, err := h.downloadService.Download(urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"result": results,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "failed to marshal json response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
