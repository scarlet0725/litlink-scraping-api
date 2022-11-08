FROM golang:1.19.2-bullseye as builder

WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /go/src
RUN go build -o prism-api

FROM gcr.io/distroless/base-debian11

COPY --from=builder /go/src/prism-api /bin/prism-api

ENTRYPOINT [ "/bin/prism-api" ]