package sales

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	dto "github.com/tcero76/marketplace/bff-service/dto/sales"
	"github.com/tcero76/marketplace/bff-service/services"
	inventoryModel "github.com/tcero76/marketplace/postgres/model/inventory"
	model "github.com/tcero76/marketplace/postgres/model/marketplace"
	salesModel "github.com/tcero76/marketplace/postgres/model/sales"
	schedulingModel "github.com/tcero76/marketplace/postgres/model/scheduling"
)

type SalesService struct {
	db *gorm.DB
}

func NewSalesService(db *gorm.DB) services.ISales {
	return &SalesService{
		db: db,
	}
}

func (s *SalesService) GetOrder(ctx context.Context, req dto.GetOrderRequest) (dto.GetOrderResponse, error) {
	var order salesModel.Order
	err := s.db.WithContext(ctx).
		Preload("Items").
		First(&order, "id = ?", req.OrderID).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.GetOrderResponse{}, ErrOrderNotFound
		}
		return dto.GetOrderResponse{}, err
	}
	items := make([]dto.OrderItemResponse, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, dto.OrderItemResponse{
			ID:            item.ID,
			AppointmentID: item.AppointmentID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			UnitPrice:     item.UnitPrice,
			TotalAmount:   item.TotalAmount,
		})
	}
	return dto.GetOrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		Status:      string(order.Status),
		TotalAmount: order.TotalAmount,
		Currency:    order.Currency,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		Items:       items,
	}, nil
}

func (s *SalesService) CreateOrder(
	ctx context.Context,
	req dto.CreateOrderRequest,
) (dto.CreateOrderResponse, error) {
	if req.UserID == uuid.Nil {
		return dto.CreateOrderResponse{}, ErrInvalidUser
	}
	if err := validateCurrency(req.Currency); err != nil {
		return dto.CreateOrderResponse{}, err
	}
	if len(req.Items) == 0 {
		return dto.CreateOrderResponse{}, ErrEmptyOrder
	}
	quantities := make(map[uuid.UUID]int)
	for _, item := range req.Items {
		if item.ProductID == uuid.Nil {
			return dto.CreateOrderResponse{}, ErrInvalidProduct
		}
		if item.Quantity <= 0 {
			return dto.CreateOrderResponse{}, ErrInvalidQuantity
		}
		if quantities[item.ProductID] > int(^uint(0)>>1)-item.Quantity {
			return dto.CreateOrderResponse{}, ErrInvalidQuantity
		}
		quantities[item.ProductID] += item.Quantity
	}
	productIDs := make([]uuid.UUID, 0, len(quantities))
	for id := range quantities {
		productIDs = append(productIDs, id)
	}
	sort.Slice(productIDs, func(i, j int) bool {
		return productIDs[i].String() < productIDs[j].String()
	})
	var response dto.CreateOrderResponse
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user model.User
		err := tx.
			Where("user_id = ?", req.UserID).
			First(&user).
			Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		if err != nil {
			return err
		}
		var products []inventoryModel.Product
		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id IN ?", productIDs).
			Find(&products).
			Error; err != nil {
			return err
		}
		productMap := make(map[uuid.UUID]inventoryModel.Product)
		for _, product := range products {

			if product.Quantity < quantities[product.ID] {
				return ErrInsufficientStock
			}

			productMap[product.ID] = product

		}
		for _, id := range productIDs {
			product, exists := productMap[id]
			if !exists {
				return ErrProductNotFound
			}

			if !product.Active {
				return ErrProductDisabled
			}
		}
		for _, id := range productIDs {

			quantity := quantities[id]

			result := tx.
				Model(&inventoryModel.Product{}).
				Where(
					"id = ? AND quantity >= ?",
					id,
					quantity,
				).
				Updates(map[string]interface{}{
					"quantity":   gorm.Expr("quantity - ?", quantity),
					"updated_at": time.Now(),
				})

			if result.Error != nil {
				return result.Error
			}

			if result.RowsAffected != 1 {
				return ErrInsufficientStock
			}
		}
		order := salesModel.Order{
			UserID:      req.UserID,
			Status:      salesModel.OrderPending,
			TotalAmount: decimal.Zero,
			Currency:    req.Currency,
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		responseItems := make(
			[]dto.CreateOrderItemResponse,
			0,
			len(req.Items),
		)
		for _, requestedItem := range req.Items {
			product := productMap[requestedItem.ProductID]
			total := product.Price.Mul(
				decimal.NewFromInt(int64(requestedItem.Quantity)),
			)
			item := salesModel.OrderItem{
				OrderID:     order.ID,
				ProductID:   &product.ID,
				Quantity:    requestedItem.Quantity,
				UnitPrice:   product.Price,
				TotalAmount: total,
			}
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
			order.TotalAmount = order.TotalAmount.Add(total)
			responseItems = append(
				responseItems,
				dto.CreateOrderItemResponse{
					OrderItemID: item.ID,
					ProductID:   product.ID,
					Quantity:    item.Quantity,
					UnitPrice:   item.UnitPrice,
					TotalAmount: item.TotalAmount,
				},
			)
		}
		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		movements := make([]inventoryModel.InventoryMovement, 0, len(productIDs))

		for _, id := range productIDs {
			movements = append(movements, inventoryModel.InventoryMovement{
				ProductID:    id,
				UserID:       req.UserID,
				OrderID:      &order.ID,
				MovementType: "SALE",
				Quantity:     quantities[id],
			})
		}

		if err := tx.Create(&movements).Error; err != nil {
			return err
		}

		response = dto.CreateOrderResponse{
			OrderID:     order.ID,
			Status:      string(order.Status),
			TotalAmount: order.TotalAmount,
			Items:       responseItems,
		}
		return nil
	})
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}
	return response, nil
}

