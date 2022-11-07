FROM golang:1.19.2-bullseye as builder

COPY . /go/src
WORKDIR /go/src
RUN go mod tidy
RUN go build -o prism-api

FROM gcr.io/distroless/base-debian11

COPY --from=builder /go/src/prism-api /bin/prism-api

ENTRYPOINT [ "/bin/prism-api" ]