package dockerclient

import (
	"context"
	"io"
	"strings"

	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/versions"
	"github.com/dyweb/gommon/httpclient"
)

// image.go is merged all the image_*.go into one file

// https://github.com/moby/moby/blob/master/client/image_pull.go difference between pull and create is pull try to auth
// https://github.com/moby/moby/blob/master/client/image_create.go
func (dc *Client) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {

	hCtx := httpclient.ConvertContext(ctx)

	ref, err := reference.ParseNormalizedNamed(refStr)
	if err != nil {
		return nil, err
	}
	hCtx.SetParam("fromImage", reference.FamiliarName(ref))
	hCtx.SetParam("tag", getAPITagFromNamedRef(ref))
	if options.Platform != "" {
		hCtx.SetParam("platform", strings.ToLower(options.Platform))
	}
	// TODO: handle auth, this is needed to pull from private registry
	res, err := dc.h.PostRaw(hCtx, "/images/create", nil)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

// https://github.com/moby/moby/blob/master/client/image_list.go
func (dc *Client) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	var images []types.ImageSummary

	hCtx := httpclient.ConvertContext(ctx)
	optionFilters := options.Filters
	referenceFilters := optionFilters.Get("reference")
	if versions.LessThan(dc.version, "1.25") && len(referenceFilters) > 0 {
		hCtx.SetParam("filter", referenceFilters[0])
		for _, filterValue := range referenceFilters {
			optionFilters.Del("reference", filterValue)
		}
	}
	if optionFilters.Len() > 0 {
		filterJSON, err := filters.ToJSON(optionFilters)
		if err != nil {
			return images, err
		}
		hCtx.SetParam("filters", filterJSON)
	}
	if options.All {
		hCtx.SetParam("all", "1")
	}
	return images, dc.h.Get(hCtx, "/images/json", &images)
}

// getAPITagFromNamedRef returns a tag from the specified reference.
// This function is necessary as long as the docker "server" api expects
// digests to be sent as tags and makes a distinction between the name
// and tag/digest part of a reference.
func getAPITagFromNamedRef(ref reference.Named) string {
	if digested, ok := ref.(reference.Digested); ok {
		return digested.Digest().String()
	}
	ref = reference.TagNameOnly(ref)
	if tagged, ok := ref.(reference.Tagged); ok {
		return tagged.Tag()
	}
	return ""
}
