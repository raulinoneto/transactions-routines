package httpadapters

import (
	"net/http"

	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
)

type TransactionsHttpAdapter struct {
	service *transactions.Service
}

func NewTransactionsHttpAdapter(service *transactions.Service) *TransactionsHttpAdapter {
	return &TransactionsHttpAdapter{service}
}

func (t *TransactionsHttpAdapter) CreateTransactions(w http.ResponseWriter, r *http.Request) {
	tb := new(TransactionBody)
	if err := parseBody(r, tb); err != nil {
		BuildBadRequestResponse(err, w)
		return
	}
	err := t.service.SaveTransaction(tb)
	BuildCreatedResponse(tb, err, w)
}
