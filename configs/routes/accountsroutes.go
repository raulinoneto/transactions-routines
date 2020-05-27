package routes

import (
	"net/http"

	"github.com/raulinoneto/transactions-routines/configs/container"
	"github.com/raulinoneto/transactions-routines/internal/primary/httpadapters"
)

func getAccountsRoutes(c *container.Container) httpadapters.Route {
	return httpadapters.Route{
		Pattern: "/accounts",
		Endpoints: []httpadapters.Endpoint{
			{
				Method:   http.MethodPost,
				Pattern:  "/",
				Function: c.GetAccountsHttpAdapter().CreateAccount,
			},
			{
				Method:   http.MethodGet,
				Pattern:  "/{id}",
				Function: c.GetAccountsHttpAdapter().GetAccount,
			},
		},
	}
}
