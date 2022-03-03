package app

import (
	"encoding/json"
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionHandlers struct {
	service service.TransactionService
}

func (handlers TransactionHandlers) createTransaction(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	accountId := vars["account_id"]
	transactionRequest := dto.TransactionRequest{}
	err := json.NewDecoder(request.Body).Decode(&transactionRequest)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error)
	}
	transactionRequest.AccountId = accountId

	transaction, appError := handlers.service.ProcessTransaction(transactionRequest)
	if appError != nil {
		writeResponse(writer, appError.Code, appError.Message)
	} else {
		writeResponse(writer, http.StatusCreated, transaction)
	}
}
