package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
)

func (srv *Server) ListContainers(w http.ResponseWriter, res *http.Request) {
	// TODO: allow pass flags in from UI
	containers, err := srv.dc.ContainerList(res.Context(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		writeErr(w, err)
		return
	}
	writeJSON(w, containers)
}

func (srv *Server) ListImages(w http.ResponseWriter, res *http.Request) {
	images, err := srv.dc.ImageList(res.Context(), types.ImageListOptions{
		All: true,
	})
	if err != nil {
		writeErr(w, err)
		return
	}
	writeJSON(w, images)
}

func writeErr(w http.ResponseWriter, err error) {
	w.Write([]byte(err.Error()))
}

func writeJSON(w http.ResponseWriter, val interface{}) {
	if err := json.NewEncoder(w).Encode(val); err != nil {
		log.Warnf("error encode json to http response: %s", err)
	}
}
