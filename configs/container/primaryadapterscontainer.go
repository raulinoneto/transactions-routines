package container

import "github.com/raulinoneto/transactions-routines/internal/primary/httpadapters"

func (c *Container) GetHTTPAdapter(routes []httpadapters.Route) *httpadapters.HTTPAdapter {
	if c.httpAdapter == nil {
		c.httpAdapter = httpadapters.NewHTTPAdapter(routes, 30)
	}
	return c.httpAdapter
}

func (c *Container) GetAccountsHttpAdapter() *httpadapters.AccountsHttpAdapter {
	if c.accountsHttpAdapter == nil {
		c.accountsHttpAdapter = httpadapters.NewAccountsHttpAdapter(c.GetAccountsService())
	}

	return c.accountsHttpAdapter
}

func (c *Container) GetTransactionsHttpAdapter() *httpadapters.TransactionsHttpAdapter {
	if c.transactionsHttpAdapter == nil {
		c.transactionsHttpAdapter = httpadapters.NewTransactionsHttpAdapter(c.GetTransactionsService())
	}
	return c.transactionsHttpAdapter
}
