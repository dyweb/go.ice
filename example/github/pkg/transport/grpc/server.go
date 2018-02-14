package grpc

import (
	"context"
	icehub "github.com/at15/go.ice/example/github/pkg/icehubpb"
)

// check it implemented the interface
var _ IceHubServer = (*Server)(nil)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) Ping(context.Context, *icehub.Ping) (*icehub.Pong, error) {
	return &icehub.Pong{Name: "pong"}, nil
}
