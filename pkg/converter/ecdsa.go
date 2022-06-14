package converter

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
)

func String2BigIntTuple(s string) (big.Int, big.Int, error) {
	bx, err := hex.DecodeString(s[:64])
	if err != nil {
		return big.Int{}, big.Int{}, err
	}
	by, err := hex.DecodeString(s[64:])
	if err != nil {
		return big.Int{}, big.Int{}, err
	}
	var bix big.Int
	var biy big.Int

	_ = bix.SetBytes(bx)
	_ = biy.SetBytes(by)
	return bix, biy, nil
}

func PublicKeyFromString(s string) (*ecdsa.PublicKey, error) {
	x, y, err := String2BigIntTuple(s)
	if err != nil {
		return nil, err
	}
	return &ecdsa.PublicKey{elliptic.P256(), &x, &y}, nil
}

func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) (*ecdsa.PrivateKey, error) {
	b, err := hex.DecodeString(s[:])
	if err != nil {
		return nil, err
	}
	var bi big.Int
	_ = bi.SetBytes(b)
	return &ecdsa.PrivateKey{*publicKey, &bi}, nil
}
