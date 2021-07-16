package linkShortener

import (
	"context"
	"math/rand"

	"github.com/Toolnado/shorter-api/pkg/api"
	"github.com/Toolnado/shorter-api/pkg/model"
	"github.com/Toolnado/shorter-api/pkg/store"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type GRPCServer struct {
	api.UnimplementedLinkShortenerServer
}

func (s *GRPCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	store := store.Store{}
	store.Open()
	defer store.Close()

	shortUrl := make([]byte, 10)
	for i := range shortUrl {
		shortUrl[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	newLink := model.Link{
		LongUrl:  req.GetUrl(),
		ShortUrl: string(shortUrl),
	}

	store.Link().Create(&newLink)

	return &api.CreateResponse{ShortUrl: string(shortUrl)}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	store := store.Store{}
	store.Open()
	defer store.Close()
	newLink := model.Link{
		LongUrl:  "",
		ShortUrl: req.GetShortUrl(),
	}
	store.Link().Get(&newLink)
	return &api.GetResponse{Url: newLink.LongUrl}, nil
}
