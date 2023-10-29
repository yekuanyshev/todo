# syntax=docker/dockerfile:1

## Build
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o /service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /service /service

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["./service"]