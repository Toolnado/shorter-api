package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Toolnado/shorter-api/internal/api"
	"github.com/Toolnado/shorter-api/internal/linkShortener"
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

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	s := grpc.NewServer()
	srv := &linkShortener.GRPCServer{}
	api.RegisterLinkShortenerServer(s, srv)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	<-stopChan

	log.Println("Server shutting dawn")
	s.GracefulStop()
	log.Println("Server stoped")
}
