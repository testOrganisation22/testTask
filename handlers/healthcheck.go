package handlers

import "net/http"

func GetHealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server works correctly"))
}
