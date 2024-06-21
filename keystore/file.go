package keystore

import (
	"github.com/rs/zerolog"
	"github.com/spf13/afero"
)

var _ KeyStore = (*FileStore)(nil)

type FileStore struct {
	logger zerolog.Logger
	osfs   afero.Fs
	root   string
}

func NewFileStore(root string, log zerolog.Logger) *FileStore {
	return &FileStore{
		logger: log,
		osfs:   afero.NewOsFs(),
		root:   root,
	}
}

func (s *FileStore) Get(key string) (val []byte, err error) {
	return afero.ReadFile(afero.NewOsFs(), s.path(key))
}

func (s *FileStore) Set(key string, val []byte) error {
	return afero.WriteFile(afero.NewOsFs(), s.path(key), val, 0644)
}

func (s *FileStore) Delete(key string) error {
	return s.osfs.RemoveAll(s.path(key))
}

func (s *FileStore) path(key string) string {
	return s.root + "/" + key
}
