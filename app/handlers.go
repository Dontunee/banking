package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Dontunee/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

//Create customer structure and marshal to json

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.ICustomerService
}

func (customerHandler *CustomerHandlers) getAllCustomers(w http.ResponseWriter, request *http.Request) {
	customers, _ := customerHandler.service.GetAllCustomer()
	if request.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func (customerHandler *CustomerHandlers) getCustomerById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]

	customer, err := customerHandler.service.GetCustomerById(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Println(writer, err.Error())
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customer)
	}

}
