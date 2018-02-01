package http

import (
	nhttp "net/http"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"

	"context"
	"github.com/at15/go.ice/ice/config"
)

type Server struct {
	config config.HttpServerConfig
	server *nhttp.Server
	log    *dlog.Logger
}

func NewServer(cfg config.HttpServerConfig, h nhttp.Handler) *Server {
	s := &Server{
		config: cfg,
		server: &nhttp.Server{Addr: cfg.Addr, Handler: h},
	}
	s.log = dlog.NewStructLogger(log, s)
	return s
}

func (srv *Server) Port() int {
	// TODO: split port from addr
	return 0
}

// TODO: check if handler is nil?
func (srv *Server) Run() error {
	cfg := srv.config
	srv.log.Infof("listen on %s", cfg.Addr)
	if cfg.Secure {
		srv.log.Infof("use tls with cert %s and key %s", cfg.Cert, cfg.Key)
		if err := srv.server.ListenAndServeTLS(cfg.Cert, cfg.Key); err != nil {
			return errors.Wrap(err, "can't start http tls server")
		}
	} else {
		srv.log.Infof("start http server without TLS")
		if err := srv.server.ListenAndServe(); err != nil {
			return errors.Wrap(err, "can't start http server")
		}
	}
	return nil
}

func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.server.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "can't shutdown http server gracefully")
	}
	return nil
}
