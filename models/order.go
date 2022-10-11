package models

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	CustomerName string     `json:"customer_name" valid:"required-customer_name is required" example:"John Dee" gorm:"not null"`
	OrderedAt    time.Time  `json:"ordered_at" valid:"required-ordered_at is required" example:"2022-10-07T18:19:24.161481554+07:00" gorm:"not null"`
	Items        []Item     `json:"items"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

func (p *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}

func (p *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}

// OrderUsecase represent the order's usecases
type OrderUsecase interface {
	Fetch(context.Context, *[]Order) error
	Store(context.Context, *Order) error
	GetByID(context.Context, *Order, string) error
	Update(context.Context, *Order, string) error
	Delete(context.Context, *Order, string) error
}

// OrderRepository represent the order's repository
type OrderRepository interface {
	Fetch(context.Context, *[]Order) error
	Store(context.Context, *Order) error
	GetByID(context.Context, *Order, string) error
	Update(context.Context, *Order, string) error
	Delete(context.Context, *Order, string) error
}
