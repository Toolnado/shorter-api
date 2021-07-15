package main

import (
	"log"
	"net"
	"os"

	"github.com/Toolnado/shorter-api/pkg/api"
	"github.com/Toolnado/shorter-api/pkg/linkShortener"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("port not found")
	}

	s := grpc.NewServer()
	srv := &linkShortener.GRPCServer{}
	api.RegisterLinkShortenerServer(s, srv)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
