package model_rusers

import "time"

type UserDTO struct {
	ID           int       `gorm:"column:id" json:"id"`
	Username     string    `gorm:"column:username" json:"username"`
	Password     string    `gorm:"column:password" json:"password"`
	RefreshToken string    `gorm:"column:refresh_token" json:"refresh_token"`
	Email        string    `gorm:"column:email" json:"email"`
	IsActive     bool      `gorm:"column:is_active" json:"is_active"`
	CreatedBy    string    `gorm:"column:created_by" json:"created_by"`
	ModifiedBy   string    `gorm:"column:modified_by" json:"modified_by"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	ModifiedAt   time.Time `gorm:"column:modified_at" json:"modified_at"`
}
