package entity

// Download represents a single file download operation.
type Download struct {
	URL      string // the URL of the file to download
	Filename string // the name of the file to save the downloaded contents to.
	Retries  int    // the number of times to retry downloading the file in case of errors.
}

// DownloadResult represents the result of a file download operation.
type DownloadResult struct {
	Download Download
	Error    error
}
