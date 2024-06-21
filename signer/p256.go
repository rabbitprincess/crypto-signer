package signer

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"
)

// P256Signer implements cryptographic operations for the NIST P-256 (secp256r1) elliptic curve.
type P256Signer struct{}

// NewP256Signer creates a new P256Signer instance.
func NewP256Signer() *P256Signer {
	return &P256Signer{}
}

// Generate creates a new ECDSA P-256 key pair.
func (p *P256Signer) Generate() (privkey []byte, pubkey []byte, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	privkey = privateKey.D.Bytes()
	// get compressed public key
	pubkey = elliptic.MarshalCompressed(elliptic.P256(), privateKey.X, privateKey.Y)
	return privkey, pubkey, nil
}

func (p *P256Signer) Pubkey(privkey []byte) (pubkey []byte, err error) {
	d := new(big.Int).SetBytes(privkey)
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult(d.Bytes())
	pubkey = elliptic.MarshalCompressed(elliptic.P256(), x, y)

	return pubkey, nil
}

func (p *P256Signer) Sign(msg []byte, privkey []byte) (signature []byte, err error) {
	d := new(big.Int).SetBytes(privkey)

	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
		},
		D: d,
	}

	hash := crypto.SHA256.New()
	hash.Write(msg)
	digest := hash.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, digest)
	if err != nil {
		return nil, err
	}

	return append(r.Bytes(), s.Bytes()...), nil
}

func (p *P256Signer) Verify(msg []byte, signature []byte, pubkey []byte) (bool, error) {
	x, y := elliptic.Unmarshal(elliptic.P256(), pubkey)
	if x == nil {
		return false, errors.New("invalid public key")
	}

	publicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	hash := crypto.SHA256.New()
	hash.Write(msg)
	digest := hash.Sum(nil)

	r := big.Int{}
	s := big.Int{}
	r.SetBytes(signature[:len(signature)/2])
	s.SetBytes(signature[len(signature)/2:])

	return ecdsa.Verify(publicKey, digest, &r, &s), nil
}
