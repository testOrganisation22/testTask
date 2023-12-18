package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"testTask/client"
	"testTask/handlers"
)

func main() {

	err := client.CreateRPCConnect()
	if err != nil {
		log.Fatal(err)
		return
	}
	r := mux.NewRouter().StrictSlash(true)

	handlers.SetupRoutes(r)

	serverAddr := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	log.Println("Starting server on port ", serverAddr)

	log.Fatal(http.ListenAndServe(serverAddr, r))
}
