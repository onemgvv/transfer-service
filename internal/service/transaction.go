package service

import (
	"fmt"
	"transaction-service/internal/domain"
	"transaction-service/internal/repository"
)

type TransactionService struct {
	repo *repository.Repository
}

func NewTransactionService(repo *repository.Repository) *TransactionService {
	return &TransactionService{repo}
}

func (ts *TransactionService) Transfer(fromId, toId, amount uint) error {
	trn := domain.NewTransaction(fromId, toId, amount)
	id, err := ts.repo.Transaction.Create(trn)
	if err != nil {
		return err
	}

	fmt.Printf("Created transaction id: %d\n", id)

	pub, err := ts.repo.Account.GetById(fromId)
	if err != nil {
		return err
	}

	sub, err := ts.repo.Account.GetById(toId)
	if err != nil {
		return err
	}

	trn.Status = domain.TR_HOLD

	if err = pub.Transaction(sub, amount); err != nil {
		trn.Status = domain.TR_FAILED
		if err := ts.repo.Update(id, trn.Status); err != nil {
			return err
		}
		return err
	}

	if err = ts.repo.Account.Transfer(*sub, *pub); err != nil {
		trn.Status = domain.TR_FAILED
		if err := ts.repo.Update(id, trn.Status); err != nil {
			return err
		}
		return err
	}

	trn.Status = domain.TR_SUCCESS
	return ts.repo.Update(id, trn.Status)
}
