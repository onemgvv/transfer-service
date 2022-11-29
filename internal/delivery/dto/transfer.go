package dto

type CreateTransferDto struct {
	PubId  uint `json:"pub_id"`
	SubId  uint `json:"sub_id"`
	Amount uint `json:"amount"`
}
