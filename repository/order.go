package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/szczynk/Assignment2/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository will create an object that represent the orderRepository interface
func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (or orderRepository) Fetch(c context.Context, m *[]models.Order) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = or.db.Debug().WithContext(ctx).Preload("Items").Find(&m).Error
	if err != nil {
		return err
	}
	return
}

func (or *orderRepository) Store(c context.Context, m *models.Order) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = or.db.Debug().WithContext(ctx).Create(&m).Error
	if err != nil {
		return err
	}
	return
}

func (or orderRepository) GetByID(c context.Context, m *models.Order, id string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = or.db.Debug().WithContext(ctx).Preload("Items").Where("id = ?", id).First(&m).Error
	if err != nil {
		return err
	}
	return
}

func (or *orderRepository) Update(c context.Context, m *models.Order, id string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = or.db.Debug().WithContext(ctx).Preload("Items").Where("id = ?", id).First(&models.Order{}).Error
	if err != nil {
		return err
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	m.ID = uint(orderId)
	err = or.db.Debug().WithContext(ctx).Save(&m).Error
	if err != nil {
		return err
	}
	return
}

func (or orderRepository) Delete(c context.Context, m *models.Order, id string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = or.db.Debug().WithContext(ctx).Preload("Items").Where("id = ?", id).First(&models.Order{}).Error
	if err != nil {
		return err
	}

	err = or.db.Debug().WithContext(ctx).Select("Items").Delete(&m).Error
	if err != nil {
		return err
	}
	return
}
