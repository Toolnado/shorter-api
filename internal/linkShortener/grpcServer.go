package linkShortener

import (
	"context"
	"log"

	"github.com/Toolnado/shorter-api/internal/api"
	"github.com/Toolnado/shorter-api/internal/model"
	"github.com/Toolnado/shorter-api/internal/store"
)

type GRPCServer struct {
	api.UnimplementedLinkShortenerServer
}

func (s *GRPCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	store := store.NewStore()
	shortUrl := make([]byte, 10)

	store.Open()
	defer store.Close()

	store.Link().GenerateShortUrl(shortUrl)

	check, err := store.Link().CheckShortUrl(string(shortUrl))
	if err != nil {
		log.Println(err)
	}
	for check != "" {
		log.Println("The shortened link already exists\n", string(shortUrl), check)
		shortUrl = store.Link().GenerateShortUrl(shortUrl)
		check, err = store.Link().CheckShortUrl(string(shortUrl))
		if err != nil {
			log.Println(err)
		}
	}

	newLink := model.Link{
		LongUrl:  req.GetUrl(),
		ShortUrl: string(shortUrl),
	}

	store.Link().Create(&newLink)

	return &api.CreateResponse{ShortUrl: string(shortUrl)}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	store := store.NewStore()
	store.Open()
	defer store.Close()
	newLink := model.Link{
		LongUrl:  "",
		ShortUrl: req.GetShortUrl(),
	}
	store.Link().Get(&newLink)
	return &api.GetResponse{Url: newLink.LongUrl}, nil
}
