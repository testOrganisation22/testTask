package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"testTask/handlers"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/healthcheck", handlers.GetHealthcheckHandler).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{txHash}", handlers.GetTransactionInfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/wallet/balance/{wallet}", handlers.GetWalletBalance).Methods(http.MethodGet)

	serverAddr := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	log.Println("Starting server on port ", serverAddr)

	log.Fatal(http.ListenAndServe(serverAddr, r))
}
