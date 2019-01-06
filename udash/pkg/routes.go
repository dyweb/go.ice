package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routers.go defines all the routes
// TODO: will switch to api package when it is ready

func (srv *Server) Handler() http.Handler {
	if srv.mux == nil {
		panic("call mountHandlers to init mux")
	}
	return srv.mux
}

//func (srv *Server) LoggedHandler() http.Handler {
//	// TODO: the traced writer need a factory method ...
//	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
//		//tracker := httputil.TrackedWriter{
//		//	w: res,
//		//}
//		//start := time.Now()
//		//l.mux.ServeHTTP(&tracker, req)
//		//duration := time.Now().Sub(start)
//		//l.logger.InfoF(fmt.Sprintf("%d %s %s", tracker.Status(), req.Method, req.URL.Path),
//		//	dlog.Int("status", tracker.Status()),
//		//	dlog.Int("size", tracker.Size()),
//		//	dlog.Str("duration", duration.String()),
//		//	dlog.Str("method", req.Method),
//		//	dlog.Str("url", req.URL.String()),
//		//	dlog.Str("remote", req.RemoteAddr),
//		//	dlog.Str("refer", req.Referer()), // TODO: refer could be empty, not sure if dlog can handle it properly ...
//		//	dlog.Str("proto", req.Proto),
//		//)
//	})
//
//}

func (srv *Server) mountHandlers() {
	r := mux.NewRouter()
	// a flat way to list all the routes
	r.HandleFunc("/api/local/docker/images", srv.ListImages)
	srv.mux = r
}
