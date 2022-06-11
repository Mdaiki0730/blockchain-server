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

func NewSignature(signature string) *Signature {
  x, y := converter.String2BigIntTuple(signature)
  return &Signature{&x, &y}
}

func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}
