package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

//Create customer structure and marshal to jsaon
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Hello world") // io writer and response to write
}

func getAllCustomers(w http.ResponseWriter, request *http.Request) {
	//create list of customers
	customers := []Customer{
		{"Tunde", "lagos", "1000"},
		{"Femi", "Oyebanji", "1453423"},
	}
	if request.Header.Get(("Content-Type")) == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
