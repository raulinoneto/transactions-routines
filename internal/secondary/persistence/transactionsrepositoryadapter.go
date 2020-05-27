package persistence

import (
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
)

type TransactionsMySqlAdapter struct {
	driver *MySqlAdapter
}

func NewTransactionsMySqlAdapter(driver *MySqlAdapter) transactions.TransactionRepository {
	return &TransactionsMySqlAdapter{driver}
}

func (ma *TransactionsMySqlAdapter) CreateTransaction(transaction transactions.Transaction) error {
	return nil
}
