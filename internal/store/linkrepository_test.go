package store_test

import (
	"testing"

	"github.com/Toolnado/shorter-api/internal/model"
	"github.com/Toolnado/shorter-api/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestLinkRepository_Create(t *testing.T) {
	store, teardown := store.TestStore(t)
	defer teardown("links")
	link, err := store.Link().Create(&model.Link{
		LongUrl:  "https://www.twitch.tv/",
		ShortUrl: "qwertyuiop",
	})
	assert.NoError(t, err)
	assert.NotNil(t, link)
}

func TestLinkRepository_Get(t *testing.T) {
	store, teardown := store.TestStore(t)
	shortUrl := "qwertyuiop"
	longUrl := "https://www.twitch.tv/"

	defer teardown("links")
	store.Link().Create(&model.Link{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	})

	link, err := store.Link().Get(&model.Link{
		ShortUrl: shortUrl,
	})

	assert.NoError(t, err)
	assert.NotNil(t, link)
	assert.Equal(t, longUrl, link.LongUrl)
}

func TestLinkRepository_GetShortUrl(t *testing.T) {
	shortUrl := "qwertyuiop"
	longUrl := "https://www.twitch.tv/"
	store, teardown := store.TestStore(t)
	defer teardown("links")
	store.Link().Create(&model.Link{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	})

	check, err := store.Link().CheckShortUrl(shortUrl)
	assert.NoError(t, err)
	assert.NotNil(t, check)
	assert.Equal(t, longUrl, check)
}

func TestLinkRepository_GenerateShortUrl(t *testing.T) {
	genShortUrl := store.NewStore().Link().GenerateShortUrl(make([]byte, 10))
	assert.NotNil(t, genShortUrl)
}
