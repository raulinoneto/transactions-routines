package observer

import (
	"github.com/reactivex/rxgo/v2"
)

type TransactionsObserverAdapter struct {
	channel chan rxgo.Item
}

func NewTransactionsObserverAdapter() *TransactionsObserverAdapter {
	return &TransactionsObserverAdapter{make(chan rxgo.Item)}
}

func (o *TransactionsObserverAdapter) Add(transaction interface{}) {
	o.channel <- rxgo.Of(transaction)
}

func (o *TransactionsObserverAdapter) GetChannel() chan rxgo.Item {
	return o.channel
}
