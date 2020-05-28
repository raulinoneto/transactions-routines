package accounts

type (
	// Account model to treat in logical business
	Account interface {
		GetDocumentNumber() int
		GetID() int
		SetID(id int)
		SetDocumentNumber(documentNumber int)
	}
	// Repository needed to use the logic
	Repository interface {
		CreateAccount(a Account) error
		GetAccount(id int) (Account, error)
	}
)
