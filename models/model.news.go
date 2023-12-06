package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	ID        string    `json:"id" gorm:"primary_key;"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Body      string    `json:"body" gorm:"type:text;not null"`
	CreatedBy string    `json:"created_by" gorm:"index;not null"`
	UpdatedBy string    `json:"updated_by" gorm:"index;default:null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Reference to the User struct
	CreatedUser Users `json:"user_created" gorm:"foreignkey:CreatedBy"`
	UpdatedUser Users `json:"user_updated" gorm:"foreignkey:UpdatedBy"`
}

func (model *News) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now().Local()
	return nil
}

func (model *News) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
