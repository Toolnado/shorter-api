package store

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T) (*Store, func(...string)) {
	t.Helper()
	store := NewStore()
	db, err := store.Open()
	if err != nil {
		t.Fatal(err)
	}

	return store, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(context.Background(), fmt.Sprintf("TRUNCATE %s  CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		store.Close()
	}
}
