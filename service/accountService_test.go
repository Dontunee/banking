package service

import (
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/dto"
	"github.com/Dontunee/banking/errs"
	moqDomain "github.com/Dontunee/banking/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func Test_CreateAccount_ShouldReturnValidationError_WhenRequestIsntValidated(testing *testing.T) {
	//Arrange
	accountRequest := dto.NewAccountRequest{
		"1010",
		"saving",
		0,
	}
	service := NewAccountService(nil)

	//Act
	_, appError := service.CreateAccount(accountRequest)

	//Assert
	if appError == nil {
		testing.Error("failed while testing the new account validation")
	}
}

var mockRepo *moqDomain.MockIAccountRepository
var service IAccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = moqDomain.NewMockIAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_CreateAccount_ShouldReturnServerSideError_IfNewAccountCannotBeCreated(testing *testing.T) {
	//Arrange
	tearDown := setup(testing)
	defer tearDown()

	accountRequest := dto.NewAccountRequest{
		"1010",
		"saving",
		6000,
	}
	account := domain.Account{
		"",
		accountRequest.CustomerID,
		time.Now().Format("2006-01-02 15:04:05"),
		accountRequest.AccountType,
		accountRequest.Amount,
		true,
	}
	mockRepo.EXPECT().Create(account).Return(nil, errs.NewUnexpectedError("unexpected database error "))

	//Act
	_, appError := service.CreateAccount(accountRequest)

	//Assert
	if appError == nil {
		testing.Error("Test failed while validating error for new account")
	}

}
