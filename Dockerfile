############################
# STEP 1 build base image from buster
############################
FROM golang:1.16.4-buster AS build

WORKDIR /go/src/app

# copy mod files and download them for caching benefits
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

# copy the app code
COPY . /go/src/app

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s' \
        -o /go/bin/mk-srv github.com/omaralsoudanii/mk-srv/cmd/mk-srv

############################
# STEP 2 build distroless image
############################
FROM gcr.io/distroless/static

# copy the binary , and config file (due to pathing problem when loading the config)
COPY --from=build /go/bin/mk-srv ./
COPY --from=build /go/src/app/config.yaml ./
# Run the server binary.
CMD ["./mk-srv"]
