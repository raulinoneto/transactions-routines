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

var accountTableName = "account"

func NewAccountMySqlAdapter(driver *MySqlAdapter) AccountAdapter {
	return &AccountMySqlAdapter{driver}
}

func (ma *AccountMySqlAdapter) CreateAccount(account accounts.Account) error {
	db := ma.driver.Open()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("INSERT INTO " + accountTableName + " () VALUES ()")
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	account.SetID(int(id))
	return tx.Commit()
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
