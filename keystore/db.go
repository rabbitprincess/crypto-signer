package keystore

import (
	"github.com/cockroachdb/pebble"
	"github.com/rs/zerolog"
)

var _ KeyStore = (*DbStore)(nil)

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

func (s *DbStore) Get(path, key string) (val []byte, err error) {
	val, closer, err := s.db.Get(s.makePathKey(path, key))
	defer closer.Close()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (s *DbStore) Set(path, key string, val []byte) error {
	return s.db.Set(s.makePathKey(path, key), val, pebble.Sync)
}

func (s *DbStore) Delete(path, key string) error {
	return s.db.Delete(s.makePathKey(path, key), pebble.Sync)
}

func (s *DbStore) Close() error {
	return s.db.Close()
}

func (s *DbStore) makePathKey(path, key string) []byte {
	return []byte(path + "/" + key)
}