func validateCurrency(currency string) error {
	if len(currency) != 3 {
		return ErrInvalidCurrency
	}

	for _, r := range currency {
		if r < 'A' || r > 'Z' {
			return ErrInvalidCurrency
		}
	}

	return nil
}

func (s *SalesService) AddProduct(ctx context.Context, req dto.AddProductRequest) (dto.AddProductResponse, error) {
	if req.Quantity <= 0 {
		return dto.AddProductResponse{}, ErrInvalidQuantity
	}

	var response dto.AddProductResponse

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var order salesModel.Order

		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&order, "id = ?", req.OrderID).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrOrderNotFound
			}

			return err
		}

		if order.Status != salesModel.OrderPending {
			return ErrOrderNotPending
		}

		var product inventoryModel.Product

		err = tx.
			First(&product, "id = ?", req.ProductID).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrProductNotFound
			}

			return err
		}

		if !product.Active {
			return ErrProductDisabled
		}

		total := product.Price.Mul(
			decimal.NewFromInt(int64(req.Quantity)),
		)

		item := salesModel.OrderItem{
			OrderID:     order.ID,
			ProductID:   &product.ID,
			Quantity:    req.Quantity,
			UnitPrice:   product.Price,
			TotalAmount: total,
		}

		if err := tx.Create(&item).Error; err != nil {
			return err
		}

		order.TotalAmount = order.TotalAmount.Add(total)

		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		response = dto.AddProductResponse{
			OrderItemID: item.ID,
			UnitPrice:   item.UnitPrice,
			TotalAmount: item.TotalAmount,
		}

		return nil
	})

	if err != nil {
		return dto.AddProductResponse{}, err
	}

	return response, nil
}

