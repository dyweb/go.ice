package dockerclient

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/versions"
	"github.com/dyweb/go.ice/httpclient"
)

// image.go is merged all the image_*.go into one file

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
	return images, dc.h.GetTo(hCtx, "/images/json", &images)
}
