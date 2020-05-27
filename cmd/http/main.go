package main

import (
	"log"
	"os"

	"github.com/raulinoneto/transactions-routines/configs/container"
	"github.com/raulinoneto/transactions-routines/configs/routes"
)

func main() {
	c := new(container.Container)
	log.Println("STARTING OBSERVER")
	c.GetTransactionsObserverAdapter().Observe()
	log.Println("OBSERVER STARTED")

	log.Println("SERVER IS STARTING")
	ha := c.GetHTTPAdapter(routes.GetRoutes(c))

	port := os.Getenv("PORT")
	log.Println("running at: " + os.Getenv("HOST") + ":" + port)
	ha.Serve(port)
}