func (s *SalesService) AddAppointment(ctx context.Context, req dto.AddAppointmentRequest) (dto.AddAppointmentResponse, error) {
	var response dto.AddAppointmentResponse
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var order salesModel.Order
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&order, "id = ?", req.OrderID).
			Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrOrderNotFound
			}
			return err
		}
		if order.Status != salesModel.OrderPending {
			return ErrOrderNotPending
		}
		var appointment schedulingModel.Appointment
		err = tx.
			Preload("Resource").
			First(&appointment, "id = ?", req.AppointmentID).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrAppointmentNotFound
			}
			return err
		}
		if appointment.UserID != order.UserID {
			return ErrAppointmentUserMismatch
		}
		if appointment.Status != schedulingModel.AppointmentConfirmed {
			return ErrAppointmentNotAvailable
		}
		price := appointment.Resource.Price
		item := salesModel.OrderItem{
			OrderID:       order.ID,
			AppointmentID: &appointment.ID,
			Quantity:      1,
			UnitPrice:     price,
			TotalAmount:   price,
		}
		if err := tx.Create(&item).Error; err != nil {
			return err
		}
		order.TotalAmount = order.TotalAmount.Add(price)
		if err := tx.Save(&order).Error; err != nil {
			return err
		}
		response = dto.AddAppointmentResponse{
			OrderItemID: item.ID,
			UnitPrice:   item.UnitPrice,
			TotalAmount: item.TotalAmount,
		}
		return nil
	})
	if err != nil {
		return dto.AddAppointmentResponse{}, err
	}
	return response, nil
}

func (s *SalesService) CancelOrder(ctx context.Context, req dto.CancelOrderRequest) (dto.CancelOrderResponse, error) {
	var response dto.CancelOrderResponse
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var order salesModel.Order
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&order, "id = ?", req.OrderID).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrOrderNotFound
			}
			return err
		}
		if order.Status != salesModel.OrderPending {
			return ErrOrderNotPending
		}
		order.Status = salesModel.OrderCancelled
		if err := tx.Save(&order).Error; err != nil {
			return err
		}
		response = dto.CancelOrderResponse{
			OrderID: order.ID,
			Status:  string(order.Status),
		}
		return nil
	})
	if err != nil {
		return dto.CancelOrderResponse{}, err
	}
	return response, nil
}

func (s *SalesService) CreateReturn(ctx context.Context, req dto.CreateReturnRequest) (dto.CreateReturnResponse, error) {
	var response dto.CreateReturnResponse
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var order salesModel.Order
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&order, "id = ?", req.OrderID).
			Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrOrderNotFound
			}
			return err
		}
		if order.Status != salesModel.OrderCompleted {
			return ErrOrderNotCompleted
		}
		if len(req.Items) == 0 {
			return ErrInvalidReturn
		}
		returnOrder := salesModel.ReturnOrder{
			OrderID:     order.ID,
			UserID:      req.UserID,
			Status:      salesModel.ReturnOrderPending,
			TotalAmount: decimal.Zero,
			Currency:    order.Currency,
		}
		if err := tx.Create(&returnOrder).Error; err != nil {
			return err
		}
		for _, reqItem := range req.Items {
			if reqItem.Quantity <= 0 {
				return ErrInvalidQuantity
			}
			var orderItem salesModel.OrderItem
			err := tx.
				First(&orderItem, "id = ? AND order_id = ?", reqItem.OrderItemID, order.ID).
				Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrOrderItemNotFound
				}
				return err
			}
			if reqItem.Quantity > orderItem.Quantity {
				return ErrInvalidReturnQuantity
			}
			unitPrice := orderItem.UnitPrice
			total := unitPrice.Mul(
				decimal.NewFromInt(int64(reqItem.Quantity)),
			)
			returnItem := salesModel.ReturnOrderItem{
				ReturnOrderID: returnOrder.ID,
				OrderItemID:   orderItem.ID,
				Quantity:      reqItem.Quantity,
				UnitPrice:     unitPrice,
				TotalAmount:   total,
			}
			if err := tx.Create(&returnItem).Error; err != nil {
				return err
			}
			returnOrder.TotalAmount = returnOrder.TotalAmount.Add(total)
		}
		if err := tx.Save(&returnOrder).Error; err != nil {
			return err
		}
		response = dto.CreateReturnResponse{
			ReturnOrderID: returnOrder.ID,
			TotalAmount:   returnOrder.TotalAmount,
			Currency:      returnOrder.Currency,
		}
		return nil
	})
	if err != nil {
		return dto.CreateReturnResponse{}, err
	}
	return response, nil
}
