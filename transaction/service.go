package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignId(campaignId int, user user.User) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func TransactionService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignId(campaignId int, user user.User) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindById(campaignId)
	if err != nil {
		return []Transaction{}, err
	}

	// Check if logged in user is the owner's campaign
	if campaign.User.ID != user.ID {
		return []Transaction{}, errors.New("you are not owner the campaign")
	}

	transactions, err := s.repository.GetCampaignId(campaignId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
