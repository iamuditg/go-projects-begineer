package main

import (
	"github.com/imuditg/handler"
	"github.com/imuditg/respository"
	"github.com/imuditg/server"
	"github.com/imuditg/service"
	"log"
)

func main() {
	// Create the download repository.
	downloadRepo := respository.NewDownloadRepository()

	// Create the download Service
	downloadService := service.NewDownloadService(downloadRepo)

	// Create the download handler.
	downloadHandler := handler.NewDownloadHandler(downloadService)

	// Create the Http server and start listening for requests.
	httpServer := server.NewHttpServer(":8080", downloadHandler)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
