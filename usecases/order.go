package usecases

import (
	"context"

	"github.com/szczynk/Assignment2/models"
)

type orderUsecase struct {
	or models.OrderRepository
}

func NewOrderUsecase(o models.OrderRepository) *orderUsecase {
	return &orderUsecase{or: o}
}

func (ouc *orderUsecase) Fetch(c context.Context, m *[]models.Order) (err error) {
	if err = ouc.or.Fetch(c, m); err != nil {
		return err
	}
	return
}

func (ouc *orderUsecase) Store(c context.Context, m *models.Order) (err error) {
	if err = ouc.or.Store(c, m); err != nil {
		return err
	}
	return
}

func (ouc *orderUsecase) GetByID(c context.Context, m *models.Order, id string) (err error) {
	if err = ouc.or.GetByID(c, m, id); err != nil {
		return err
	}
	return
}

func (ouc *orderUsecase) Update(c context.Context, m *models.Order, id string) (err error) {
	if err = ouc.or.Update(c, m, id); err != nil {
		return err
	}
	return
}

func (ouc *orderUsecase) Delete(c context.Context, m *models.Order, id string) (err error) {
	if err = ouc.or.Delete(c, m, id); err != nil {
		return err
	}
	return
}
