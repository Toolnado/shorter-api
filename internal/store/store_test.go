package store_test

import (
	"log"
	"os"
	"testing"

	"github.com/Toolnado/shorter-api/internal/store"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	os.Exit(m.Run())
}

func TestStore_Open(t *testing.T) {
	store := store.NewStore()
	db, err := store.Open()
	defer store.Close()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func Test_NewStore(t *testing.T) {
	store := store.NewStore()
	assert.NotNil(t, store)
}

func Test_StoreLink(t *testing.T) {
	linkRepository := store.NewStore().Link()
	assert.NotNil(t, linkRepository)
}
