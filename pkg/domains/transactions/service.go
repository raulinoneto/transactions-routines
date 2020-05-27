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

func (s Service) SaveTransaction(t Transaction) error {
	isBlocked, err := s.accountsRepo.AccountIsBlocked(t.GetAccountID())
	if err != nil {
		return err
	}
	if isBlocked {
		s.observer.Add(t)
		return nil
	}

	if t.GetOperationType() != core.OperationsTypeDeposit && t.GetAmount() > 0 {
		t.SetAmount(0 - t.GetAmount())
	}

	err = s.accountsRepo.BlockAccount(t.GetAccountID())
	if err != nil {
		return err
	}

	err = s.transactionsRepo.CheckLimit(t.GetAccountID(), t.GetAmount())
	if err != nil {
		return err
	}
	defer s.accountsRepo.UnlockAccount(t.GetAccountID())
	return s.transactionsRepo.CreateTransaction(t)
}
