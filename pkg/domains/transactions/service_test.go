package transactions

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	service := NewService(nil, nil, nil)
	if !reflect.DeepEqual(*service, *new(Service)) {
		t.Error("Is not instance of service")
	}
}

type tCase struct {
	service *Service
	err     error
}

var tCases = map[string]tCase{
	"success": {
		NewService(
			new(TransactionRepositoryMock),
			&AccountsRepositoryMock{
				isBlocked:         1,
				blockAccountError: false,
			},
			new(TransactionObserverMock),
		),
		nil,
	},
	"isBlocked": {
		NewService(
			new(TransactionRepositoryMock),
			&AccountsRepositoryMock{
				isBlocked:         2,
				blockAccountError: false,
			},
			new(TransactionObserverMock),
		),
		nil,
	},
	"error on isBlocked": {
		NewService(
			new(TransactionRepositoryMock),
			&AccountsRepositoryMock{
				isBlocked:         0,
				blockAccountError: false,
			},
			new(TransactionObserverMock),
		),
		errorIsBlocked,
	},
	"error on BlockAccount": {
		NewService(
			new(TransactionRepositoryMock),
			&AccountsRepositoryMock{
				isBlocked:         1,
				blockAccountError: true,
			},
			new(TransactionObserverMock),
		),
		errorBlockAccount,
	},
	"error on CheckLimit": {
		NewService(
			&TransactionRepositoryMock{
				errorCheckLimit: true,
			},
			&AccountsRepositoryMock{
				isBlocked:         1,
				blockAccountError: false,
			},
			new(TransactionObserverMock),
		),
		errorCheckLimit,
	},
	"error on CreateTransaction": {
		NewService(
			&TransactionRepositoryMock{
				errorCreateTransaction: true,
			},
			&AccountsRepositoryMock{
				isBlocked:         1,
				blockAccountError: false,
			},
			new(TransactionObserverMock),
		),
		errorCreateTransaction,
	},
}

func TestService_SaveTransaction(t *testing.T) {
	for name, payload := range tCases {
		err := payload.service.SaveTransaction(new(TransactionMock))
		if err != nil && err.Error() != payload.err.Error() {
			t.Errorf("Invalid return in case %s\n returned: %+v", name, err)
		}
	}
}
