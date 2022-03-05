package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindById(id int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func CampaignRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	if err := r.db.Preload("CampaignImages", "is_primary =?", 1).Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	if err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "is_primary = ?", 1).Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindById(id int) (Campaign, error) {
	var campaign Campaign

	if err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", id).Find(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	if err := r.db.Create(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	if err := r.db.Save(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}
