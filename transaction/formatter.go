package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Amount     int       `json:"amount"`
	CampaignID string    `json:"campaign_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	var transactionFormatter CampaignTransactionFormatter
	transactionFormatter.ID = transaction.ID
	transactionFormatter.Name = transaction.User.Name
	transactionFormatter.Amount = transaction.Amount
	transactionFormatter.CampaignID = transaction.CampaignID
	transactionFormatter.CreatedAt = transaction.CreatedAt
	return transactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}
	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID        string            `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(transactions Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transactions.ID
	formatter.Amount = transactions.Amount
	formatter.Status = transactions.Status
	formatter.CreatedAt = transactions.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transactions.Campaign.Name
	campaignFormatter.ImageUrl = ""

	if len(transactions.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = transactions.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}
	var userTransactions []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		userTransactions = append(userTransactions, formatter)
	}

	return userTransactions
}

type TransactionFormatter struct {
	ID         string `json:"id"`
	CampaignID string `json:"campaign_id"`
	UserID     string `json:"user_id"`
	Amount     int64  `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.CampaignID = transaction.CampaignID
	formatter.UserID = transaction.UserID
	formatter.Amount = int64(transaction.Amount)
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL
	return formatter
}
