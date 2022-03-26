package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"bwastartup/user"
	"errors"
	"strconv"
)

type Service interface {
	GetTransactionsByCampaignId(campaignId int, user user.User) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func TransactionService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
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

	transactions, err := s.repository.GetByCampaignId(campaignId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserId(userId int) ([]Transaction, error) {

	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}

// Create new transaction for backer
func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignId
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.UserID = input.User.ID

	// Create new tansaction
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	// Mapping transaction entity to payment transaction entity
	paymentTransaction := payment.Transaction{
		Id:     newTransaction.Id,
		Amount: newTransaction.Amount,
	}

	// Get Payment redirect URL from midtrans
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	// Add payment URL on transaction object
	newTransaction.PaymentURL = paymentURL

	// Update transaction
	return s.repository.Update(newTransaction)
}

// Notification after process payment from user
func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transactionId, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetById(transactionId)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "canceled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindById(updatedTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
