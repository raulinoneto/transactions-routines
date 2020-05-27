package persistence

import (
	"errors"
	"github.com/raulinoneto/transactions-routines/internal/apierror"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"net/http"
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

func (ma *TransactionsMySqlAdapter) CheckLimit(accountId int, value float64) error {
	var limit float64
	result, err := ma.driver.query("SELECT COALESCE(SUM(amount),0) FROM "+transactionTableName+" WHERE account_id=?", accountId)
	if err != nil {
		return err
	}
	err = result.Scan(&limit)
	if err != nil {
		return err
	}
	if limit < value {
		return apierror.NewWarning(
			InsufficientFoundsErrorCode,
			InsufficientFoundsError.Error(),
			http.StatusBadRequest,
			InsufficientFoundsError,
		)
	}
	return nil
}
