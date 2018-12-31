# NOTE: copied from dyweb/gommon
FROM dyweb/go-dev:1.11.4 as builder

LABEL maintainer="contact@dongyue.io"

ARG PROJECT_ROOT=/go/src/github.com/dyweb/go.ice/

WORKDIR $PROJECT_ROOT

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers will only be re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml $PROJECT_ROOT
RUN dep ensure -v -vendor-only

# Copy all project and build it
COPY . $PROJECT_ROOT
RUN make install

# NOTE: use ubuntu instead of alphine
#
# When using alpine I saw standard_init_linux.go:190: exec user process caused "no such file or directory",
# because I didn't compile go with static flag
# https://stackoverflow.com/questions/49535379/binary-compiled-in-prestage-doesnt-work-in-scratch-container
FROM ubuntu:18.04 as runner
LABEL maintainer="contact@dongyue.io"
LABEL github="github.com/dyweb/go.ice"
WORKDIR /usr/bin
COPY --from=builder /go/bin/dk .
ENTRYPOINT ["dk"]
CMD ["help"]