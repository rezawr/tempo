package models

import (
	"tempo/pkg"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string    `json:"id" gorm:"primary_key;"`
	Fullname  string    `json:"fullname,omitempty" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email,omitempty" gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"password,omitempty" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (model *Users) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	model.Password = pkg.HashPassword(model.Password)
	model.CreatedAt = time.Now().Local()
	return nil
}

func (model *Users) BeforeUpdate(db *gorm.DB) error {
	model.UpdatedAt = time.Now().Local()
	return nil
}
