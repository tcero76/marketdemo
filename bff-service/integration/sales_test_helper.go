package integration

import (
	"testing"

	"github.com/google/uuid"
	inventoryModel "github.com/tcero76/marketplace/postgres/model/inventory"
	model "github.com/tcero76/marketplace/postgres/model/sales"
	"gorm.io/gorm"
)

func deleteOrder(t *testing.T, db *gorm.DB, orderID uuid.UUID) error {
	t.Helper()
	if err := db.
		Where("order_id = ?", orderID).
		Delete(&model.OrderItem{}).Error; err != nil {
		return err
	}
	if err := db.
		Where("order_id = ?", orderID).
		Delete(&inventoryModel.InventoryMovement{}).Error; err != nil {
		return err
	}
	if err := db.
		Where("id = ?", orderID).
		Delete(&model.Order{}).Error; err != nil {
		return err
	}
	return nil
}
