package observer

import (
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"github.com/reactivex/rxgo/v2"
)

type TransactionsObserverAdapter struct {
	channel chan rxgo.Item
	service *transactions.Service
}

func NewTransactionsObserverAdapter(service *transactions.Service) *TransactionsObserverAdapter {
	return &TransactionsObserverAdapter{make(chan rxgo.Item), service}
}

func (o *TransactionsObserverAdapter) Add(transaction transactions.Transaction) {
	o.channel <- rxgo.Of(transaction)
}

func (o *TransactionsObserverAdapter) Observe() {
	observable := rxgo.FromChannel(o.channel)
	for item := range observable.Observe() {
		o.service.SaveTransaction(item.V.(transactions.Transaction))
	}
}
