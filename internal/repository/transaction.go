package repository

import (
	"fmt"
	"transaction-service/internal/domain"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (tr *TransactionRepository) Create(trn *domain.Transaction) (uint, error) {
	var id uint
	query := fmt.Sprintf("INSERT INTO %s (pub_id, sub_id, value, status) VALUES ($1, $2, $3, $4) RETURNING id", transactionsTable)

	row := tr.db.QueryRow(query, trn.PubId, trn.SubId, trn.Value, trn.Status)

	if err := row.Err(); err != nil {
		return 0, err
	}
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (tr *TransactionRepository) All() (domain.Transaction, error) {
	var transaction domain.Transaction

	query := fmt.Sprintf("SELECT * FROM %s", transactionsTable)

	if err := tr.db.Select(transaction, query); err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (tr *TransactionRepository) ByStatus(status domain.TransactionStatus) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	query := fmt.Sprintf("SELECT * FROM %s WHERE status = $1", transactionsTable)
	if err := tr.db.Select(transactions, query, status); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (tr *TransactionRepository) Update(id uint, status domain.TransactionStatus) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", transactionsTable)
	row := tr.db.QueryRow(query, status, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil
}