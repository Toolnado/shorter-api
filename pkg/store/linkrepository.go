package store

import "github.com/Toolnado/shorter-api/pkg/model"

type LinkRepository struct {
	store *Store
}

func (r *LinkRepository) Create(l *model.Link) (*model.Link, error) {
	if err := r.store.db.QueryRow("INSERT INTO links (long_url, short_url) VALUES ($1, $2) RETURNING id",
		l.LongUrl,
		l.ShortUrl,
	).Scan(&l.ID); err != nil {
		return nil, err
	}
	return l, nil
}
