FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.26-alpine AS builder

ARG TARGETOS=linux
ARG TARGETARCH=amd64

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -trimpath -o /chromagoth ./cmd/chromagoth

FROM alpine:latest
COPY --from=builder /chromagoth /usr/local/bin/chromagoth
ENTRYPOINT ["/chromagoth"]
