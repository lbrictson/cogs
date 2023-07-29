# Build in app in a Go container
FROM docker.io/golang:1.20.6-bullseye as builder
COPY . /app
WORKDIR /app
RUN go env -w GOPROXY=direct && go env -w GOSUMDB=off
RUN CGO_ENABLED=0 go build -o main cmd/server/main.go
# Move artifact to smaller container with no Go tools installed
FROM docker.io/alpine:3.15.0
RUN apk update && apk upgrade && apk --no-cache add ca-certificates bash curl wget
WORKDIR /app
COPY --from=builder /app/main app
ENTRYPOINT ["/app/app"]