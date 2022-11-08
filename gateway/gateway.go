package gateway

import (
	"net/http"
)

type APIServer interface {
	Serve() error
	AddRoute(string, func(http.ResponseWriter, *http.Request))
}

type apiServer struct {
	Server *http.Server
	mux    *http.ServeMux
}

func (s *apiServer) Serve() error {
	err := s.Server.ListenAndServe()
	return err
}

func (s *apiServer) AddRoute(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(path, handler)
}

func InitAPIServer(addr string) APIServer {
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &apiServer{
		Server: s,
		mux:    mux,
	}
}
