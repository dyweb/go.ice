package pkg

import (
	"io"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/kr/pty"
)

// host.go allow user to run shell command on host

func (srv *Server) HostShellStatic(w http.ResponseWriter, r *http.Request) {
	// FIXME: hard coded html path
	http.ServeFile(w, r, "/home/at15/workspace/src/github.com/dyweb/go.ice/udash/pkg/shell.html")
}

func (srv *Server) HostShell(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		writeErr(w, err)
		return
	}

	defer ws.Close()
	cmd := exec.Command("/bin/bash")
	tty, err := pty.Start(cmd)
	if err != nil {
		writeErr(w, err)
		return
	}
	defer tty.Close()

	// FIXME: the websocket is just echoing ... it is not executing bash
	wrapper := wsWrapper{ws}
	go func() {
		// stdout
		if _, err := io.Copy(tty, &wrapper); err != nil {
			log.Warnf("error read input from ws to tty: %s", err)
		}
	}()
	//go func() {
	// stdin
	if _, err := io.Copy(&wrapper, tty); err != nil {
		log.Warnf("error write output from tty to ws: %s", err)
	}
	//}()

	if err := cmd.Wait(); err != nil {
		log.Warnf("error wait cmd: %s", err)
	}
	srv.logger.Info("closed")
}

var upgrader = websocket.Upgrader{}

// FIXME: copied from gotty
type wsWrapper struct {
	*websocket.Conn
}

func (wsw *wsWrapper) Write(p []byte) (n int, err error) {
	log.Infof("I am going to write %s", string(p))
	writer, err := wsw.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	defer writer.Close()
	return writer.Write(p)
}

func (wsw *wsWrapper) Read(p []byte) (n int, err error) {
	for {
		msgType, reader, err := wsw.Conn.NextReader()
		if err != nil {
			return 0, err
		}

		if msgType != websocket.TextMessage {
			continue
		}

		return reader.Read(p)
	}
}
