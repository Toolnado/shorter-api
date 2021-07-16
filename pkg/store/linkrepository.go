package store

import (
	"context"

	"github.com/Toolnado/shorter-api/pkg/model"
)

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) Create(l *model.Link) (*model.Link, error) {
	if err := r.store.db.QueryRow(context.Background(), "INSERT INTO links as m (long_url, short_url) values ($1, $2) ON CONFLICT(long_url) DO UPDATE SET short_url = $2 where m.long_url = $1 RETURNING id;",
		l.LongUrl,
		l.ShortUrl,
	).Scan(&l.ID); err != nil {
		return nil, err
	}
	return l, nil
}

func (r *LinkRepository) Get(l *model.Link) (*model.Link, error) {
	if err := r.store.db.QueryRow(context.Background(), "SELECT long_url FROM links WHERE short_url=$1", l.ShortUrl).Scan(&l.LongUrl); err != nil {
		return nil, err
	}
	return l, nil
}
