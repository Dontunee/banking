package app

import (
	"log"
	"net/http"
)

func Start() {

	//define the route and handler which provides response and brings request from client
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8902", nil)) // starts the server on asn ip and port and use default multiplexer
}
