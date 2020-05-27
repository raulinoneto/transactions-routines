package httpadapters

type AccountBody struct {
	ID, DocumentNumber int
}

func NewAccount(id, documentNumber int) *AccountBody {
	return &AccountBody{id, documentNumber}
}

func (a *AccountBody) GetID() int {
	return a.ID
}

func (a *AccountBody) GetDocumentNumber() int {
	return a.DocumentNumber
}

func (a *AccountBody) SetID(id int) {
	a.ID = id
}

func (a *AccountBody) SetDocumentNumber(documentNumber int) {
	a.DocumentNumber = documentNumber
}
