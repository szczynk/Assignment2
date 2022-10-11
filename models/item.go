package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Item struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	OrderID     uint       `json:"order_id"`
	ItemCode    string     `json:"item_code" valid:"required-item_code is required" example:"BXC-100" gorm:"not null"`
	Description string     `json:"description" valid:"required-item_code is required" example:"Fancy Glass" gorm:"not null"`
	Quantity    int        `json:"quantity" valid:"required-quantity is required" example:"3" gorm:"not null"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (p *Item) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}

func (p *Item) BeforeUpdate(tx *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}
