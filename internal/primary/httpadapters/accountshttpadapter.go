package httpadapters

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/raulinoneto/transactions-routines/pkg/domains/accounts"
)

type AccountsHttpAdapter struct {
	service *accounts.Service
}

func NewAccountsHttpAdapter(service *accounts.Service) *AccountsHttpAdapter {
	return &AccountsHttpAdapter{service}
}

func (u *AccountsHttpAdapter) CreateAccount(w http.ResponseWriter, r *http.Request) {
	a := new(AccountBody)
	if err := parseBody(r, a); err != nil {
		BuildBadRequestResponse(err, w)
		return
	}
	err := u.service.SaveAccount(a)
	BuildCreatedResponse(a, err, w)
}

func (u *AccountsHttpAdapter) GetAccount(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		BuildBadRequestResponse(err, w)
		return
	}
	usr, err := u.service.GetAccount(id)
	BuildOkResponse(usr, err, w)
}
