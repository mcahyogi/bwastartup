package user

import (
	"time"
)

type User struct {
	ID             string `gorm:"type:uuid;primaryKey"`
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
