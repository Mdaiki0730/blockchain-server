package model

import (
	"fmt"
	"math/big"

	"garicoin/pkg/converter"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func NewSignature(signature string) (*Signature, error) {
	x, y, err := converter.String2BigIntTuple(signature)
	if err != nil {
		return nil, err
	}
	return &Signature{&x, &y}, nil
}

func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}
