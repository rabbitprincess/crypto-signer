package keystore

type KeyStore interface {
	// GetKey returns the key for the given keyID
	Get(key string) (val []byte, err error)
	Set(key string, val []byte) error
	Delete(key string) error
}
