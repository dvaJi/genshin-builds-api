##
## Build
##
FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env .
COPY commands ./commands
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /genshindata-api

COPY . .

##
## Deploy
##
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=build /genshindata-api /genshindata-api
COPY .env /genshindata-api

ENTRYPOINT ["/genshindata-api"]