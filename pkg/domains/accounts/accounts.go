package accounts

type (
	Account interface {
		GetDocumentNumber() int
		GetID() int
		SetID(id int)
		SetDocumentNumber(documentNumber int)
	}
	Repository interface {
		CreateAccount(a Account) error
		GetAccount(id int) (Account, error)
	}
)
