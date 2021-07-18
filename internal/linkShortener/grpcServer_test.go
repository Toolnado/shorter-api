package linkShortener_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/Toolnado/shorter-api/internal/api"
	"github.com/Toolnado/shorter-api/internal/linkShortener"
	"github.com/Toolnado/shorter-api/internal/model"
	"github.com/Toolnado/shorter-api/internal/store"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
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

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	os.Exit(m.Run())
}

func TestGrpcServer_Create(t *testing.T) {
	url := "https://www.twitch.tv/"

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("port not found")
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewLinkShortenerClient(conn)
	res, err := client.Create(context.Background(), &api.CreateRequest{Url: url})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGrpcServer_Get(t *testing.T) {
	store, teardown := store.TestStore(t)
	shortUrl := "qwertyuiop"
	longUrl := "https://www.twitch.tv/"

	defer teardown("links")

	store.Link().Create(&model.Link{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	})

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("port not found")
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewLinkShortenerClient(conn)
	res, err := client.Get(context.Background(), &api.GetRequest{ShortUrl: shortUrl})
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, longUrl, res.Url)
}
