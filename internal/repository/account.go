package repository

import (
	"context"
	"database/sql"
	"fmt"
	"transaction-service/internal/delivery/dto"
	"transaction-service/internal/domain"

	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) Create(acc dto.CreateAccountDto) (uint, error) {
	var id uint
	query := fmt.Sprintf("INSERT INTO %s (user_id, balance) VALUES ($1, $2) RETURNING id", accountsTable)
	row := ar.db.QueryRow(query, acc.UserId, acc.Balance)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Err(); err != nil {
		return 0, err
	}

	return id, nil
}

func (ar *AccountRepository) GetById(id uint) (*domain.Account, error) {
	var account domain.Account
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", accountsTable)
	row := ar.db.QueryRowx(query, id)
	if err := row.StructScan(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (ar *AccountRepository) GetByUser(userId uint) (*domain.Account, error) {
	var account domain.Account
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", accountsTable)
	row := ar.db.QueryRowx(query, userId)
	if err := row.StructScan(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (ar *AccountRepository) Transfer(sub, pub domain.Account) error {
	tx, err := ar.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	updateBalancePub := fmt.Sprintf("UPDATE %s SET balance = $1 WHERE user_id = $2", accountsTable)
	updateBalanceSub := fmt.Sprintf("UPDATE %s SET balance = $1 WHERE user_id = $2", accountsTable)

	if _, err = tx.Exec(updateBalancePub, pub.Balance, pub.UserId); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if _, err = tx.Exec(updateBalanceSub, sub.Balance, sub.UserId); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
