package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID string) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID string) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("User").Order("id desc").Where("campaign_id = ?", campaignID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
