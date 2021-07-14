package main

import "google.golang.org/grpc"

func main() {
	s := grpc.NewServer()
	svr := linkShort.GRPCServer{}
}
