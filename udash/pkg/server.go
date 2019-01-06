package pkg

import (
	"net/http"
	"time"

	"github.com/dyweb/go.ice/lib/dockerclient"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
)

var log, logReg = dlog.NewPackageLoggerAndRegistryWithSkip("udash", 0)

type Server struct {
	mux http.Handler
	dc  *dockerclient.Client

	logger *dlog.Logger
}

func NewServer() (*Server, error) {
	dc, err := dockerclient.New(dockerclient.DefaultLocalHost)
	if err != nil {
		return nil, err
	}
	srv := Server{
		dc: dc,
	}
	dlog.NewStructLogger(log, &srv)
	srv.mountHandlers()
	return &srv, nil
}

func (srv *Server) Run(addr string) error {
	httpd := http.Server{
		Addr:    addr,
		Handler: srv.LoggedHandler(),
	}
	srv.logger.Infof("listen on %s", addr)
	// wait for 3s
	go func() {
		time.Sleep(3 * time.Second)
		srv.logger.Infof("server started on %s", addr)
	}()
	if err := httpd.ListenAndServe(); err != nil {
		return errors.Wrapf(err, "error listen on %s", addr)
	}
	// TODO: will this section get called if the server is shutdown gracefully?
	return nil
}
