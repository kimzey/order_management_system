package repository

import (
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type orderRepositoryImpl struct {
	db database.Database
}

func NewOrderRepositoryImpl(db database.Database) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) Create(ecommerce *_interface.Ecommerce) (*entities.Order, error) {

	tx := r.db.Connect().Begin()
	if tx.Error != nil {
		return nil, errors.New("failed to start transaction")
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	modelOrder := ToOrderModel(ecommerce.Order)
	stock := new(model.Stock)

	if err := tx.Where("product_id = ?", ecommerce.ProductID).First(&stock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("stock not found: %s", err.Error()))
	}

	if stock.Quantity < ecommerce.TransactionQuantity {
		return nil, errors.New("stock not enough")
	}
	stock.Quantity -= ecommerce.TransactionQuantity

	if err := tx.Model(&stock).Where(
		"id = ? AND quantity >= ?", stock.ID, ecommerce.TransactionQuantity).Updates(&stock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to update stock: %s", err.Error()))
	}
	if err := tx.Create(&modelOrder).Preload("Transaction").Where("id = ?", modelOrder.ID).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("create order error : %s", err.Error()))
	}
	return modelOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) FindAll() (*[]entities.Order, error) {
	orders := new([]model.Order)

	if err := r.db.Connect().Preload("Transaction").Find(orders).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find all order error : %s", err.Error()))
	}
	allOrder := ConvertOrderModelsToEntities(orders)

	return allOrder, nil
}
func (r *orderRepositoryImpl) FindByID(id string) (*entities.Order, error) {
	order := new(model.Order)
	if err := r.db.Connect().Preload("Transaction").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find by id order error : %s", err.Error()))
	}
	return order.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Update(id string, ecommerce *_interface.Ecommerce) (*entities.Order, error) {
	tx := r.db.Connect().Begin()
	if tx.Error != nil {
		return nil, errors.New("failed to start transaction")
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	modelOrder := ToOrderModel(ecommerce.Order)
	stock := new(model.Stock)

	if err := tx.Where("product_id = ?", ecommerce.ProductID).First(&stock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("stock not found: %s", err.Error()))
	}

	if stock.Quantity < ecommerce.TransactionQuantity {
		return nil, errors.New("stock not enough")
	}
	stock.Quantity -= ecommerce.TransactionQuantity

	if err := tx.Model(&stock).Where(
		"id = ? AND quantity >= ?", stock.ID, ecommerce.TransactionQuantity).Updates(&stock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to update stock: %s", err.Error()))
	}

	if err := r.db.Connect().Model(&modelOrder).Where("id = ?", id).Updates(&modelOrder).Scan(modelOrder).Where("id = ?", id).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error : %s", err.Error()))
	}
	return modelOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) UpdateStatus(id string, order *entities.Order) (*entities.Order, error) {
	orderModel := ToOrderModel(order)

	if err := r.db.Connect().Model(&orderModel).Where("id = ?", id).Updates(&orderModel).Scan(orderModel).Preload("Transaction").Where("id = ?", id).First(&orderModel).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error : %s", err.Error()))
	}
	return orderModel.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Delete(id string) error {
	if err := r.db.Connect().Where("id = ?", id).Delete(&model.Order{}).Error; err != nil {
		return errors.New(fmt.Sprintf("delete order error : %s", err.Error()))
	}
	return nil
}

func ConvertOrderModelsToEntities(orders *[]model.Order) *[]entities.Order {
	entityOrders := new([]entities.Order)

	for _, order := range *orders {
		*entityOrders = append(*entityOrders, *order.ToOrderEntity())
	}

	return entityOrders
}
func ToOrderModel(e *entities.Order) *model.Order {
	//fmt.Println("e: ", e.IsDomestic)
	return &model.Order{
		TransactionID: e.TransactionID,
		Status:        e.Status,
	}
}

func ToOrderModelRes(e *entities.Order) *modelRes.Order {
	return &modelRes.Order{
		OrderID:       e.OrderID,
		TransactionID: e.TransactionID,
		//ProductID:     e.ProductID,
		Status: e.Status,
	}
}
