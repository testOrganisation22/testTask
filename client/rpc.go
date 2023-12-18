package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

var RPCClient *ethclient.Client

func GetRPCClient() *ethclient.Client {
	return RPCClient
}

func CreateRPCConnect() error {
	var err error
	RPCClient, err = ethclient.Dial(os.Getenv("RPC_URL"))
	return err
}
