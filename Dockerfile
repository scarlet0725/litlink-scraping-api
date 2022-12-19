FROM golang:1.19.4-bullseye as builder

WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /go/src
RUN go build -o prism-api

FROM gcr.io/distroless/base-debian11:latest

COPY --from=builder --chown=nonroot:nonroot /go/src/prism-api /bin/prism-api

USER nonroot
ENTRYPOINT [ "/bin/prism-api" ]