package http

import "net/http"

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) Handler() http.Handler {
	if srv.mux != nil {
		return srv.mux
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong\n"))
	})
	srv.mux = mux
	return srv.mux
}
