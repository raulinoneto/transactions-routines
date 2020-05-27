package main

import (
	"github.com/raulinoneto/transactions-routines/internal/primary/rx"
	"log"
	"os"

	"github.com/raulinoneto/transactions-routines/configs/container"
	"github.com/raulinoneto/transactions-routines/configs/routes"
)

func main() {
	c := new(container.Container)
	log.Println("STARTING OBSERVER")
	go rx.Observe(
		c.GetTransactionsObserverAdapter().GetChannel(),
		c.GetTransactionsService(),
	)
	log.Println("OBSERVER STARTED")

	log.Println("SERVER IS STARTING")
	ha := c.GetHTTPAdapter(routes.GetRoutes(c))

	port := os.Getenv("PORT")
	log.Println("running at: " + os.Getenv("HOST") + ":" + port)
	ha.Serve(port)
}
