package server

import (
	"context"

	ihttp "github.com/at15/go.ice/ice/transport/http"
	pb "github.com/at15/go.ice/example/github/pkg/icehubpb"
	"github.com/pkg/errors"
)

type HttpServer struct {
}

func (srv *HttpServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	return &pb.Pong{Name: "pong for ping " + ping.Name}, nil
}

func (srv *HttpServer) NoPayload(ctx context.Context, _ interface{}) (*pb.Pong, error) {
	return &pb.Pong{Name: "no payload"}, nil
}

func (srv *HttpServer) RegisterHandler(mux *ihttp.JsonHandlerMux) {
	mux.AddHandlerFunc("/ping", func() interface{} {
		return &pb.Ping{}
	}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		if ping, ok := req.(*pb.Ping); !ok {
			return nil, errors.New("invalid type, can't cast to *pb.Ping")
		} else {
			return srv.Ping(ctx, ping)
		}
	})
	mux.AddHandlerFunc("/nopayload", nil, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		return srv.NoPayload(ctx, nil)
	})
}
