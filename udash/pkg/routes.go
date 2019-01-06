package pkg

import (
	"net/http"

	"github.com/dyweb/gommon/log/logx"
	"github.com/gorilla/mux"
)

// routers.go defines all the routes
// TODO: will switch to api package when it is ready

func (srv *Server) Handler() http.Handler {
	if srv.mux == nil {
		panic("call mountHandlers to init mux")
	}
	return srv.mux
}

func (srv *Server) LoggedHandler() http.Handler {
	return logx.NewHttpAccessLogger(srv.logger, srv.Handler())

}

func (srv *Server) mountHandlers() {
	r := mux.NewRouter()
	// a flat way to list all the routes
	r.HandleFunc("/api/local/docker/containers", srv.ListContainers)
	r.HandleFunc("/api/local/docker/images", srv.ListImages)
	srv.mux = r
}
