# 2019-01-06 Host Shell

- [x] httputil didn't implement hijacker interface

Debug websocket

- https://kaazing.com/inspecting-websocket-traffic-with-chrome-developer-tools/
- remember to **increase height of developer tools**

The following code seems to be copying sames things in and out ...

````go
	wrapper := wsWrapper{ws}
	go func() {
		if _, err := io.Copy(&wrapper, tty); err != nil {
			log.Warnf("error write output from tty to ws: %s", err)
		}
	}()
	go func() {
		if _, err := io.Copy(tty, &wrapper); err != nil {
			log.Warnf("error read input from ws to tty: %s", err)
		}
	}()
````

- the key for having bash working is add `\n`

How the server works actually also depends on how the client is implemented

- if you use a simple text box instead of a web terminal like xterm.js, then you need to
  - add a `\n` after the input because it is a full message and need `\n` to trigger bash's response (bash has a buffer of stdin and only take action when `\n` shows up I suppose)
- pty has the echo behavior, which is why you can see what you type
  - also why type in xterm does not shown anything ...
  
When use xterm.js all those backend logic is wrong

- no need to use bufio for output, it will break the echo logic, you can't see what you type when typing
- xterm is sending every keystroke to server, server don't need to buffer, nor add `\n`

## Ref

- https://blog.nelhage.com/2009/12/a-brief-introduction-to-termios/ describes pty and the echo back
- https://github.com/gorilla/websocket/blob/master/examples/command/main.go shows how to use non xterm, but it is wrong when using xterm
  - it adds `\n` because full command message is send in the form
  - it buffers output because it dump then message by message, which is not the case for xterm
- https://github.com/gravitational/console-demo is not using gorilla websocket but /x/net/websocket

Xterm

- when using cdn instead of package manager, you need to use global var `fit` and `attach` when apply plugin to `Terminal`