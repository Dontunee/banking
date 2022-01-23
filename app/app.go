package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()
	//define the route and handler which provides response and brings request from client
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8902", mux)) // starts the server on asn ip and port and use default multiplexer
}
