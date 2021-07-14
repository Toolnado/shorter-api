package main

import (
	"github.com/Toolnado/shorter-api/pkg/api"
	"github.com/Toolnado/shorter-api/pkg/linkShortener"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &linkShortener.GRPCServer{}
	api.RegisterLinkShortenerServer(s, srv)
}
