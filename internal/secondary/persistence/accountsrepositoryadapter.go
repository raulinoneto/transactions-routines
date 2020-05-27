package persistence

import (
	"fmt"
	"github.com/raulinoneto/transactions-routines/pkg/domains/accounts"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"log"
)

type AccountAdapter interface {
	accounts.Repository
	transactions.AccountsRepository
}

type AccountMySqlAdapter struct {
	driver *MySqlAdapter
}

var accountTableName = "accounts"

func NewAccountMySqlAdapter(driver *MySqlAdapter) AccountAdapter {
	return &AccountMySqlAdapter{driver}
}

func (ma *AccountMySqlAdapter) CreateAccount(account accounts.Account) error {
	id, err := ma.driver.exec(
		fmt.Sprintf("INSERT INTO %s (document_number, is_blocked) VALUES (?,?)", accountTableName),
		account.GetDocumentNumber(), 0,
	)
	if err != nil {
		return err
	}
	account.SetID(id)
	return nil
}

func (ma *AccountMySqlAdapter) GetAccount(id int) (accounts.Account, error) {
	var account *AccountsModel
	result, err := ma.driver.query(
		fmt.Sprintf("SELECT document_number FROM %s WHERE id=?", accountTableName),
		id,
	)
	if err != nil {
		return nil, err
	}
	var documentNumber int
	err = result.Scan(&documentNumber)
	if err != nil {
		return nil, err
	}
	account = NewAccount(id, documentNumber, 0)
	return account, nil
}

func (ma *AccountMySqlAdapter) BlockAccount(id int) error {
	id, err := ma.driver.exec(
		fmt.Sprintf("UPDATE %s SET is_blocked=? WHERE id=?", accountTableName),
		1, id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ma *AccountMySqlAdapter) UnlockAccount(id int) {
	id, err := ma.driver.exec(
		fmt.Sprintf("UPDATE %s SET is_blocked=? WHERE id=?", accountTableName),
		0, id,
	)
	if err != nil {
		log.Println(err)
	}
}

func (ma *AccountMySqlAdapter) AccountIsBlocked(id int) (bool, error) {
	var isBlocked int
	result, err := ma.driver.query(
		fmt.Sprintf("SELECT is_blocked FROM %s WHERE id=?", accountTableName),
		id,
	)
	if err != nil {
		return false, err
	}
	err = result.Scan(&isBlocked)
	if err != nil {
		return false, err
	}
	return isBlocked > 0, nil
}
