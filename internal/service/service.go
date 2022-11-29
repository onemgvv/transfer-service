package service

import "transaction-service/internal/repository"

type Account interface {
}

type Transaction interface {
	Transfer(fromId, toId, amount uint) error
}

type Service struct {
	Account
	Transaction
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Account:     NewAccountService(repo),
		Transaction: NewTransactionService(repo),
	}
}
