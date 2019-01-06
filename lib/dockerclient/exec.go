package dockerclient

import (
	"context"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/dyweb/go.ice/httpclient"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/httputil"
)

// https://github.com/docker/cli/blob/master/cli/command/container/exec.go
// https://github.com/moby/moby/blob/master/client/container_exec.go
func (dc *Client) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error) {
	hCtx := httpclient.ConvertContext(ctx)
	var id types.IDResponse
	if err := dc.h.Post(hCtx, "/containers/"+container+"/exec", config, &id); err != nil {
		return id, err
	}
	return id, nil
}

// TODO: Attach is using the same start API but hijack the raw stream
// NOTE: docker is using the deprecated httputil.ClientConn
// https://github.com/moby/moby/blob/master/client/hijack.go#L53
func (dc *Client) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (net.Conn, error) {
	tr, ok := dc.h.Transport()
	if !ok {
		return nil, errors.New("can't get underlying http.Transport")
	}
	hCtx := httpclient.ConvertContext(ctx)
	hCtx.SetHeader("Connection", "Upgrade").SetHeader("Upgrade", "tcp")
	req, err := dc.h.NewRequest(hCtx, httputil.Post, "/exec/"+execID+"/start", config)
	// TODO: if it's unix, then the network and addr are actually ignored ...
	// TODO: but for tcp it is needed ...
	// TODO: actually just create a net.Dialer is fine ...
	conn, err := tr.DialContext(ctx, "", "")
	if err != nil {
		return nil, err
	}
	if err := req.Write(conn); err != nil {
		return nil, err
	}
	// TODO: actually we should read response https://golang.org/pkg/net/http/#ReadResponse
	return conn, nil
}

// TODO: Resize and Inspect
