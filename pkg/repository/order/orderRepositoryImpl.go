package order

import (
	"context"
	customTracer "github.com/kizmey/order_management_system/tracer"

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

func (r *orderRepositoryImpl) Create(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderCreateRepository")
	defer sp.End()

	modelOrder := ToOrderModel(order)
	if err := r.db.Connect().Create(&modelOrder).Preload("Transaction").Where("id = ?", modelOrder.ID).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("create order error "))
	}

	customTracer.SetSubAttributesWithJson(modelOrder.ToOrderEntity(), sp)

	return modelOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) FindAll(ctx context.Context) (*[]entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderFindAllRepository")
	defer sp.End()

	orders := new([]model.Order)

	if err := r.db.Connect().Preload("Transaction").Find(orders).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find all order error "))
	}
	allOrder := ConvertOrderModelsToEntities(orders)

	return allOrder, nil
}
func (r *orderRepositoryImpl) FindByID(ctx context.Context, id string) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderFindByIdRepository")
	defer sp.End()

	order := new(model.Order)
	if err := r.db.Connect().Preload("Transaction").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find by id order error "))
	}
	return order.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Update(ctx context.Context, id string, order *entities.Order) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderUpdateRepository")
	defer sp.End()

	modelOrder := ToOrderModel(order)
	if err := r.db.Connect().Model(&modelOrder).Where("id = ?", id).Updates(&modelOrder).Scan(modelOrder).Where("id = ?", id).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error "))
	}

	return modelOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) UpdateStatus(ctx context.Context, id string, order *entities.Order) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderUpdateRepository")
	defer sp.End()

	orderModel := ToOrderModel(order)

	if err := r.db.Connect().Model(&orderModel).Where("id = ?", id).Updates(&orderModel).Scan(orderModel).Preload("Transaction").Where("id = ?", id).First(&orderModel).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error "))
	}
	return orderModel.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Delete(ctx context.Context, id string) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderDeleteRepository")
	defer sp.End()

	order := new(model.Order)
	if err := r.db.Connect().Where("id = ?", id).First(&order).Delete(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to delete order"))
	}

	return order.ToOrderEntity(), nil
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
