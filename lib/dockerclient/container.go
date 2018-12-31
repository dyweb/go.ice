package dockerclient

import (
	"context"

	"github.com/dyweb/go.ice/httpclient"
	"github.com/pkg/errors"
)

// TODO
// - list
// - start
// - stop
// - kill

// TODO: signal should be typed
// TODO: kill -l to list all the signals
// https://www.linux.org/threads/kill-commands-and-signals.8881/
// https://github.com/docker/cli/blob/master/cli/command/container/kill.go
// https://github.com/moby/moby/blob/master/client/container_kill.go
func (dc *Client) ContainerKill(ctx context.Context, containerId, signal string) error {
	hCtx := httpclient.ConvertContext(ctx)
	if signal == "" {
		signal = "KILL"
	}
	if containerId == "" {
		return errors.New("containerId is empty for container kill")
	}
	if err := dc.h.PostIgnoreRes(hCtx, "/containers/"+containerId+"/kill", nil); err != nil {
		return err
	}
	return nil
}
