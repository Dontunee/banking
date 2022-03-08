package app

import (
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
	"github.com/Dontunee/banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var customerHandler CustomerHandlers
var mockService *service.MockICustomerService

func setup(testing *testing.T) func() {
	controller := gomock.NewController(testing)
	mockService = service.NewMockICustomerService(controller)
	customerHandler = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getAllCustomers)
	return func() {
		router = nil
		defer controller.Finish()
	}
}

func Test_GetAllCustomers_ShouldReturnStatusCode200AndCustomers(testing *testing.T) {
	//Arrange
	teardown := setup(testing)
	defer teardown()
	dummyCustomers := []dto.CustomerResponse{
		{1001, "Tunde", "lagos", "10001", "1996-04-03", "true"},
		{1002, "Femi", "lagos", "1005", "2000-04-03", "false"},
	}
	mockService.EXPECT().GetAllCustomer().Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	recorder := httptest.NewRecorder()

	//Act
	router.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK {
		testing.Error("failed while testing the status code")
	}
}

func Test_GetAllCustomers_ShouldReturnStatusCode500WithErrorMessage_WhenErrorOccurs(testing *testing.T) {
	//Arrange
	teardown := setup(testing)
	defer teardown()
	mockService.EXPECT().GetAllCustomer().Return(nil, errs.NewUnexpectedError("database error occurred"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	recorder := httptest.NewRecorder()

	//Act
	router.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		testing.Error("failed while testing the error status code")
	}
}
