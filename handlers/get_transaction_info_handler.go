package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/gorilla/mux"
	"net/http"
	"testTask/controller"
	"testTask/error"
)

func GetTransactionInfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	txHash := mux.Vars(r)["txHash"]

	transactionInfo, err := controller.GetTransactionInfo(ctx, txHash)
	if errors.Is(err, ethereum.NotFound) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.IncorrectHash.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.ServerError.Error()))
		return
	}
	responce, err := json.Marshal(transactionInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.ServerError.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responce)
}
