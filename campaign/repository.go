package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func CampaignRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (s *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	if err := s.db.Preload("CampaignImages").Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *repository) FindByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	if err := s.db.Where("user_id = ?", userId).
		Preload("CampaignImages", "is_primary = ?", 1).
		Find(&campaigns).Error; err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
