package container

import (
	"github.com/raulinoneto/transactions-routines/internal/primary/httpadapters"
	"github.com/raulinoneto/transactions-routines/internal/secondary/observer"
	"github.com/raulinoneto/transactions-routines/internal/secondary/persistence"
	"github.com/raulinoneto/transactions-routines/pkg/domains/accounts"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
)

type Container struct {
	accountsService     *accounts.Service
	transactionsService *transactions.Service

	httpAdapter             *httpadapters.HTTPAdapter
	accountsHttpAdapter     *httpadapters.AccountsHttpAdapter
	transactionsHttpAdapter *httpadapters.TransactionsHttpAdapter

	mySqlAdapter             *persistence.MySqlAdapter
	accountMySqlAdapter      persistence.AccountAdapter
	transactionsMySqlAdapter transactions.TransactionRepository
	transactionObserver      *observer.TransactionsObserverAdapter
}
