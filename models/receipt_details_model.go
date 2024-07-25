package models

import "gorm.io/gorm"

type ReceiptDetails struct {
	BaseModel
	IDReceipt uint `gorm:"column:IDReceipt" json:"idreceipt"`
	// IDItem           int     `gorm:"column:IDItem" json:"iditem"`
	ItemCode string  `gorm:"column:ItemCode" json:"itemcode"`
	ItemName string  `gorm:"column:ItemName" json:"itemname"`
	Quantity float64 `gorm:"column:Quantity" json:"quantity"`
	// Price    float64 `gorm:"column:Price" json:"price"`
	// Remarks  string  `gorm:"column:Remarks" json:"remarks"`
	// UoMCode          string  `gorm:"column:UoMCode" json:"uomcode"`
	// ExpectedDelivery string  `gorm:"column:ExpectedDelivery" json:"expecteddelivery"`
	// TotalFirst       float64 `gorm:"column:TotalFirst" json:"totalfirst"`
	// Tax              float64 `gorm:"column:Tax" json:"tax"`
	// Discount         float64 `gorm:"column:Discount" json:"discount"`
	// TotalAfter       float64 `gorm:"column:TotalAfter" json:"totalafter"`
}

// --- Tìm 1 hoặc nhiều mã có cùng ID ---
func (u *ReceiptDetails) FindReceiptDetails(receiptid string) (*ReceiptDetails, error) {
	var receiptdetail ReceiptDetails
	err := db.Where("IDReceipt = ?", receiptid).First(&receiptdetail).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &receiptdetail, err
}

// --- get id receipt ---

func GetByIdReceiptDetail_Model(id string) (*ReceiptDetails, error) {
	var receiptDetail ReceiptDetails

	err := db.Where("id = ?", id).First(&receiptDetail).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &receiptDetail, nil
}

// Update Item Receipt
func UpdateReceiptDetailItem_Model(id string, data map[string]interface{}) error {

	item := ReceiptDetails{
		// IDItem:           data["IDItem"].(int),
		ItemCode: data["ItemCode"].(string),
		ItemName: data["ItemName"].(string),
		Quantity: data["Quantity"].(float64),
		// Price:    data["Price"].(float64),
		// UoMCode:          data["UoMCode"].(string),
		// ExpectedDelivery: data["ExpectedDelivery"].(string),
		// TotalFirst:       data["TotalFirst"].(float64),
		// Tax:              data["Tax"].(float64),
		// Discount:         data["Discount"].(float64),
		// TotalAfter:       data["TotalAfter"].(float64),
		// Remarks: data["Remarks"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		// "IDItem":           item.IDItem,
		"ItemCode": item.ItemCode,
		"ItemName": item.ItemName,
		"Quantity": item.Quantity,
		// "Price":    item.Price,
		// "UoMCode":          item.UoMCode,
		// "ExpectedDelivery": item.ExpectedDelivery,
		// "TotalFirst":       item.TotalFirst,
		// "Tax":              item.Tax,
		// "Discount":         item.Discount,
		// "TotalAfter":       item.TotalAfter
		// "Remarks": item.Remarks,
	}).Error; err != nil {
		return err
	}
	return nil
}
