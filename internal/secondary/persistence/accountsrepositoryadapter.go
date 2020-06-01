package persistence

import (
	"github.com/raulinoneto/transactions-routines/internal/apierror"
	"github.com/raulinoneto/transactions-routines/pkg/domains/accounts"
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"log"
	"net/http"
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
		"INSERT INTO "+accountTableName+" (document_number, is_blocked, available_credit_limit) VALUES (?,?,?)",
		account.GetDocumentNumber(), 0, 1000,
	)
	if err != nil {
		return err
	}
	account.SetID(id)
	return nil
}

func (ma *AccountMySqlAdapter) GetAccount(id int) (accounts.Account, error) {
	var account *AccountsModel
	result, err := ma.driver.query("SELECT document_number FROM "+accountTableName+" WHERE id=?", id)
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
	id, err := ma.driver.exec("UPDATE "+accountTableName+" SET is_blocked=? WHERE id=?", 1, id)
	if err != nil {
		return err
	}
	return nil
}

func (ma *AccountMySqlAdapter) UnlockAccount(id int) {
	id, err := ma.driver.exec("UPDATE "+accountTableName+" SET is_blocked=? WHERE id=?", 0, id)
	if err != nil {
		log.Println(err)
	}
}

func (ma *AccountMySqlAdapter) ChangeLimit(amount float64, id int) error {
	id, err := ma.driver.exec("UPDATE "+accountTableName+" SET available_credit_limit=? WHERE id=?", amount, id)
	if err != nil {
		return err
	}
	return nil
}

func (ma *AccountMySqlAdapter) AccountIsBlocked(id int) (bool, error) {
	var isBlocked int
	result, err := ma.driver.query("SELECT is_blocked FROM "+accountTableName+" a WHERE id=?", id)
	if err != nil {
		return false, err
	}
	err = result.Scan(&isBlocked)
	if err != nil {
		return false, err
	}
	return isBlocked > 0, nil
}

func (ma *AccountMySqlAdapter) CheckLimit(accountId int, value float64) error {
	var limit float64
	result, err := ma.driver.query("SELECT available_credit_limit FROM "+accountTableName+" WHERE id=?", accountId)
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
