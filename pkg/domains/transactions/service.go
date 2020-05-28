package transactions

import "github.com/raulinoneto/transactions-routines/pkg/core"

type Service struct {
	transactionsRepo TransactionRepository
	accountsRepo     AccountsRepository
	observer         TransactionObserver
}

func NewService(
	transactionsRepo TransactionRepository,
	accountsRepo AccountsRepository,
	observer TransactionObserver,
) *Service {
	return &Service{transactionsRepo, accountsRepo, observer}
}

// Verify all
func (s Service) SaveTransaction(t Transaction) error {
	isBlocked, err := s.accountsRepo.AccountIsBlocked(t.GetAccountID())
	if err != nil {
		return err
	}
	if isBlocked {
		s.observer.Add(t)
		return nil
	}
	err = s.accountsRepo.BlockAccount(t.GetAccountID())
	if err != nil {
		return err
	}
	defer s.accountsRepo.UnlockAccount(t.GetAccountID())

	if t.GetOperationType() != core.OperationTypeDeposit {
		if err = s.transactionsRepo.CheckLimit(t.GetAccountID(), t.GetAmount()); err != nil {
			return err
		}
		if t.GetAmount() > 0 {
			t.SetAmount(-1 * t.GetAmount())
		}
	}
	return s.transactionsRepo.CreateTransaction(t)
}
