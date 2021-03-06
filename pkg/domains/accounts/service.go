package accounts

import (
	"errors"
	"net/http"

	"github.com/raulinoneto/transactions-routines/internal/apierror"
)

const InvalidDocumentNumberErrorCode = "account.invalid_document_number"

var InvalidDocumentNumberError = errors.New("Invalid Document ")

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

var errorDocumentInvalid = apierror.NewWarning(
	InvalidDocumentNumberErrorCode,
	InvalidDocumentNumberError.Error(),
	http.StatusBadRequest,
	InvalidDocumentNumberError,
)

// Logic business to save an account
func (s *Service) SaveAccount(a Account) error {
	if !documentNumberIsValid(a.GetDocumentNumber()) {
		return errorDocumentInvalid
	}
	return s.repo.CreateAccount(a)
}

func (s *Service) GetAccount(id int) (Account, error) {
	return s.repo.GetAccount(id)
}

func documentNumberIsValid(documentNumber int) bool {
	if documentNumber <= 0 {
		return false
	}
	return true
}
