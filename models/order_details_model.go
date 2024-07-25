package models

import (
	"fmt"

	"gorm.io/gorm"
)

type OrderDetails struct {
	BaseModel
	IDOrder  uint    `gorm:"column:IDOrder" json:"idorder"`
	ItemCode string  `gorm:"column:ItemCode" json:"itemcode"`
	ItemName string  `gorm:"column:ItemName" json:"itemname"`
	Quantity float64 `gorm:"column:Quantity" json:"quantity"`
	Price    float64 `gorm:"column:Price" json:"price"`
	Remarks  string  `gorm:"column:Remarks" json:"remarks"`
	Tax      float64 `gorm:"column:Tax" json:"tax"`
	WhsCode  string  `gorm:"column:WhsCode" json:"whscode"`
	Status   string  `gorm:"column:Status" json:"status"`
	// IDItem           int     `gorm:"column:IDItem" json:"iditem"`
	// UoMCode          string  `gorm:"column:UoMCode" json:"uomcode"`
	// ExpectedDelivery string  `gorm:"column:ExpectedDelivery" json:"expecteddelivery"`
	// TotalFirst       float64 `gorm:"column:TotalFirst" json:"totalfirst"`
	// Discount         float64 `gorm:"column:Discount" json:"discount"`
	// TotalAfter       float64 `gorm:"column:TotalAfter" json:"totalafter"`
}

type OrderDetailsAddNew struct {
	IDOrder  uint    `gorm:"column:IDOrder" json:"idorder"`
	ItemCode string  `gorm:"column:ItemCode" json:"itemcode"`
	ItemName string  `gorm:"column:ItemName" json:"itemname"`
	Quantity float64 `gorm:"column:Quantity" json:"quantity"`
	Price    float64 `gorm:"column:Price" json:"price"`
	Remarks  string  `gorm:"column:Remarks" json:"remarks"`
	Tax      float64 `gorm:"column:Tax" json:"tax"`
	WhsCode  string  `gorm:"column:WhsCode" json:"whscode"`
	Status   string  `gorm:"column:Status" json:"status"`
}

// --- Tìm 1 hoặc nhiều mã có cùng ID ---
func (u *OrderDetails) FindOrderDetails(orderid string) (*OrderDetails, error) {
	var orderdetail OrderDetails
	err := db.Where("IDOrder = ?", orderid).First(&orderdetail).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &orderdetail, err
}

// --- get id Order ---

func GetByIdOrderDetail_Model(id string) (*OrderDetails, error) {
	var orderDetail OrderDetails

	err := db.Where("id = ?", id).First(&orderDetail).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &orderDetail, nil
}

func AddOrderDetailItem_Model(data map[string]interface{}, id string) error {
	// var order Order
	fmt.Println("Model")
	fmt.Println(data["itemcode"].(string))
	fmt.Println(data["itemname"].(string))
	fmt.Println(data["quantity"].(float64))
	fmt.Println(data["price"].(float64))
	fmt.Println(data["tax"].(float64))
	fmt.Println(data["remarks"].(string))
	fmt.Println(data["whscode"].(string))
	fmt.Println(data["status"].(string))
	item := OrderDetails{
		IDOrder:  data["idorder"].(uint),
		ItemCode: data["itemcode"].(string),
		ItemName: data["itemname"].(string),
		Quantity: data["quantity"].(float64),
		Price:    data["price"].(float64),
		Tax:      data["tax"].(float64),
		Remarks:  data["remarks"].(string),
		WhsCode:  data["whscode"].(string),
		Status:   data["status"].(string),
	}

	//result := db.Debug().Find(&item)
	// result := db.Create(&item)
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
	// if err := db.Model(&order).Where("id = ?", id).Association("order_details").Append(&OrderDetails{
	// 	ItemCode: item.ItemCode,
	// 	ItemName: item.ItemName,
	// 	Quantity: item.Quantity,
	// 	Price:    item.Price,
	// 	Tax:      item.Tax,
	// 	Remarks:  item.Remarks,
	// 	WhsCode:  item.WhsCode,
	// 	Status:   item.Status,
	// }); err != nil {
	// 	return err
	// }
	// return nil
}

// Update Item Order
func UpdateOrderDetailItem_Model(id string, data map[string]interface{}) error {

	item := OrderDetails{
		ItemCode: data["ItemCode"].(string),
		ItemName: data["ItemName"].(string),
		Quantity: data["Quantity"].(float64),
		Price:    data["Price"].(float64),
		Tax:      data["Tax"].(float64),
		Remarks:  data["Remarks"].(string),
		WhsCode:  data["WhsCode"].(string),
		Status:   data["Status"].(string),
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
		"Tax":      item.Tax,
		"Remarks":  item.Remarks,
		"WhsCode":  item.WhsCode,
		"Status":   item.Status,
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
