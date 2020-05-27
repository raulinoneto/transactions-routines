package routes

import (
	"github.com/raulinoneto/transactions-routines/configs/container"
	"github.com/raulinoneto/transactions-routines/internal/primary/httpadapters"
)

func GetRoutes(c *container.Container) []httpadapters.Route {
	return []httpadapters.Route{
		{
			Pattern:   "/",
			Endpoints: []httpadapters.Endpoint{},
			SubRoutes: []httpadapters.Route{
				getAuthRoutes(c),
				getAccountsRoutes(c),
			},
		},
	}
}
