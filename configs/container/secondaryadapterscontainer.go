package container

import (
	"github.com/raulinoneto/transactions-routines/internal/secondary/persistence"
	"github.com/raulinoneto/transactions-routines/internal/secondary/observer"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"os"
)

func (c *Container) GetMySqlAdapter() *persistence.MySqlAdapter {
	if c.mySqlAdapter == nil {
		c.mySqlAdapter = persistence.NewMySqlAdapter(
			os.Getenv("DB_ADAPTER"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_ADDR"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		c.mySqlAdapter.TestConnection()
	}
	return c.mySqlAdapter
}

func (c *Container) GetAccountMySqlAdapter() persistence.AccountAdapter {
	if c.accountMySqlAdapter == nil {
		c.accountMySqlAdapter = persistence.NewAccountMySqlAdapter(c.GetMySqlAdapter())
	}
	return c.accountMySqlAdapter
}

func (c *Container) GetTransactionsMySqlAdapter() transactions.TransactionRepository {
	if c.transactionsMySqlAdapter == nil {
		c.transactionsMySqlAdapter = persistence.NewTransactionsMySqlAdapter(c.GetMySqlAdapter())
	}
	return c.transactionsMySqlAdapter
}

func (c *Container) GetTransactionsObserverAdapter() *observer.TransactionsObserverAdapter {
	if c.transactionObserver == nil {
		c.transactionObserver = observer.NewTransactionsObserverAdapter(c.GetTransactionsService())
	}
	return c.transactionObserver
}
