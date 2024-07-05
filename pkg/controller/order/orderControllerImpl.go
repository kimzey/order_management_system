package order

import (
	logger "github.com/kizmey/order_management_system/observability/logs"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_orderService "github.com/kizmey/order_management_system/pkg/service/order"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type orderControllerImpl struct {
	orderService _orderService.OrderService
}

func NewOrderControllerImpl(orderService _orderService.OrderService) OrderController {
	return &orderControllerImpl{orderService}
}

func (c *orderControllerImpl) Create(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "orderCreateController")
	defer sp.End()

	newOrderReq := new(modelReq.Order)
	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(newOrderReq); err != nil {
		logger.LogError("Failed to validate order request", logrus.Fields{"error": err.Error()})
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateOrderRequest)
	}

	order := c.orderReqToEntity(newOrderReq)
	newOrderRes, err := c.orderService.Create(ctx, order)
	if err != nil {
		logger.LogError("Failed to create order", logrus.Fields{"error": err.Error()})
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToCreateOrder)
	}

	logger.LogInfo("Order created successfully", logrus.Fields{"order_id": newOrderRes.OrderID})

	orderRes := c.orderEntityToModelRes(newOrderRes)
	return pctx.JSON(http.StatusCreated, orderRes)
}

func (c *orderControllerImpl) FindAll(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "orderFindAllController")
	defer sp.End()

	orderListingResult, err := c.orderService.FindAll(ctx)
	if err != nil {
		logger.LogError("Failed to retrieve order list", logrus.Fields{"error": err.Error()})
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToRetrieveOrders)
	}

	logger.LogInfo("Retrieved order list successfully", nil)

	allOrder := make([]modelRes.Order, 0)
	for _, order := range *orderListingResult {
		allOrder = append(allOrder, *c.orderEntityToModelRes(&order))
	}

	return pctx.JSON(http.StatusOK, allOrder)
}

func (c *orderControllerImpl) FindByID(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "orderFindByIdController")
	defer sp.End()

	id := pctx.Param("id")

	newOrderRes, err := c.orderService.FindByID(ctx, id)
	if err != nil {
		logger.LogError("Failed to find order by ID", logrus.Fields{"error": err.Error(), "order_id": id})
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrOrderNotFound)
	}

	logger.LogInfo("Found order by ID", logrus.Fields{"order_id": newOrderRes.OrderID})

	orderRes := c.orderEntityToModelRes(newOrderRes)
	return pctx.JSON(http.StatusOK, orderRes)
}

func (c *orderControllerImpl) Update(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "orderUpdateController")
	defer sp.End()

	id := pctx.Param("id")

	orderReq := new(modelReq.Order)
	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(orderReq); err != nil {
		logger.LogError("Failed to validate order update request", logrus.Fields{"error": err.Error()})
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateOrderRequest)
	}

	order := c.orderReqToEntity(orderReq)
	newOrderRes, err := c.orderService.Update(ctx, id, order)
	if err != nil {
		logger.LogError("Failed to update order", logrus.Fields{"error": err.Error(), "order_id": id})
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToUpdateOrder)
	}
	logger.LogInfo("Order updated successfully", logrus.Fields{"order_id": newOrderRes.OrderID})

	orderRes := c.orderEntityToModelRes(newOrderRes)
	return pctx.JSON(http.StatusOK, orderRes)
}

func (c *orderControllerImpl) Delete(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "orderDeleteController")
	defer sp.End()

	id := pctx.Param("id")

	newOrderEntity, err := c.orderService.Delete(ctx, id)
	if err != nil {
		logger.LogError("Failed to delete order", logrus.Fields{"error": err.Error(), "order_id": id})
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToDeleteOrder)
	}

	logger.LogInfo("Order deleted successfully", logrus.Fields{"order_id": newOrderEntity.OrderID})

	orderRes := c.orderEntityToModelRes(newOrderEntity)
	return pctx.JSON(http.StatusOK, orderRes)
}

func (c *orderControllerImpl) ChangeStatusNext(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "orderChangeStatusNextController")
	defer sp.End()

	orderId := pctx.Param("id")
	//if orderId == "" {
	//	err := custom.ErrOrderNotFound
	//	logger.LogWarn("Order id not found", logrus.Fields{"error": err.Error()})
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	newOrderRes, err := c.orderService.ChangeStatusNext(ctx, orderId)
	if err != nil {
		logger.LogError("Failed to change order status", logrus.Fields{"error": err.Error(), "order_id": orderId})
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToChangeOrderStatus)
	}

	logger.LogInfo("Order status changed successfully", logrus.Fields{"order_id": newOrderRes.OrderID, "new_status": newOrderRes.Status})
	orderRes := c.orderEntityToModelRes(newOrderRes)
	return pctx.JSON(http.StatusOK, orderRes)
}

func (c *orderControllerImpl) ChageStatusDone(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "orderChageStatusDoneController")
	defer sp.End()

	id := pctx.Param("id")

	logger.LogInfo("Changing order status to done", logrus.Fields{"order_id": id})
	newOrderRes, err := c.orderService.ChageStatusDone(ctx, id)
	if err != nil {
		logger.LogError("Failed to change order status to done", logrus.Fields{"error": err.Error(), "order_id": id})
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToChangeOrderStatus)
	}

	logger.LogInfo("Order status changed to done successfully", logrus.Fields{"order_id": newOrderRes.OrderID})
	orderRes := c.orderEntityToModelRes(newOrderRes)
	return pctx.JSON(http.StatusOK, orderRes)
}

func (c *orderControllerImpl) orderReqToEntity(orderReq *modelReq.Order) *entities.Order {
	return &entities.Order{
		TransactionID: orderReq.TransactionID,
		Status:        orderReq.Status,
	}
}

func (c *orderControllerImpl) orderEntityToModelRes(order *entities.Order) *modelRes.Order {
	return &modelRes.Order{
		OrderID:       order.OrderID,
		TransactionID: order.TransactionID,
		Status:        order.Status,
	}
}
