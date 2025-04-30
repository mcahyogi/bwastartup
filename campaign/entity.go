package campaign

import "time"

type Campaign struct {
	ID               string `gorm:"type:uuid;primaryKey"`
	UserID           string
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         string `gorm:"type:uuid;primaryKey"`
	CampaignID string
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
