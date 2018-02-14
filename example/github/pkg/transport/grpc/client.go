package grpc

import "google.golang.org/grpc"

// NOTE: client does not need any code implementation, unless there need to be hack, we can embed the auto generated client

func NewClient(con *grpc.ClientConn) IceHubClient {
	return NewIceHubClient(con)
}
