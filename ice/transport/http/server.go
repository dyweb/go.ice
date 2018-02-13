package http

import (
	nhttp "net/http"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"

	"context"
	"github.com/at15/go.ice/ice/config"
	"net"
	"strconv"
)

type Server struct {
	config config.HttpServerConfig
	server *nhttp.Server
	log    *dlog.Logger
}

// TODO: it should also has error, due to config etc.
func NewServer(cfg config.HttpServerConfig, h nhttp.Handler) *Server {
	// TODO： http server also accept stdlib logger
	httpServer := &nhttp.Server{
		Addr: cfg.Addr,
	}
	// TODO: this logic should be moved into access_log.go
	httpServer.Handler = nhttp.HandlerFunc(func(w nhttp.ResponseWriter, r *nhttp.Request) {
		tw := &TrackedWriter{w: w, status: 200}
		h.ServeHTTP(tw, r)
		// TODO: duration, but gommon/log can't handle float?
		// TODO: we should not be using package level logger, should use struct logger, or some special logger
		log.InfoF("http", dlog.Fields{
			dlog.Str("remote", r.RemoteAddr),
			dlog.Str("proto", r.Proto),
			dlog.Str("url", r.URL.String()),
			dlog.Int("size", tw.Size()),
			dlog.Int("status", tw.Status()),
		})
	})
	srv := &Server{
		config: cfg,
		server: httpServer,
	}
	srv.log = dlog.NewStructLogger(log, srv)
	return srv
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
