package transactions

import (
	"errors"
	"time"
)

type TransactionRepositoryMock struct {
	errorCheckLimit        bool
	errorCreateTransaction bool
}

var errorCheckLimit = errors.New("generic errorCheckLimit")
var errorCreateTransaction = errors.New("generic errorCreateTransaction")

func (t *TransactionRepositoryMock) CreateTransaction(Transaction) error {
	if t.errorCheckLimit {
		return errorCheckLimit
	}
	return nil
}

func (t *TransactionRepositoryMock) CheckLimit(_ int, _ float64) error {
	if t.errorCreateTransaction {
		return errorCreateTransaction
	}
	return nil
}

type AccountsRepositoryMock struct {
	// value 1: return not blocked
	// value 2: return blocked
	// otherwise: return error
	isBlocked int

	blockAccountError bool
}

var errorIsBlocked = errors.New("generic errorIsBlocked")
var errorBlockAccount = errors.New("generic errorBlockAccount")

func (a *AccountsRepositoryMock) BlockAccount(int) error {
	if a.blockAccountError {
		return errorBlockAccount
	}
	return nil
}

func (AccountsRepositoryMock) UnlockAccount(int) {}
func (AccountsRepositoryMock) ChangeLimit(amount float64, id int) error { return nil}
func (AccountsRepositoryMock) CheckLimit(accountID int, value float64) error { return nil}

func (a *AccountsRepositoryMock) AccountIsBlocked(int) (bool, error) {
	switch a.isBlocked {
	case 1:
		return false, nil
	case 2:
		return true, nil
	}
	return false, errorIsBlocked
}

type TransactionObserverMock int

func (TransactionObserverMock) Add(_ interface{}) {}

type TransactionMock int

func (TransactionMock) GetID() int              { return 1 }
func (TransactionMock) GetAccountID() int       { return 1 }
func (TransactionMock) GetOperationType() int   { return 1 }
func (TransactionMock) GetAmount() float64      { return 1 }
func (TransactionMock) GetEventDate() time.Time { return time.Now() }

func (TransactionMock) SetID(int)              {}
func (TransactionMock) SetAccountID(int)       {}
func (TransactionMock) SetOperationType(int)   {}
func (TransactionMock) SetAmount(float64)      {}
func (TransactionMock) SetEventDate(time.Time) {}
