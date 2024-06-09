package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/model"
)

type orderRepositoryImpl struct {
	db database.Database
}

func NewOrderRepositoryImpl(db database.Database) OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) Create(order *entities.Order) (*entities.Order, error) {
	modelOrder := ToOrderModel(order)
	newOrder := new(model.Order)
	//fmt.Println(modelOrder.Transaction.IsDomestic)

	if err := r.db.Connect().Create(modelOrder).Scan(&newOrder).Error; err != nil {
		return nil, err
	}

	if err := r.db.Connect().Preload("Product").Preload("Transaction").First(&newOrder, newOrder.ID).Error; err != nil {
		return nil, err
	}
	//fmt.Println("newOrder: ", newOrder)
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
func (r *orderRepositoryImpl) FindByID(id uint64) (*entities.Order, error) {
	order := new(model.Order)
	if err := r.db.Connect().Preload("Product").Preload("Transaction").First(&order, id).Error; err != nil {
		return nil, err
	}
	return order.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Update(id uint64, order *entities.Order) (*entities.Order, error) {
	orderModel := ToOrderModel(order)

	if err := r.db.Connect().Model(&orderModel).Where("id = ?", id).Updates(&orderModel).Scan(orderModel).Error; err != nil {
		return nil, err

	}

	if err := r.db.Connect().Preload("Product").Preload("Transaction").First(&orderModel, id).Error; err != nil {
		return nil, err
	}

	return orderModel.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) Delete(id uint64) error {
	return r.db.Connect().Delete(&model.Order{}, id).Error

}

func (r *orderRepositoryImpl) ChangeStatusNext(id uint64) (*entities.Order, error) {
	newOrder := new(model.Order)

	if err := r.db.Connect().Preload("Product").Preload("Transaction").First(&newOrder, id).Error; err != nil {
		return nil, err
	}

	if err := newOrder.NextStatus(); err != nil {
		return nil, err
	}

	if err := r.db.Connect().Save(&newOrder).Error; err != nil {
		return nil, err
	}

	return newOrder.ToOrderEntity(), nil
}

func (r *orderRepositoryImpl) ChangeStatusDone(id uint64) (*entities.Order, error) {
	newOrder := new(model.Order)

	if err := r.db.Connect().Preload("Product").Preload("Transaction").First(&newOrder, id).Error; err != nil {
		return nil, err
	}

	if err := newOrder.NextPaidToDone(); err != nil {
		return nil, err
	}

	if err := r.db.Connect().Save(&newOrder).Error; err != nil {
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
	//fmt.Println("e: ", e.IsDomestic)
	return &model.Order{
		TransactionID: e.TransactionID,
		ProductID:     e.ProductID,
		Status:        e.Status,
	}
}
