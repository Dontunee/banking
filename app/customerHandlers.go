package app

import (
	"encoding/json"
	"github.com/Dontunee/banking/errs"
	"github.com/Dontunee/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.ICustomerService
}

func (customerHandler *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	keys, ok := request.URL.Query()["status"]
	var customers interface{}
	var err *errs.AppError

	if !ok {
		customers, err = customerHandler.service.GetAllCustomer()
	} else {
		status := keys[0]
		customers, err = customerHandler.getCustomersByStatus(status)
	}

	if err != nil {
		writeResponse(writer, http.StatusInternalServerError, err.ReturnMessage())
	}
	writeResponse(writer, http.StatusOK, customers)

}

func (customerHandler *CustomerHandlers) getCustomerById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]

	customer, err := customerHandler.service.GetCustomerById(id)
	if err != nil {
		writeResponse(writer, err.Code, err.ReturnMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}

}
func (customerHandler *CustomerHandlers) getCustomersByStatus(status string) (interface{}, *errs.AppError) {
	inputStatus := false
	if status == "active" {
		inputStatus = true
	}
	customers, err := customerHandler.service.GetCustomersByStatus(inputStatus)
	return customers, err
}

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		panic(err)
	}

}
