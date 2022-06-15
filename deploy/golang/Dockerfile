# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18.2-buster AS build

ENV GO111MODULE=on \
	GOOS=linux \
	GOARCH=amd64

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
COPY ./app/controller /app/vendor/imports/dependencies/app/controller
COPY ./app/model /app/vendor/imports/dependencies/app/model
COPY ./config/.env /app/config/

RUN go install /app/app
COPY ./config/. /app/config/
RUN go build ./app/main.go
RUN echo $(ls -a /app/config)

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/main /main
COPY --from=build /app/config /config

EXPOSE 3200

USER nonroot:nonroot

ENTRYPOINT ["/main"]