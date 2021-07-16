FROM golang:1.16-buster

RUN go version

ENV GOPATH=/ 

COPY ./ ./

RUN go mod download 
RUN go build -o shorter-api-server.exe ./cmd/server/main.go 

CMD ["./shorter-api-server.exe"]