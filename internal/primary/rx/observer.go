package rx

import (
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"github.com/reactivex/rxgo/v2"
	"time"
)

func Observe(ch chan rxgo.Item, service *transactions.Service) {
	observable := rxgo.FromChannel(ch)
	for item := range observable.Observe() {
		time.Sleep(time.Second * 2)
		err := service.SaveTransaction(item.V.(transactions.Transaction))
		if err != nil {
			// ToDo Retry or notify
		}
	}
}
