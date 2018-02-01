package server

import (
	"net/http"

	idb "github.com/at15/go.ice/ice/db"
	"fmt"
)

//type HTTPHandler struct {
//}
//
//func NewHTTP() *HTTPHandler {
//	return nil
//}

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
	return nil
}

func (r *Registry) HTTPHandler() http.Handler {
	return r.h
}
