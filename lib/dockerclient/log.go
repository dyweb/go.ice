package dockerclient

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	dockertime "github.com/docker/docker/api/types/time"
	"github.com/dyweb/go.ice/httpclient"
	"github.com/dyweb/gommon/errors"
)

// log.go is container log with handy wrapper func

// https://github.com/moby/moby/blob/master/client/container_logs.go
// https://github.com/docker/cli/blob/master/cli/command/container/logs.go
// https://docs.docker.com/engine/reference/commandline/logs/#options
//
// docker run --name test -d busybox sh -c "while true; do $(echo date); sleep 1; done"
func (dc *Client) ContainerLog(ctx context.Context, containerNameOrId string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	hCtx := httpclient.ConvertContext(ctx)
	if options.ShowStdout {
		hCtx.SetParam("stdout", "1")
	}

	if options.ShowStderr {
		hCtx.SetParam("stderr", "1")
	}

	if options.Since != "" {
		ts, err := dockertime.GetTimestamp(options.Since, time.Now())
		if err != nil {
			return nil, errors.Wrap(err, `invalid value for "since"`)
		}
		hCtx.SetParam("since", ts)
	}

	if options.Until != "" {
		ts, err := dockertime.GetTimestamp(options.Until, time.Now())
		if err != nil {
			return nil, errors.Wrap(err, `invalid value for "until"`)
		}
		hCtx.SetParam("until", ts)
	}

	if options.Timestamps {
		hCtx.SetParam("timestamps", "1")
	}

	if options.Details {
		hCtx.SetParam("details", "1")
	}

	if options.Follow {
		hCtx.SetParam("follow", "1")
	}
	hCtx.SetParam("tail", options.Tail)
	res, err := dc.h.GetRaw(hCtx, "/containers/"+containerNameOrId+"/logs")
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
