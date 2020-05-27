package observer

import (
	"github.com/raulinoneto/transactions-routines/pkg/domains/transactions"
	"github.com/reactivex/rxgo/v2"
	"log"
	"time"
)

func Observe(ch chan rxgo.Item, service *transactions.Service) {
	observable := rxgo.FromChannel(ch)
	for item := range observable.Observe() {
		time.Sleep(time.Second * 2)
		t := item.V.(transactions.Transaction)
		log.Println(t)
		err := service.SaveTransaction(t)
		if err != nil {
			ch <- item
		}
	}
}
