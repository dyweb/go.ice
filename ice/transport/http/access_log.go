package http

import "net/http"

// HTTP Access Log
// TODO: might rename this to writer, and create a handler.go for logging handler

var _ http.ResponseWriter = (*TrackedWriter)(nil)
var _ http.Flusher = (*TrackedWriter)(nil)
var _ http.Pusher = (*TrackedWriter)(nil)
var _ http.CloseNotifier = (*TrackedWriter)(nil)

// TrackedWriter keeps track of status code and bytes written so it can be used by logger.
// It proxies all the interfaces except Hijacker, since it is not supported by HTTP/2.
// Most methods comments are copied from net/http
// It is based on https://github.com/gorilla/handlers but put all interface on one struct
type TrackedWriter struct {
	w      http.ResponseWriter
	status int
	size   int
}

// Header returns the header map of the underlying ResponseWriter
//
// Changing the header map after a call to WriteHeader (or
// Write) has no effect unless the modified headers are
// trailers.
func (tracker *TrackedWriter) Header() http.Header {
	return tracker.w.Header()
}

// Write keeps track of bytes written of the underlying ResponseWriter
//
// Write writes the data to the connection as part of an HTTP reply.
//
// If WriteHeader has not yet been called, Write calls
// WriteHeader(http.StatusOK) before writing the data. If the Header
// does not contain a Content-Type line, Write adds a Content-Type set
// to the result of passing the initial 512 bytes of written data to
// DetectContentType.
func (tracker *TrackedWriter) Write(b []byte) (int, error) {
	size, err := tracker.w.Write(b)
	tracker.size += size
	return size, err
}

// WriteHeader keep track of status code and call the underlying ResponseWriter
//
// WriteHeader sends an HTTP response header with status code.
// If WriteHeader is not called explicitly, the first call to Write
// will trigger an implicit WriteHeader(http.StatusOK).
// Thus explicit calls to WriteHeader are mainly used to
// send error codes.
func (tracker *TrackedWriter) WriteHeader(status int) {
	tracker.status = status
	tracker.w.WriteHeader(status)
}

// Status return the tracked status code, returns 0 if WriteHeader has not been called
func (tracker *TrackedWriter) Status() int {
	return tracker.status
}

// Size return number of bytes written through Write, returns 0 if Write has not been called
func (tracker *TrackedWriter) Size() int {
	return tracker.size
}

// Flush calls Flush on underlying ResponseWriter if it implemented http.Flusher
//
// Flusher interface is implemented by ResponseWriters that allow
// an HTTP handler to flush buffered data to the client.
// The default HTTP/1.x and HTTP/2 ResponseWriter implementations
// support Flusher
func (tracker *TrackedWriter) Flush() {
	if f, ok := tracker.w.(http.Flusher); ok {
		f.Flush()
	}
}

// Push returns http.ErrNotSupported if underlying ResponseWriter does not implement http.Pusher
//
// Push initiates an HTTP/2 server push, returns ErrNotSupported if the client has disabled push or if push
// is not supported on the underlying connection.
func (tracker *TrackedWriter) Push(target string, opts *http.PushOptions) error {
	if p, ok := tracker.w.(http.Pusher); ok {
		return p.Push(target, opts)
	}
	return http.ErrNotSupported
}

// CloseNotify returns nil if underlying ResponseWriter does not implement http.CloseNotifier
//
// The CloseNotifier interface is implemented by ResponseWriters which
// allow detecting when the underlying connection has gone away
func (tracker *TrackedWriter) CloseNotify() <-chan bool {
	if c, ok := tracker.w.(http.CloseNotifier); ok {
		return c.CloseNotify()
	}
	// FIXME: what should be returned, may nil would work? .. never used CloseNotify before
	return nil
}
