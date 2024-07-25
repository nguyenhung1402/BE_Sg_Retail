package models

import "gorm.io/gorm"

type POSDetails struct {
	BaseModel
	IDPOS    uint    `gorm:"column:IDPOS" json:"idpos"`
	ItemCode string  `gorm:"column:ItemCode" json:"itemcode"`
	ItemName string  `gorm:"column:ItemName" json:"itemname"`
	Quantity float64 `gorm:"column:Quantity" json:"quantity"`
	Price    float64 `gorm:"column:Price" json:"price"`
	Category string  `gorm:"column:Category" json:"category"`
	Tax      float64 `gorm:"column:Tax" json:"tax"`
	WhsCode  string  `gorm:"column:WhsCode" json:"whscode"`
	// IDItem           int     `gorm:"column:IDItem" json:"iditem"`
	// UoMCode          string  `gorm:"column:UoMCode" json:"uomcode"`
	// ExpectedDelivery string  `gorm:"column:ExpectedDelivery" json:"expecteddelivery"`
	// TotalFirst       float64 `gorm:"column:TotalFirst" json:"totalfirst"`
	// Discount         float64 `gorm:"column:Discount" json:"discount"`
	// TotalAfter       float64 `gorm:"column:TotalAfter" json:"totalafter"`
}

// --- Tìm 1 hoặc nhiều mã có cùng ID ---
func (u *POSDetails) FindPOSDetails(posid string) (*POSDetails, error) {
	var posdetail POSDetails
	err := db.Where("IDPOS = ?", posid).First(&posdetail).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &posdetail, err
}

// --- get id POS ---

func GetByIdPOSDetail_Model(id string) (*POSDetails, error) {
	var posDetail POSDetails

	err := db.Where("id = ?", id).First(&posDetail).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &posDetail, nil
}

// Update Item POS
func UpdatePOSDetailItem_Model(id string, data map[string]interface{}) error {

	item := POSDetails{
		ItemCode: data["ItemCode"].(string),
		ItemName: data["ItemName"].(string),
		Quantity: data["Quantity"].(float64),
		Price:    data["Price"].(float64),
		Category: data["Category"].(string),
		Tax:      data["Tax"].(float64),
		WhsCode:  data["WhsCode"].(string),
		// IDItem:           data["IDItem"].(int),
		// UoMCode:          data["UoMCode"].(string),
		// ExpectedDelivery: data["ExpectedDelivery"].(string),
		// TotalFirst:       data["TotalFirst"].(float64),
		// Discount:         data["Discount"].(float64),
		// TotalAfter:       data["TotalAfter"].(float64),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"ItemCode": item.ItemCode,
		"ItemName": item.ItemName,
		"Quantity": item.Quantity,
		"Price":    item.Price,
		"Category": item.Category,
		"Tax":      item.Tax,
		"WhsCode":  item.WhsCode,
		// "IDItem":           item.IDItem,
		// "UoMCode":          item.UoMCode,
		// "ExpectedDelivery": item.ExpectedDelivery,
		// "TotalFirst":       item.TotalFirst,
		// "Discount":         item.Discount,
		// "TotalAfter":       item.TotalAfter
	}).Error; err != nil {
		return err
	}
	return nil
}
