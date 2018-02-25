package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonFunc func(ctx context.Context, req interface{}) (res interface{}, err error)

type JsonHandler interface {
	HasPayload() bool
	NewPayload() interface{}
	Func() JsonFunc
}

type JsonHandlerMux struct {
	handlers map[string]JsonHandler
}

type JsonHandlerRegister func(mux *JsonHandlerMux)

func (m *JsonHandlerMux) AddHandlerFunc(path string, payloadFactory func() interface{}, f JsonFunc) {
	h := &jsonHandler{f: f}
	if payloadFactory() == nil {
		h.hasPayload = false
	} else {
		h.hasPayload = true
		h.payloadFactory = payloadFactory
	}
	m.handlers[path] = h
}

func (m *JsonHandlerMux) MountToStd(mux *http.ServeMux) {
	for path, h := range m.handlers {
		if !h.HasPayload() {
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				res, err := h.Func()(ctx, nil)
				if err != nil {
					jsonInternalError(w, err)
				}
				jsonRes(w, res)
			})
		} else {
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				req := h.NewPayload()
				if err := json.NewDecoder(r.Body).Decode(req); err != nil {
					jsonInvalidFormat(w, err)
				}
				res, err := h.Func()(ctx, req)
				if err != nil {
					jsonInternalError(w, err)
				}
				jsonRes(w, res)
			})
		}
	}
}

var _ JsonHandler = (*jsonHandler)(nil)

type jsonHandler struct {
	hasPayload     bool
	payloadFactory func() interface{}
	f              JsonFunc
}

func (h *jsonHandler) HasPayload() bool {
	return h.hasPayload
}

func (h *jsonHandler) NewPayload() interface{} {
	return h.payloadFactory()
}

func (h *jsonHandler) Func() JsonFunc {
	return h.Func()
}

func jsonInvalidFormat(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	// TODO: escape
	fmt.Fprintf(w, `{"msg":"invalid json", "err":"%v"`, err)
}

func jsonInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	// TODO: escape
	fmt.Fprintf(w, `{"msg":"internal error", "err":"%v"`, err)
}

func jsonRes(w http.ResponseWriter, res interface{}) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(res); err != nil {
		// TODO: should not use package level logger
		log.Warnf("error encoding json %v", err)
		jsonInternalError(w, err)
		return
	}
	if _, err := w.Write(buf.Bytes()); err != nil {
		// TODO: should not use package level logger
		log.Warnf("can't write to http connection %v", err)
	}
}
