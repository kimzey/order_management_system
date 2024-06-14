package repository

import (
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
)

type orderRepositoryImpl struct {
	db database.Database
}

func NewOrderRepositoryImpl(db database.Database) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) Create(order *entities.Order) (*entities.Order, error) {

	modelOrder := ToOrderModel(order)
	if err := r.db.Connect().Create(&modelOrder).Preload("Transaction").Where("id = ?", modelOrder.ID).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("create order error "))
	}

	return modelOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) FindAll() (*[]entities.Order, error) {
	orders := new([]model.Order)

	if err := r.db.Connect().Preload("Transaction").Find(orders).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find all order error "))
	}
	allOrder := ConvertOrderModelsToEntities(orders)

	return allOrder, nil
}
func (r *orderRepositoryImpl) FindByID(id string) (*entities.Order, error) {
	order := new(model.Order)
	if err := r.db.Connect().Preload("Transaction").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find by id order error "))
	}
	return order.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Update(id string, order *entities.Order) (*entities.Order, error) {

	modelOrder := ToOrderModel(order)
	if err := r.db.Connect().Model(&modelOrder).Where("id = ?", id).Updates(&modelOrder).Scan(modelOrder).Where("id = ?", id).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error "))
	}

	return modelOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) UpdateStatus(id string, order *entities.Order) (*entities.Order, error) {
	orderModel := ToOrderModel(order)

	if err := r.db.Connect().Model(&orderModel).Where("id = ?", id).Updates(&orderModel).Scan(orderModel).Preload("Transaction").Where("id = ?", id).First(&orderModel).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error "))
	}
	return orderModel.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Delete(id string) error {
	if err := r.db.Connect().Where("id = ?", id).Delete(&model.Order{}).Error; err != nil {
		return errors.New(fmt.Sprintf("delete order error "))
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

//func ToOrderModelRes(e *entities.Order) *modelRes.Order {
//	return &modelRes.Order{
//		OrderID:       e.OrderID,
//		TransactionID: e.TransactionID,
//		//ProductID:     e.ProductID,
//		Status: e.Status,
//	}
//}
