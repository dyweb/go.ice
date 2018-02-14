package http

import (
	"net/http"

	dlog "github.com/dyweb/gommon/log"
)

// TODO: panic, cors handler, and combination of handlers ...

func NewLoggingHandler(h http.Handler, log *dlog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tw := &TrackedWriter{w: w, status: 200}
		h.ServeHTTP(tw, r)
		// TODO: duration, but gommon/log can't handle float?
		log.InfoF("http", dlog.Fields{
			dlog.Str("remote", r.RemoteAddr),
			dlog.Str("proto", r.Proto),
			dlog.Str("url", r.URL.String()),
			dlog.Int("size", tw.Size()),
			dlog.Int("status", tw.Status()),
		})
	})
}
