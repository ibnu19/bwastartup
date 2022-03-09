package transaction

type GetCampaignTransactionsInput struct {
	Id int `uri:"id" binding:"required"`
}
