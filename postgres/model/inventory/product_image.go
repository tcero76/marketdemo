package inventory

type ProductImage struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	ProductID int  `gorm:"index"`
	URL       string
}

func (ProductImage) TableName() string {
	return "inventory.product_images"
}
