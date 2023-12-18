package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"testTask/client"
	"testTask/models"
	"testing"
)

func TestGetTransactionInfoHandler(t *testing.T) {
	os.Setenv("RPC_URL", "https://rpc.ankr.com/eth")
	route := "/transaction/"
	transaction := "0xa4ff4bff034b1a005c0ab14bac73c84660d75738f86a45ce67a3e690266639ce"

	err := client.CreateRPCConnect()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := mux.NewRouter()
	SetupRoutes(r)

	t.Run("Correct address response test", func(t *testing.T) {
		hashInfo := &models.TransactionInfo{
			Hash:              "0xa4ff4bff034b1a005c0ab14bac73c84660d75738f86a45ce67a3e690266639ce",
			Sender:            "0x6bD1A64227866Ac25B7F9c918B3Dc32D2a9CFA19",
			Receiver:          "0xf466f27fB811Ab1572CA67ab438E966910f5d9C1",
			SumCost:           big.NewInt(25000000000000000),
			Value:             big.NewInt(0),
			Nonce:             106074,
			Gas:               500000,
			GasPrice:          "50000000000",
			TransactionStatus: "Success",
		}
		hashInfoExpected, err := json.Marshal(hashInfo)
		if err != nil {
			t.Error(err)
			return
		}
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route+transaction, nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(rr, req)
		if bytes.Compare(rr.Body.Bytes(), hashInfoExpected) != 0 {
			t.Errorf("Expected response '%s', got '%s'", hashInfoExpected, rr.Body.String())
		}
	})
	t.Run("Incorrenct address response test", func(t *testing.T) {
		hashInfo := &models.TransactionInfo{
			Hash:              "0xa4ff4bff034b1a005c0ab14bac73c84660d75738f86a45ce67a3e690266639ce",
			Sender:            "0x6bD1A64227866Ac25B7F9c918B3Dc32D2a9CFA19",
			Receiver:          "0xf466f27fB811Ab1572CA67ab438E966910f5d9C1",
			SumCost:           big.NewInt(25000000000000000),
			Value:             big.NewInt(0),
			Nonce:             10607,
			Gas:               50000,
			GasPrice:          "50000000000",
			TransactionStatus: "Success",
		}
		hashInfoExpected, err := json.Marshal(hashInfo)
		if err != nil {
			t.Error(err)
			return
		}
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route+transaction, nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(rr, req)
		if bytes.Compare(rr.Body.Bytes(), hashInfoExpected) == 0 {
			t.Errorf("Expected response '%s', got '%s'", hashInfoExpected, rr.Body.String())
		}
	})
	t.Run("Correct address test", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route+transaction, nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
	t.Run("Incorrect address test", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route+"877", nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
	t.Run("Empty address test", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route, nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}
