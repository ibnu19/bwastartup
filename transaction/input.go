package transaction

import "bwastartup/user"

type GetCampaignTransactionsInput struct {
	Id int `uri:"id" binding:"required"`
}

type CreateTransactionInput struct {
	CampaignId int `json:"campaign_id" binding:"required"`
	Amount     int `json:"amount" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
