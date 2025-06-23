package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID string) ([]Transaction, error)
	GetByUserID(userID string) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
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

func (r *repository) GetByUserID(userID string) ([]Transaction, error) {
	var transactions []Transaction
	// Param Preload ke-1 = relasi ke Campaign & relasi ke CampaignImages
	// Param Preload ke-2 = kondisi di CampaignImages dengan nama CampaignImages di db yaitu campaign_images
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
