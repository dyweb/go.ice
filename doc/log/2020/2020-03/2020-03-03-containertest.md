# 2020-03-03 Container Test

For https://github.com/dyweb/go.ice/issues/56, it is required by benchhub for testing relational database.

## TODO

- [ ] revive the [old docker client](https://github.com/dyweb/go.ice/tree/archive/2020-01-13/lib/dockerclient)
- [ ] allow start/stop mysql container and wait for ready
- [ ] create database from testdata

It's a bit hard to have go mod working with docker, suggested way is https://github.com/moby/moby/issues/39302#issuecomment-504146736

Moved part of dockerclient part, but maybe shell out is a better idea, compared with manually vendoring the part I need ...