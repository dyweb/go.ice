package pkg

import (
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
	// FIXME: it seems to be /bin/bash problem ...
	cmd := exec.Command("bash")
	//cmd := exec.Command("/usr/bin/
	tty, err := pty.Start(cmd)
	if err != nil {
		writeErr(w, err)
		return
	}
	defer tty.Close()

	// FIXME: the websocket is just echoing ... it is not executing bash
	//wrapper := wsWrapper{ws}
	go func() {
		// stdout
		//if _, err := io.Copy(tty, &wrapper); err != nil {
		//	log.Warnf("error read input from ws to tty: %s", err)
		//}
		//s := bufio.NewScanner(tty)
		//for s.Scan() {
		//	if err := ws.WriteMessage(websocket.TextMessage, s.Bytes()); err != nil {
		//		log.Warnf("write err: %s", err)
		//		ws.Close()
		//		break
		//	} else {
		//		log.Infof("scanned %s", s.Text())
		//	}
		//}
		//if s.Err() != nil {
		//	log.Warnf("scan err: %s", s.Err())
		//}
		for {
			buf := make([]byte, 200)
			n, err := tty.Read(buf)
			if err != nil {
				log.Warnf("read tty error: %s", err)
				break
			}
			log.Infof("read tty got %d %s", n, buf[:n])
		}
	}()
	//go func() {
	// stdin
	//if _, err := io.Copy(&wrapper, tty); err != nil {
	//	log.Warnf("error write output from tty to ws: %s", err)
	//}
	//}()
	go func() {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Warnf("error read ws message %s", err)
				break
			}
			// FIXED: this is the key have bash working ....
			message = append(message, '\n')
			if n, err := tty.Write(message); err != nil {
				log.Warnf("error write ws message to stdin %s", err)
				break
			} else {
				log.Infof("write %d into tty", n)
			}
		}
	}()
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
