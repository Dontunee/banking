package app

import (
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	customerHandler := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define the route and handler which provides response and brings request from client

	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomerById).Methods(http.MethodGet)

	// starts the server on asn ip and port and use default multiplexer

	log.Fatal(http.ListenAndServe("localhost:8902", router))
}
