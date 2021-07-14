package api

import (
	"context"
)

type GRPCServer struct{}

func (s *GRPCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {

}
