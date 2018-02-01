package server

import (
	"fmt"
	"net/http"

	"github.com/at15/go.ice/_example/github/pkg/server/auth"
	idb "github.com/at15/go.ice/ice/db"
)

type Registry struct {
	h  *http.ServeMux
	db *idb.Manager
}

// TODO: entity etc.
// TODO: return http handler and grpc handler?

func NewRegistry(cfg Config) (*Registry, error) {
	r := &Registry{}
	r.db = idb.NewManager(cfg.DatabaseManager)
	return r, nil
}

func (r *Registry) ConfigHttpHandler() error {
	r.h = http.NewServeMux()
	r.h.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong\n")
	})
	// FIXME: hack to get gh up and running
	gh := auth.NewGh()
	r.h.HandleFunc("/start", gh.Start)
	r.h.HandleFunc("/cb", gh.Cb)
	return nil
}

func (r *Registry) HTTPHandler() http.Handler {
	return r.h
}
