package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/labstack/echo/v4"
)

type orderRepositoryImpl struct {
	db     database.Database
	logger echo.Logger
}

func NewOrderRepositoryImpl(db database.Database, logger echo.Logger) OrderRepository {
	return &orderRepositoryImpl{db: db, logger: logger}
}

func (r *orderRepositoryImpl) Create(order *entities.Order) (*entities.Order, error) {

	newOrder := new(entities.Order)
	if err := r.db.Connect().Create(order).Scan(&newOrder).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return nil, err
	}
	return newOrder, nil
}

func (r *orderRepositoryImpl) ChangeStatusNext(id uint64) (*entities.Order, error) {
	newOrder := new(entities.Order)

	if err := r.db.Connect().First(&newOrder, id).Error; err != nil {
		r.logger.Error("Failed to find order:", err.Error())
		return nil, err
	}

	if err := newOrder.NextStatus(); err != nil {
		r.logger.Error("Next Status failed:", err.Error())
		return nil, err
	}

	if err := r.db.Connect().Save(&newOrder).Error; err != nil {
		r.logger.Error("Failed to save order:", err.Error())
		return nil, err
	}

	return newOrder, nil
}

func (r orderRepositoryImpl) ChageStatusDone(id uint64) (*entities.Order, error) {
	newOrder := new(entities.Order)

	if err := r.db.Connect().First(&newOrder, id).Error; err != nil {
		r.logger.Error("Failed to find order:", err.Error())
		return nil, err
	}

	if err := newOrder.NextPaidToDone(); err != nil {
		r.logger.Error("Next Status failed:", err.Error())
		return nil, err
	}

	if err := r.db.Connect().Save(&newOrder).Error; err != nil {
		r.logger.Error("Failed to save order:", err.Error())
		return nil, err
	}

	return newOrder, nil
}
