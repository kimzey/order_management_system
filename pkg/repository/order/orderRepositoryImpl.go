package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type orderRepositoryImpl struct {
	db database.Database
}

func NewOrderRepositoryImpl(db database.Database) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) Create(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	_, sp := tracer.Start(ctx, "orderCreateRepository")
	defer sp.End()

	modelOrder := r.ToOrderModel(order)
	if err := r.db.Connect().Create(&modelOrder).Preload("Transaction").Where("id = ?", modelOrder.ID).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("create order error "))
	}

	orderEntity := modelOrder.ToOrderEntity()
	r.SetOrderSubAttributes(orderEntity, sp)
	return orderEntity, nil
}

func (r *orderRepositoryImpl) FindAll(ctx context.Context) (*[]entities.Order, error) {
	_, sp := tracer.Start(ctx, "orderFindAllRepository")
	defer sp.End()

	orders := new([]model.Order)

	if err := r.db.Connect().Preload("Transaction").Find(orders).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find all order error "))
	}
	allOrder := r.ConvertOrderModelsToEntities(orders)

	r.SetOrderSubAttributes(allOrder, sp)
	return allOrder, nil
}
func (r *orderRepositoryImpl) FindByID(ctx context.Context, id string) (*entities.Order, error) {
	_, sp := tracer.Start(ctx, "orderFindByIdRepository")
	defer sp.End()

	order := new(model.Order)
	if err := r.db.Connect().Preload("Transaction").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("find by id order error "))
	}

	orderEntity := order.ToOrderEntity()
	r.SetOrderSubAttributes(order, sp)
	return orderEntity, nil
}

func (r *orderRepositoryImpl) Update(ctx context.Context, id string, order *entities.Order) (*entities.Order, error) {
	_, sp := tracer.Start(ctx, "orderUpdateRepository")
	defer sp.End()

	modelOrder := r.ToOrderModel(order)
	if err := r.db.Connect().Model(&modelOrder).Where("id = ?", id).Updates(&modelOrder).Scan(modelOrder).Where("id = ?", id).First(&modelOrder).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error "))
	}

	orderEntiry := modelOrder.ToOrderEntity()
	r.SetOrderSubAttributes(orderEntiry, sp)
	return orderEntiry, nil
}

func (r *orderRepositoryImpl) UpdateStatus(ctx context.Context, id string, order *entities.Order) (*entities.Order, error) {
	_, sp := tracer.Start(ctx, "orderUpdateRepository")
	defer sp.End()

	orderModel := r.ToOrderModel(order)

	if err := r.db.Connect().Model(&orderModel).Where("id = ?", id).Updates(&orderModel).Scan(orderModel).Preload("Transaction").Where("id = ?", id).First(&orderModel).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("update order error "))
	}

	orderEntiry := orderModel.ToOrderEntity()
	r.SetOrderSubAttributes(orderEntiry, sp)
	return orderEntiry, nil
}

func (r *orderRepositoryImpl) Delete(ctx context.Context, id string) (*entities.Order, error) {
	_, sp := tracer.Start(ctx, "orderDeleteRepository")
	defer sp.End()

	order := new(model.Order)
	if err := r.db.Connect().Where("id = ?", id).First(&order).Delete(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to delete order"))
	}

	orderEntiry := order.ToOrderEntity()
	r.SetOrderSubAttributes(orderEntiry, sp)
	return orderEntiry, nil
}

func (r *orderRepositoryImpl) ConvertOrderModelsToEntities(orders *[]model.Order) *[]entities.Order {
	entityOrders := new([]entities.Order)

	for _, order := range *orders {
		*entityOrders = append(*entityOrders, *order.ToOrderEntity())
	}

	return entityOrders
}
func (r *orderRepositoryImpl) ToOrderModel(e *entities.Order) *model.Order {
	return &model.Order{
		TransactionID: e.TransactionID,
		Status:        e.Status,
	}
}

func (r *orderRepositoryImpl) SetOrderSubAttributes(orderData any, sp trace.Span) {
	if orders, ok := orderData.(*[]entities.Order); ok {
		orderIDs := make([]string, len(*orders))
		transactionIDs := make([]string, len(*orders))
		statuses := make([]string, len(*orders))

		for _, order := range *orders {
			orderIDs = append(orderIDs, order.OrderID)
			transactionIDs = append(transactionIDs, order.TransactionID)
			statuses = append(statuses, order.Status)
		}

		sp.SetAttributes(
			attribute.StringSlice("OrderID", orderIDs),
			attribute.StringSlice("TransactionID", transactionIDs),
			attribute.StringSlice("Status", statuses),
		)
	} else if order, ok := orderData.(*entities.Order); ok {
		sp.SetAttributes(
			attribute.String("OrderID", order.OrderID),
			attribute.String("TransactionID", order.TransactionID),
			attribute.String("Status", order.Status),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
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
