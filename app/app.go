package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	//define the route and handler which provides response and brings request from client
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/createCustomer", createCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8902", router)) // starts the server on asn ip and port and use default multiplexer
}
