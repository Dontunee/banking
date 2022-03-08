package dto

import (
	"net/http"
	"testing"
)

func Test_IsValid_ShouldReturnError_WhenTransactionTypeIsNotDepositOrWithdrawal(testing *testing.T) {
	//Arrange
	transactionRequest := TransactionRequest{TransactionType: "invalid transaction type"}

	//Act
	err, isValid := transactionRequest.Validate()

	//Assert
	if err.Message != "Transaction type can only be withdrawal or deposit" && isValid != false {
		testing.Error("invalid message while testing transaction type")
	}

	if err.Code != http.StatusUnprocessableEntity {
		testing.Error("invalid error code")
	}
}

func Test_IsValid_ShouldReturnError_WhenAmountIsLessThanZero(testing *testing.T) {
	//Arrange
	transactionRequest := TransactionRequest{
		Amount:          -100,
		TransactionType: "deposit",
	}

	//Act
	appError, isValid := transactionRequest.Validate()

	//Assert

	if appError.Message != "Amount can only be a positive number" && !isValid {
		testing.Error("invalid amount")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		testing.Error("invalid code while validating amount")
	}
}
