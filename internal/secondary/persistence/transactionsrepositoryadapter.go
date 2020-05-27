package persistence

import (
	"errors"
	"fmt"
	"github.com/raulinoneto/transactions-routines/internal/apierror"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"net/http"
)

type TransactionsMySqlAdapter struct {
	driver *MySqlAdapter
}

const InsufficientFoundsErrorCode = "account.insuficient_founds"

var (
	InsufficientFoundsError = errors.New("Invalid Document ")
	transactionTableName    = "transactions"
)

func NewTransactionsMySqlAdapter(driver *MySqlAdapter) transactions.TransactionRepository {
	return &TransactionsMySqlAdapter{driver}
}

func (ma *TransactionsMySqlAdapter) CreateTransaction(transaction transactions.Transaction) error {
	id, err := ma.driver.exec(
		fmt.Sprintf(""+
			"INSERT INTO %s (account_id, amount, operation_type) VALUES (?,?,?)",
			transactionTableName,
		),
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
	result, err := ma.driver.query(
		fmt.Sprintf("SELECT COUNT(amount) as limit FROM %s WHERE account_id=?", accountTableName),
		accountId,
	)
	if err != nil {
		return err
	}
	err = result.Scan(&limit)
	if err != nil {
		return err
	}
	if limit > value {
		return apierror.NewWarning(
			InsufficientFoundsErrorCode,
			InsufficientFoundsError.Error(),
			http.StatusBadRequest,
			InsufficientFoundsError,
		)
	}
	return nil
}
