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
