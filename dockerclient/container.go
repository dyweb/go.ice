package dockerclient

import (
	"context"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/httpclient"
)

// TODO
// - start
// - stop
// - kill

type configWrapper struct {
	*container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
}

// https://docs.docker.com/engine/api/sdk/examples/#run-a-container
// https://github.com/docker/cli/blob/master/cli/command/container/run.go
func (dc *Client) ContainerCreate(ctx context.Context, config *container.Config,
	hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig,
	containerName string) (container.ContainerCreateCreatedBody, error) {
	hCtx := httpclient.ConvertContext(ctx)
	if containerName != "" {
		hCtx.SetParam("name", containerName)
	}
	body := configWrapper{
		Config:           config,
		HostConfig:       hostConfig,
		NetworkingConfig: networkingConfig,
	}

	var created container.ContainerCreateCreatedBody
	err := dc.h.Post(hCtx, "/containers/create", body, &created)
	return created, err
}

// https://github.com/docker/cli/blob/master/cli/command/container/list.go
// https://github.com/moby/moby/blob/master/client/container_list.go
// https://docs.docker.com/engine/reference/commandline/ps/#usage
func (dc *Client) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	hCtx := httpclient.ConvertContext(ctx)

	if options.All {
		hCtx.SetParam("all", "1")
	}
	if options.Limit != -1 {
		hCtx.SetParam("limit", strconv.Itoa(options.Limit))
	}
	if options.Since != "" {
		hCtx.SetParam("since", options.Since)
	}
	if options.Before != "" {
		hCtx.SetParam("before", options.Before)
	}
	if options.Size {
		hCtx.SetParam("size", "1")
	}
	if options.Filters.Len() > 0 {
		if filterJSON, err := filters.ToJSON(options.Filters); err != nil {
			return nil, err
		} else {
			hCtx.SetParam("filters", filterJSON)
		}
	}

	var containers []types.Container
	if err := dc.h.Get(hCtx, "/containers/json", &containers); err != nil {
		return nil, err
	}
	return containers, nil
}

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
