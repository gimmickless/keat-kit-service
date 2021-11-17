# syntax=docker/dockerfile:1
FROM golang:1.17-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /app ./cmd/app

FROM alpine
WORKDIR /
RUN apk add --no-cache tini tzdata
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/sbin/tini", "--"]
CMD /app