package persistence

import (
	"github.com/raulinoneto/transactions-routines/pkg/domains/accounts"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
)

type AccountAdapter interface {
	accounts.Repository
	transactions.AccountsRepository
}

type AccountMySqlAdapter struct {
	driver *MySqlAdapter
}

func NewAccountMySqlAdapter(driver *MySqlAdapter) AccountAdapter {
	return &AccountMySqlAdapter{driver}
}

func (ma *AccountMySqlAdapter) CreateAccount(account accounts.Account) error {
	return nil
}

func (ma *AccountMySqlAdapter) GetAccount(id int) (accounts.Account, error) {
	return nil, nil
}

func (ma *AccountMySqlAdapter) BlockAccount(id int) error {
	return nil
}

func (ma *AccountMySqlAdapter) UnlockAccount(id int) {

}

func (ma *AccountMySqlAdapter) AccountIsBlocked(id int) (bool, error) {
	return true, nil
}
