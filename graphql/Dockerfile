FROM golang:1.16.5 as builder
WORKDIR /app
COPY .  ./
RUN go mod download
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/server \
    -ldflags '-s -w'

FROM scratch as runner
COPY --from=builder /go/bin/server /app/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]