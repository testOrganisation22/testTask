package controller

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"testTask/constants"
	"testTask/models"
)

func GetTransactionInfo(ctx context.Context, txHash string) (*models.TransactionInfo, error) {
	var (
		transactionHash = common.HexToHash(txHash)
	)
	//TODO: refactor it
	ethereumClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	transaction, _, err := ethereumClient.TransactionByHash(ctx, transactionHash)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sender, err := types.Sender(types.NewLondonSigner(transaction.ChainId()), transaction)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	transactionStatus := constants.TransactionPending
	if transaction.To() != nil {
		transactionStatus = constants.TransactionSuccess
	}
	return &models.TransactionInfo{
		Hash:              transaction.Hash().Hex(),
		Sender:            sender.Hex(),
		Receiver:          transaction.To().Hex(),
		SumCost:           transaction.Cost(),
		Value:             transaction.Value(),
		GasPrice:          transaction.GasPrice().String(),
		Gas:               transaction.Gas(),
		Nonce:             transaction.Nonce(),
		TransactionStatus: transactionStatus,
	}, nil
}
