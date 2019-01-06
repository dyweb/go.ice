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

// TODO: config upgrader and limit its scope to the server instance
var upgrader = websocket.Upgrader{}

func (srv *Server) HostShell(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		writeErr(w, err)
		return
	}

	defer ws.Close()
	cmd := exec.Command("bash")
	tty, err := pty.Start(cmd)
	if err != nil {
		writeErr(w, err)
		return
	}
	defer tty.Close()

	go func() {
		// read from stdout and write to websocket, no need to buffer and add `\n`
		// NOTE: pty has echo, that's how you see what you type
		// (also change terminal mode to not echo back your password in plain text)
		for {
			// TODO: handle buffer recycle and buffer overflow
			b := make([]byte, 1000)
			n, err := tty.Read(b)
			if err != nil {
				log.Warnf("read err: %s", err)
				ws.Close()
				break
			}
			if n > 0 {
				if err := ws.WriteMessage(websocket.TextMessage, b[:n]); err != nil {
					log.Warnf("write err: %s", err)
					ws.Close()
					break
				}
			}
		}
	}()
	go func() {
		//var buf bytes.Buffer
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Warnf("error read ws message %s", err)
				break
			}
			if len(message) == 0 {
				continue
			}
			//buf.Write(message)
			//if message[len(message) - 1] != '\n' {
			//	continue
			//}
			//if message[len(message)-1] != '\n' {
			//	// FIXED: this is the key to have bash working, must have a trailing \n
			//	message = append(message, '\n')
			//}
			if _, err := tty.Write(message); err != nil {
				log.Warnf("error write ws message to stdin %s", err)
				break
			}
		}
	}()
	if err := cmd.Wait(); err != nil {
		log.Warnf("error wait cmd: %s", err)
	}
	srv.logger.Info("closed")
}
