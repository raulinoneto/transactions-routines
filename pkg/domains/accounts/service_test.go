package accounts

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	service := NewService(nil)
	if !reflect.DeepEqual(*service, *new(Service)) {
		t.Error("Is not instance of service")
	}
}

type AccountMock struct {
	documentNumber int
}

func (a AccountMock) GetDocumentNumber() int { return a.documentNumber }
func (AccountMock) GetID() int               { return 1 }
func (AccountMock) SetID(_ int)              {}
func (AccountMock) SetDocumentNumber(_ int)  {}

type RepositoryMock struct {
	errorCreateAccount, errorGetAccount bool
}

var errorCreateAccount = errors.New("generic errorCreateAccount")
var errorGetAccount = errors.New("generic errorGetAccount")

func (r *RepositoryMock) CreateAccount(_ Account) error {
	if r.errorCreateAccount {
		return errorCreateAccount
	}
	return nil
}
func (r *RepositoryMock) GetAccount(id int) (Account, error) {
	if r.errorGetAccount {
		return nil, errorGetAccount
	}
	result := AccountMock{id}
	return result, nil
}

type tCaseSaveAccount struct {
	service *Service
	payload *AccountMock
	err     error
}

var tCasesSave = map[string]tCaseSaveAccount{
	"success": {
		NewService(
			&RepositoryMock{},
		),
		&AccountMock{1},
		nil,
	},
	"error on CreateAccount": {
		NewService(
			&RepositoryMock{},
		),
		&AccountMock{1},
		errorCreateAccount,
	},
	"error DocumentInvalid": {
		NewService(
			&RepositoryMock{},
		),
		&AccountMock{0},
		errorDocumentInvalid,
	},
}

func TestService_SaveAccount(t *testing.T) {
	for name, tCase := range tCasesSave {
		err := tCase.service.SaveAccount(tCase.payload)
		if err != nil && err.Error() != tCase.err.Error() {
			t.Errorf("Invalid return in case %s\n returned: %+v", name, err)
		}
	}

}

type tCaseGetAccount struct {
	service *Service
	payload int
	err     error
}

var tCasesGet = map[string]tCaseGetAccount{
	"success": {
		NewService(
			&RepositoryMock{},
		),
		1,
		nil,
	},
	"error on GetAccount": {
		NewService(
			&RepositoryMock{
				errorGetAccount: true,
			},
		),
		1,
		errorGetAccount,
	},
}

func TestService_GetAccount(t *testing.T) {
	for name, tCase := range tCasesGet {
		expected := AccountMock{tCase.payload}
		result, err := tCase.service.GetAccount(tCase.payload)
		if err != nil && err.Error() != tCase.err.Error() {
			t.Errorf("Invalid error return in case %s\n returned: %+v", name, err)
		}
		if result != nil && expected.documentNumber != result.GetDocumentNumber() {
			t.Errorf("Invalid return in case %s\n returned: %+v", name, result)
		}
	}
}
