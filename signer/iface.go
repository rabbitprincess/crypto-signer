package signer

type Signer interface {
	Generate() (privkey []byte, pubkey []byte, err error)
	Pubkey(privKey []byte) (pubKey []byte, err error)
	Sign(msg []byte, privkey []byte) (signature []byte, err error)
	Verify(signature []byte, pubkey []byte) error
}

type SignAggregator interface {
	Signer
	Aggregate(sigs ...[]byte) ([]byte, error)
}
