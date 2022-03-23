package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetByCampaignId(campaignId int) ([]Transaction, error)
	GetByUserId(userId int) ([]Transaction, error)
	GetById(Id int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func TransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get transaction by campaign id
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

// Get transaction by user
func (r *repository) GetByUserId(userId int) ([]Transaction, error) {
	var transactions []Transaction

	if err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = ?", 1).
		Where("user_id = ?", userId).
		Order(("created_at desc")).
		Find(&transactions).Error; err != nil {

		return transactions, err
	}

	return transactions, nil
}

// Get transaction by id
func (r *repository) GetById(Id int) (Transaction, error) {
	var transaction Transaction

	if err := r.db.Where("id = ?", Id).Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

// Create new transaction to database
func (r *repository) Save(transaction Transaction) (Transaction, error) {

	if err := r.db.Create(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

// Update transaction on database
func (r *repository) Update(transaction Transaction) (Transaction, error) {

	if err := r.db.Save(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}
