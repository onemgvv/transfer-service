package domain

import "errors"

type Account struct {
	Id      uint `json:"-" db:"id"`
	UserId  uint `json:"userId" db:"user_id"`
	Balance uint `json:"balance" db:"balance"`
}

func NewAccount(userId, balance uint) *Account {
	return &Account{
		UserId:  userId,
		Balance: balance,
	}
}

func (a *Account) Transaction(sub *Account, amount uint) error {
	if a.Balance < amount {
		return errors.New("не достаточно средств")
	}

	a.Balance -= amount
	sub.Balance += amount

	return nil
}
