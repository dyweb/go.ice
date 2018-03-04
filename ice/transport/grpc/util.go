package grpc

import (
	"context"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc/peer"
)

func RemoteAddr(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	return p.Addr.String()
}

func SplitHostPort(addr string) (string, int64) {
	host, ps, err := net.SplitHostPort(addr)
	if err != nil {
		log.Warnf("failed to split host port %s %v", addr, err)
		return host, 0
	}
	// TODO: protobuf generated struct has omit empty ... which would leave bind ip as blank, so we fill it with 0.0.0.0
	if host == "" {
		host = "0.0.0.0"
	}
	p, err := strconv.Atoi(ps)
	if err != nil {
		log.Warnf("failed to convert port number %s to int %v", ps, err)
		return host, int64(p)
	}
	return host, int64(p)
}

func Hostname() string {
	if host, err := os.Hostname(); err != nil {
		log.Warnf("can't get hostname %v", err)
		return "unknown"
	} else {
		return host
	}
}
