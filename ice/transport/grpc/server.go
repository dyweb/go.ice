package grpc

import (
	"context"
	"net"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/at15/go.ice/ice/config"
)

type Server struct {
	config config.GrpcServerConfig
	server *grpc.Server

	log *dlog.Logger
}

func NewServer(cfg config.GrpcServerConfig, register func(s *grpc.Server)) (*Server, error) {
	if cfg.ShutdownDuration <= 0 {
		cfg.ShutdownDuration = config.ShutdownDuration
	}
	srv := &Server{
		config: cfg,
	}
	srv.log = dlog.NewStructLogger(log, srv)
	var opts []grpc.ServerOption
	if cfg.Secure {
		srv.log.Infof("use tls with cert %s and key %s", cfg.Cert, cfg.Key)
		if creds, err := credentials.NewServerTLSFromFile(cfg.Cert, cfg.Key); err != nil {
			return nil, errors.Wrap(err, "can't generate grpc server credential from file")
		} else {
			opts = append(opts, grpc.Creds(creds))
		}
	}
	grpcServer := grpc.NewServer(opts...)
	register(grpcServer)
	srv.server = grpcServer
	return srv, nil
}

func (srv *Server) Run() error {
	cfg := srv.config
	srv.log.Infof("grpc listen on %s", cfg.Addr)
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

func (srv *Server) RunWithContext(ctx context.Context) error {
	waitCh := make(chan error)
	go func() {
		err := srv.Run()
		waitCh <- err
	}()
	select {
	case err := <-waitCh:
		return err
	case <-ctx.Done():
		merr := errors.NewMultiErr()
		merr.Append(ctx.Err())
		shutCtx, _ := context.WithTimeout(context.Background(), srv.config.ShutdownDuration)
		merr.Append(srv.Shutdown(shutCtx))
		return merr.ErrorOrNil()
	}
}

func (srv *Server) Shutdown(ctx context.Context) error {
	srv.log.Info("graceful shutdown grpc server")
	// TODO: make use of context
	srv.server.GracefulStop()
	return nil
}
