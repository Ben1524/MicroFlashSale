FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/user_srv

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY . .
RUN make proto tidy build

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/user_srv/user_srv /user_srv
ENTRYPOINT ["/user_srv"]
CMD []
