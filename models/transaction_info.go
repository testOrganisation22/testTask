package models

import "math/big"

type TransactionInfo struct {
	Hash              string
	Sender            string
	Receiver          string
	SumCost           *big.Int
	Value             *big.Int
	Nonce             uint64
	Gas               uint64
	GasPrice          string
	TransactionStatus string
}
