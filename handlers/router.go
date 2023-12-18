package handlers

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/healthcheck", GetHealthcheckHandler).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{txHash}", GetTransactionInfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/wallet/balance/{wallet}", GetWalletBalance).Methods(http.MethodGet)
	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

}
