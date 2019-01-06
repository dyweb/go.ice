# 2019-01-05 Docker exec log

## Exec

- exec is hjacking http request's tcp connection
- only client side, need upgrade and get the raw tcp connection
- to make it interactive, just copy stdout from tcp connection and copy stdin to tcp connection

## Logs

- docker multiplex stream when write to json
- docker has timestamp for each line (it is splitting them in go code) in `Copier` but it's in UTC
  - https://github.com/moby/moby/blob/master/daemon/logger/copier.go#L118
- use `pkg/stdcopy.StdCopy` to deduplex to stdout and stderr
- the log file is by default stored as a json file in the format `{log, stream, time}`