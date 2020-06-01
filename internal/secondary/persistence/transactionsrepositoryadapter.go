package persistence

import (
	"errors"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
)

type TransactionsMySqlAdapter struct {
	driver *MySqlAdapter
}

const InsufficientFoundsErrorCode = "account.insufficient_founds"

var (
	InsufficientFoundsError = errors.New("Insufficient Founds ")
	transactionTableName    = "transactions"
)

func NewTransactionsMySqlAdapter(driver *MySqlAdapter) transactions.TransactionRepository {
	return &TransactionsMySqlAdapter{driver}
}

func (ma *TransactionsMySqlAdapter) CreateTransaction(transaction transactions.Transaction) error {
	id, err := ma.driver.exec(
		"INSERT INTO "+transactionTableName+" (account_id, amount, operation_type) VALUES (?,?,?)",
		transaction.GetAccountID(), transaction.GetAmount(), transaction.GetOperationType(),
	)
	if err != nil {
		return err
	}
	transaction.SetID(id)
	return nil
}
