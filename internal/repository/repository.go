package repository

import (
	"transaction-service/internal/delivery/dto"
	"transaction-service/internal/domain"

	"github.com/jmoiron/sqlx"
)

const (
	accountsTable     = "accounts"
	transactionsTable = "transactions"
)

type Account interface {
	Create(acc dto.CreateAccountDto) (uint, error)
	GetById(id uint) (*domain.Account, error)
	GetByUser(userId uint) (*domain.Account, error)
	Transfer(sub, pub domain.Account) error
}

type Transaction interface {
	Create(trn *domain.Transaction) (uint, error)
	All() (domain.Transaction, error)
	ByStatus(status domain.TransactionStatus) ([]domain.Transaction, error)
	Update(id uint, status domain.TransactionStatus) error
}

type Repository struct {
	Account
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account:     NewAccountRepository(db),
		Transaction: NewTransactionRepository(db),
	}
}
