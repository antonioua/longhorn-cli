FROM golang:1.20 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd/ cmd/

RUN go mod download

COPY main.go main.go

ARG VERSION=undefined

# CGO_ENABLED=0 is needed to avoid building the possible C files since this part of the library may be missing in golang:alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build \
    -ldflags "-X main.Version=$VERSION" \
    -a \
    -o bin/longhorn-cli main.go

FROM alpine:latest

RUN apk update && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates

RUN addgroup --gid 901 appuser && \
    adduser -u 901 -G appuser -D -h /home/appuser appuser

WORKDIR /home/appuser

COPY --from=builder --chown=appuser:appuser /workspace/bin/longhorn-cli ./
USER appuser

ENTRYPOINT ["./longhorn-cli"]
