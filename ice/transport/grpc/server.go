package grpc

import (
	"net"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"
	ngrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/at15/go.ice/ice/config"
)

type Server struct {
	config config.GrpcServerConfig
	server *ngrpc.Server
	log    *dlog.Logger
}

func NewServer(cfg config.GrpcServerConfig, register func(server *ngrpc.Server)) (*Server, error) {
	srv := &Server{
		config: cfg,
	}
	srv.log = dlog.NewStructLogger(log, srv)
	var opts []ngrpc.ServerOption
	if cfg.Secure {
		srv.log.Infof("use tls with cert %s and key %s", cfg.Cert, cfg.Key)
		if creds, err := credentials.NewServerTLSFromFile(cfg.Cert, cfg.Key); err != nil {
			return nil, errors.Wrap(err, "can't generate grpc server credential from file")
		} else {
			opts = append(opts, ngrpc.Creds(creds))
		}
	}
	grpcServer := ngrpc.NewServer(opts...)
	register(grpcServer)
	srv.server = grpcServer
	return srv, nil
}

func (srv *Server) Run() error {
	cfg := srv.config
	srv.log.Infof("listen on %s", cfg.Addr)
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return errors.Wrap(err, "net can't listen tcp")
	}
	// reflection can be used by client to discover service, and can be used with cli
	reflection.Register(srv.server)
	srv.log.Info("start grpc server")
	if err := srv.server.Serve(lis); err != nil {
		return errors.Wrap(err, "grpc server can't serve")
	}
	return nil
}
