FROM golang:1.16-buster

RUN go version

ENV GOPATH=/ 

COPY ./ ./

RUN go mod download 
RUN go build -o shorter-api-client.exe ./cmd/client/main.go

CMD ["./shorter-api-client.exe"]