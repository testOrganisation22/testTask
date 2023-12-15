package models

import "math/big"

type Wallet struct {
	Wallet  string
	Balance *big.Int
}
