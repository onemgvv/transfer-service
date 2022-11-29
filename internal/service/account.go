package service

import "transaction-service/internal/repository"

type AccountService struct {
	repo *repository.Repository
}

func NewAccountService(repo *repository.Repository) *AccountService {
	return &AccountService{repo}
}
