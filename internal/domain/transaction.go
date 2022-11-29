package domain

import "time"

type TransactionStatus = string

const (
	TR_SUCCESS TransactionStatus = "success"
	TR_HOLD                      = "hold"
	TR_FAILED                    = "failed"
	TR_WAITING                   = "waiting"
)

type Transaction struct {
	Id           uint              `json:"-" db:"id"`
	PubId        uint              `json:"pubId" db:"pub_id"`               // Идентификатор аккаунта отправителя
	SubId        uint              `json:"subId" db:"sub_id"`               // Идентификатор аккаунта получателя
	Value        uint              `json:"value" db:"value"`                // сумма транзакции
	Status       TransactionStatus `json:"status" db:"status"`              // статус транзакции
	TransactDate time.Time         `json:"transactionDate" db:"created_at"` // дата создания транзакции
}

func NewTransaction(pubId, subId, value uint) *Transaction {
	return &Transaction{
		PubId:  pubId,
		SubId:  subId,
		Status: TR_WAITING,
		Value:  value,
	}
}
