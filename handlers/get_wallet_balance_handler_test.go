package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testTask/client"
	"testing"
)

func TestGetWalletBalance(t *testing.T) {
	os.Setenv("RPC_URL", "https://rpc.ankr.com/eth")
	route := "/wallet/balance/"
	err := client.CreateRPCConnect()
	if err != nil {
		log.Fatal(err)
		return
	}

	r := mux.NewRouter()
	SetupRoutes(r)
	t.Run("Correct address test", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route+"0x6bD1A64227866Ac25B7F9c918B3Dc32D2a9CFA19", nil)
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
