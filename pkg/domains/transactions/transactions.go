package transactions

import (
	"time"
)

type (
	// The object to be treated in transaction
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

	// Repository Behavior
	TransactionRepository interface {
		CreateTransaction(Transaction) error
	}

	// Account behavior needed to logical business
	AccountsRepository interface {
		BlockAccount(int) error
		UnlockAccount(int)
		AccountIsBlocked(int) (bool, error)
		CheckLimit(accountID int, value float64) error
		ChangeLimit(amount float64, id int) error
	}

	// Behavior for observer and treat transaction before
	TransactionObserver interface {
		Add(t interface{})
	}
)
