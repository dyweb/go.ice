package http

import (
	"context"
	"net"
	"net/http"
	"strconv"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"

	"github.com/at15/go.ice/ice/config"
	"github.com/opentracing/opentracing-go"
)

type Server struct {
	config config.HttpServerConfig
	server *http.Server
	log    *dlog.Logger
	tracer opentracing.Tracer
}

// TODO: check if there is any error in config and return error
func NewServer(cfg config.HttpServerConfig, h http.Handler, tracer opentracing.Tracer) (*Server, error) {
	if cfg.EnableTracing && tracer == nil {
		return nil, errors.New("tracer is nil but tracing is enabled")
	}
	srv := &Server{
		config: cfg,
		tracer: tracer,
	}
	srv.log = dlog.NewStructLogger(log, srv)
	// TODO： http server also accept stdlib logger, we might hijack it ...
	httpServer := &http.Server{
		Addr: cfg.Addr,
	}
	// TODO: there could be more than just logging handler, panic, cors etc.
	// TODO: http log might need special logger, we are using struct's logger for now...
	httpServer.Handler = NewLoggingHandler(h, srv.log)
	srv.server = httpServer
	return srv, nil
}

func (srv *Server) Port() int {
	var (
		port  int
		portS string
		err   error
	)
	if _, portS, err = net.SplitHostPort(srv.config.Addr); err != nil {
		return 0
	}
	if port, err = strconv.Atoi(portS); err != nil {
		return 0
	} else {
		return port
	}
}

func (srv *Server) Run() error {
	if srv.server.Handler == nil {
		return errors.New("nil handler")
	}
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
