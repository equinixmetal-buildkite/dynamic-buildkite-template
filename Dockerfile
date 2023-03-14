FROM golangci/golangci-lint:latest
LABEL authors="Equinix Metal"
# using golangci-lint base image as it comes prepackaged with go 1.20
# it saves installation steps
# refer: https://hub.docker.com/layers/golangci/golangci-lint/latest/images/sha256-b0464d4d425f95f73a3443496c68be5fc3b5c456d2bd379f0d29ebaf7eeab657?context=explore

WORKDIR /go/src/workspace/dynamic-buildkite-template
COPY . .
RUN make all 

ENTRYPOINT [ "./dynamic-buildkite-template" ]