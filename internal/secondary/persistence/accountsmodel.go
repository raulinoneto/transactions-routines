package persistence

type AccountsModel struct {
	ID, DocumentNumber int
	IsBlocked          uint8
}

func NewAccount(id, documentNumber int, isBlocked uint8) *AccountsModel {
	return &AccountsModel{id, documentNumber, isBlocked}
}

func (a *AccountsModel) GetID() int {
	return a.ID
}

func (a *AccountsModel) GetDocumentNumber() int {
	return a.DocumentNumber
}

func (a *AccountsModel) SetID(id int) {
	a.ID = id
}

func (a *AccountsModel) SetDocumentNumber(documentNumber int) {
	a.DocumentNumber = documentNumber
}
