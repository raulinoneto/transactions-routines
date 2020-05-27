package transactions

import (
	"time"
)

type (
	Transaction interface {
		GetID() int
		GetAccountID() int
		GetOperationType() int
		GetAmount() float64
		GetEventDate() time.Time
		SetID(int)
		SetAccountID(int)
		SetOperationType(int)
		SetAmount(float64)
		SetEventDate(time.Time)
	}

	TransactionRepository interface {
		CreateTransaction(Transaction) error
		CheckLimit(accountID int, value float64) error
	}

	AccountsRepository interface {
		BlockAccount(int) error
		UnlockAccount(int)
		AccountIsBlocked(int) (bool, error)
	}

	TransactionObserver interface {
		Add(t interface{})
	}
)
