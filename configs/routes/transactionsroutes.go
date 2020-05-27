package routes

import (
	"net/http"

	"github.com/raulinoneto/transactions-routines/configs/container"
	"github.com/raulinoneto/transactions-routines/internal/primary/httpadapters"
)

func getAuthRoutes(c *container.Container) httpadapters.Route {
	return httpadapters.Route{
		Pattern: "/transactions",
		Endpoints: []httpadapters.Endpoint{
			{
				Method:   http.MethodPost,
				Pattern:  "/",
				Function: c.GetTransactionsHttpAdapter().CreateTransactions,
			},
		},
	}
}
