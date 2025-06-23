package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"bwastartup/user"
	"errors"
	"log"

	"github.com/google/uuid"
)

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserID(userID string) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	// paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if input.User.ID != campaign.UserID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID string) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln(err.Error())
	}

	transaction := Transaction{}
	transaction.ID = uuid.String()
	transaction.Amount = input.Amount
	transaction.CampaignID = input.CampaignID
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	var paymentTrans = payment.Transaction{
		ID:     newTransaction.ID,
		Amount: int64(newTransaction.Amount),
	}
	var user = user.User{
		Name:  input.User.Name,
		Email: input.User.Email,
	}
	paymentURL, err := payment.NewService().GetPaymentURL(paymentTrans, user)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
