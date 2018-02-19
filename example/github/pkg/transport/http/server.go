package http

import (
	"github.com/at15/go.ice/example/github/pkg/server/auth"
	"net/http"
)

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
	gh := auth.New()
	mux.HandleFunc("/github/login", http.HandlerFunc(gh.GitHubLogin))
	mux.HandleFunc("/github/callback", http.HandlerFunc(gh.GitHubLoginCallback))
	srv.mux = mux
	return srv.mux
}
