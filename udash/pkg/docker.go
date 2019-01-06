package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
)

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
