package server

import "net/http"

// HttpServer is an HTTP server
type HttpServer struct {
	addr    string
	handler http.Handler
	server  *http.Server
}

// NewHttpServer creates a new instance of HTTP Server
func NewHttpServer(addr string, handler http.Handler) *HttpServer {
	return &HttpServer{
		addr:    addr,
		handler: handler,
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

// ListenAndServe starts the Http server and listen for incoming requests.
func (s *HttpServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}
