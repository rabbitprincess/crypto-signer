package keystore

import (
	"github.com/cockroachdb/pebble"
	"github.com/rs/zerolog"
)

type DbStore struct {
	logger zerolog.Logger
	db     *pebble.DB
	root   string
}

func NewDbStore(root string, log zerolog.Logger) *DbStore {
	db, err := pebble.Open(root, &pebble.Options{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open db")
	}
	return &DbStore{
		logger: log,
		db:     db,
		root:   root,
	}
}

func (s *DbStore) Get(key string) (val []byte, err error) {
	val, closer, err := s.db.Get([]byte(key))
	defer closer.Close()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *DbStore) Set(key string, val []byte) error {
	return s.db.Set([]byte(key), val, pebble.Sync)
}

func (s *DbStore) Delete(key string) error {
	return s.db.Delete([]byte(key), pebble.Sync)
}
