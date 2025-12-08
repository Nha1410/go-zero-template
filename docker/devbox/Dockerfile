FROM golang:1.24-alpine

RUN apk add --no-cache \
    git \
    make \
    bash \
    curl

# Install goctl
RUN go install github.com/zeromicro/go-zero/tools/goctl@latest

# Install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install migration tools (optional - uncomment if needed)
# RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /go/src

CMD ["/bin/sh"]

