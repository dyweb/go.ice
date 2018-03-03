package grpc

import (
	"context"
	"google.golang.org/grpc/peer"
)

func RemoteAddr(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	return p.Addr.String()
}
