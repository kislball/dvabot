FROM golang:1.16
WORKDIR /dvamon/apps/gateway
COPY . .
RUN whoami
RUN go mod download && go build -ldflags "-s -w" -o ./gateway ./cmd/gateway/main.go
ENTRYPOINT ./gateway