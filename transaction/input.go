package transaction

type GetCampaignTransactionsInput struct {
	ID int `json:"id" binding:"required"`
}
