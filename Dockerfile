FROM golang:1.19.2-bullseye as builder

WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /go/src
RUN go build -o prism-api

FROM golang:1.19.2-bullseye as healthcheck-builder

WORKDIR /go/src
COPY go.mod .
COPY go.sum .
COPY cmd/healthcheck/healthcheck.go /go/src/healthcheck.go
RUN go build -o healthcheck
RUN chmod 755 healthcheck

FROM gcr.io/distroless/base-debian11

COPY --from=builder --chown=nonroot:nonroot /go/src/prism-api /bin/prism-api
COPY --from=healthcheck-builder --chown=nonroot:nonroot /go/src/healthcheck /bin/healthcheck

USER nonroot
ENTRYPOINT [ "/bin/prism-api" ]