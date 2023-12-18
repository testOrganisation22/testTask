package handlers

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"net/http"
	"testTask/client"
	"testTask/error"
	"testTask/models"
)

func GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	ethereumClient := client.GetRPCClient()

	wallet := mux.Vars(r)["wallet"]
	if !common.IsHexAddress(wallet) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.IncorrectAddress.Error()))
		return
	}
	walletAddress := common.HexToAddress(wallet)
	balance, err := ethereumClient.BalanceAt(context.Background(), walletAddress, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.ServerError.Error()))
		return
	}

	walletWithBalance := models.Wallet{
		Wallet:  wallet,
		Balance: balance,
	}
	resp, err := json.Marshal(walletWithBalance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.ServerError.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
