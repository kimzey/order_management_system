package repository

import (
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/model"
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
	modelOrder := ToOrderModel(order)
	newOrder := new(model.Order)

	if err := r.db.Connect().Create(modelOrder).Scan(&newOrder).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return nil, err
	}

	if err := r.db.Connect().Preload("Product").Preload("Transaction").First(&newOrder, newOrder.ID).Error; err != nil {
		r.logger.Error("Failed to find order:", err.Error())
		return nil, err
	}
	fmt.Println("newOrder: ", newOrder)
	return newOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) FindAll() (*[]entities.Order, error) {
	orders := new([]model.Order)

	if err := r.db.Connect().Preload("Product").Preload("Transaction").Find(orders).Error; err != nil {
		return nil, err
	}
	allOrder := ConvertOrderModelsToEntities(orders)
	return allOrder, nil
}
func (r *orderRepositoryImpl) ChangeStatusNext(id uint64) (*entities.Order, error) {
	newOrder := new(model.Order)

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

	return newOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) ChangeStatusDone(id uint64) (*entities.Order, error) {
	newOrder := new(model.Order)

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

	return newOrder.ToOrderEntity(), nil
}

func ConvertOrderModelsToEntities(orders *[]model.Order) *[]entities.Order {
	entityOrders := new([]entities.Order)

	for _, order := range *orders {
		*entityOrders = append(*entityOrders, *order.ToOrderEntity())
	}

	return entityOrders
}
func ToOrderModel(e *entities.Order) *model.Order {
	return &model.Order{
		TransactionID: e.TransactionID,
		ProductID:     e.ProductID,
		Status:        e.Status,
	}
}
