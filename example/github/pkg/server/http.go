package server

import (
	"context"
	"net/http"

	pb "github.com/dyweb/go.ice/example/github/pkg/icehubpb"
	"github.com/dyweb/go.ice/example/github/pkg/server/auth"
	ihttp "github.com/dyweb/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
)

type HttpServer struct {
}

func NewHttpServer() (*HttpServer, error) {
	return &HttpServer{}, nil
}

func (srv *HttpServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	return &pb.Pong{Name: "pong for ping " + ping.Name}, nil
}

func (srv *HttpServer) NoPayload(ctx context.Context, _ interface{}) (*pb.Pong, error) {
	return &pb.Pong{Name: "no payload"}, nil
}

func (srv *HttpServer) Handler() http.Handler {
	mux := http.NewServeMux()
	jMux := ihttp.NewJsonHandlerMux()
	srv.RegisterHandler(jMux)
	jMux.MountToStd(mux)
	gh := auth.New()
	mux.HandleFunc("/github/login", http.HandlerFunc(gh.GitHubLogin))
	mux.HandleFunc("/github/callback", http.HandlerFunc(gh.GitHubLoginCallback))
	return mux
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
