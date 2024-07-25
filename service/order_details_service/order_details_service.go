package orderdetails_service

import (
	"fmt"
	"sap-crm/models"
)

type OrderDetails struct {
	ItemCode string
	ItemName string
	Quantity float64
	Price    float64
	Tax      float64
	Remarks  string
	WhsCode  string
	Status   string
	// IDItem           int
	// UoMCode          string
	// ExpectedDelivery string
	// TotalFirst       float64
	// Discount         float64
	// TotalAfter       float64
}

type OrderDetailsAddNew struct {
	IDOrder  uint
	ItemCode string
	ItemName string
	Quantity float64
	Price    float64
	Tax      float64
	Remarks  string
	WhsCode  string
	Status   string
	// IDItem           int
	// UoMCode          string
	// ExpectedDelivery string
	// TotalFirst       float64
	// Discount         float64
	// TotalAfter       float64
}

// --- Get ID OrderDetails ---
func (a *OrderDetails) GetByIdOrderDetail_Service(id string) (*models.OrderDetails, error) {
	orderdetails, err := models.GetByIdOrderDetail_Model(id)
	if err != nil {
		return nil, err
	}

	return orderdetails, nil
}

func (a *OrderDetailsAddNew) AddOrderDetails_Service(id string) error {
	// theTime := currentime.Format("2006-1-2 15:4:5")

	fmt.Println("Service")
	fmt.Println(a.ItemCode)
	fmt.Println(a.ItemName)
	fmt.Println(a.Quantity)
	fmt.Println(a.Price)
	fmt.Println(a.Remarks)
	fmt.Println(a.Tax)
	fmt.Println(a.Tax)
	fmt.Println(a.WhsCode)
	fmt.Println(a.Status)
	item := map[string]interface{}{
		"idorder":  a.IDOrder,
		"itemcode": a.ItemCode,
		"itemname": a.ItemName,
		"quantity": a.Quantity,
		"price":    a.Price,
		"remarks":  a.Remarks,
		"tax":      a.Tax,
		"whscode":  a.WhsCode,
		"status":   a.Status,
	}
	fmt.Println(item)
	if err := models.AddOrderDetailItem_Model(item, id); err != nil {
		return err
	}

	return nil
}

// --- Update Item OrderDetails ---
func (a *OrderDetails) UpdateOrderDetaiItem_Service(id string) error {

	item := map[string]interface{}{
		"ItemCode": a.ItemCode,
		"ItemName": a.ItemName,
		"Quantity": a.Quantity,
		"Price":    a.Price,
		"Remarks":  a.Remarks,
		"Tax":      a.Tax,
		"WhsCode":  a.WhsCode,
		"Status":   a.Status,
		// "IDItem":           a.IDItem,
		// "UoMCode":          a.UoMCode,
		// "ExpectedDelivery": a.ExpectedDelivery,
		// "TotalFirst":       a.TotalFirst,
		// "Discount":         a.Discount,
		// "TotalAfter":       a.TotalAfter,
	}

	if err := models.UpdateOrderDetailItem_Model(id, item); err != nil {
		return err
	}

	return nil
}
