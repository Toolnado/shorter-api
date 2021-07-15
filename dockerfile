FROM golang:1.16.5-buster

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o shorter-api ./cmd/server/main.go

CMD ["./shorter-api"]