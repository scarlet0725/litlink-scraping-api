ARG GO_VERSION=1.19.4

FROM golang:${GO_VERSION}-bullseye as builder

WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /go/src
RUN go build -o prism-api

FROM gcr.io/distroless/base-debian11:nonroot

COPY --from=builder --chown=nonroot:nonroot /go/src/prism-api /bin/prism-api

ENTRYPOINT [ "/bin/prism-api" ]
