package hasher

// sha256, sha3, blake2
type Hasher interface {
	Hash(data []byte) ([]byte, error)
	Compare(data []byte, hash []byte) error
}

// hmac, pbkdf2, bcrypt
type SaltHasher interface {
	Hash(data []byte, salt []byte) ([]byte, error)
	Compare(data []byte, salt []byte, hash []byte) error
}

// aes
type CryptHasher interface {
	Encrypt(key, data []byte) ([]byte, error)
	Compare(data []byte, hash []byte) error
	Decrypt(data []byte, key []byte) error
}
