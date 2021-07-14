package shorsServer

import (
	"context"
	"math/rand"

	"github.com/Toolnado/shorter-api/pkg/api"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type GRPCServer struct{}

func (s *GRPCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	short := make([]byte, 5)
	for i := range short {
		short[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return &api.CreateResponse{ShortUrl: string(short)}, nil
}
