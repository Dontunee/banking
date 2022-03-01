package app

import (
	"encoding/json"
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandlers struct {
	service service.AccountService
}

func (handler AccountHandlers) createAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customer_id"]
	if customerId == "" {
		writeResponse(writer, http.StatusBadRequest, error.Error)
	}
	var newAccount dto.NewAccountRequest
	err := json.NewDecoder(request.Body).Decode(&newAccount)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		newAccount.CustomerID = customerId
		account, appError := handler.service.CreateAccount(newAccount)
		if appError != nil {
			writeResponse(writer, appError.Code, appError.Message)
		} else {
			writeResponse(writer, http.StatusCreated, account)
		}
	}
}
