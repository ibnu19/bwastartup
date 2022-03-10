package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignId(campaignId int) ([]Transaction, error)
	GetByUserId(userId int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func TransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignId(campaignId int) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("User").
		Where("campaign_id = ?", campaignId).
		Order("created_at desc").
		Find(&transactions).Error; err != nil {

		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByUserId(userId int) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = ?", 1).
		Where("user_id =?", userId).
		Order(("created_at desc")).
		Find(&transactions).Error; err != nil {

		return transactions, err
	}

	return transactions, nil
}
