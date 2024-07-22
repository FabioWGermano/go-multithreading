package main

import (
	"log"
	"net/http"

	controller "github.com/FabioWGermano/go-multithreading/internal/controller/action"
)

func main() {
	http.HandleFunc("/buscar-cep", controller.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
