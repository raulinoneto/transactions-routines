package container

import (
	"github.com/raulinoneto/transactions-routines/pkg/domains/accounts"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
)

func (c *Container) GetAccountsService() *accounts.Service {
	if c.accountsService == nil {
		c.accountsService = accounts.NewService(c.GetAccountMySqlAdapter())
	}

	return c.accountsService
}

func (c *Container) GetTransactionsService() *transactions.Service {
	if c.transactionsService == nil {
		c.transactionsService = transactions.NewService(
			c.GetTransactionsMySqlAdapter(),
			c.GetAccountMySqlAdapter(),
			c.GetTransactionsObserverAdapter(),
		)
	}
	return c.transactionsService
}
