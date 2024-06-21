package keystore

type KeyStore interface {
	// GetKey returns the key for the given keyID
	Get(path, key string) (val []byte, err error)
	Set(path, key string, val []byte) error
	Delete(path, key string) error
}
