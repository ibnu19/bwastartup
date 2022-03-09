package transaction

import "gorm.io/gorm"

type Repository interface {
	GetCampaignId(campaignId int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func TransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignId(campaignId int) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Order("created_at desc").Find(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}
