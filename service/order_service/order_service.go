package order_service

import (
	"fmt"
	"sap-crm/models"
	"time"
)

type Order struct {
	TableName    string
	TableNumber  string
	DocDate      time.Time
	Remarks      string
	Status       bool
	Type         string
	PostingDate  string
	Total        float64
	Discount     int
	DocTotal     float64
	Customerpay  string
	Refund       float64
	CardCode     string
	CardName     string
	VAT          int
	Creator      string
	ViewPayment  string
	CodeAuto     string
	OrderDetails []OrderDetails
	// IDVen        int
	// PONum        string
	// UoMUnit      string
	// DeliveryDate string
}

type OrderDetails struct {
	ItemCode string
	ItemName string
	Quantity float64
	Price    float64
	Remarks  string
	Tax      float64
	WhsCode  string
	Status   string
	// IDItem           int
	// UoMCode          string
	// ExpectedDelivery string
	// TotalFirst       float64
	// Discount         float64
	// TotalAfter       float64
}

// Add Order
func (a *Order) AddOrder() error {
	currentime := time.Now()
	list := make([]map[string]interface{}, len(a.OrderDetails))
	for i, v := range a.OrderDetails {
		fmt.Printf("v: %v\n", v)
		item := map[string]interface{}{
			"ItemCode": v.ItemCode,
			"ItemName": v.ItemName,
			"Quantity": v.Quantity,
			"Price":    v.Price,
			"Remarks":  v.Remarks,
			"Tax":      v.Tax,
			"WhsCode":  v.WhsCode,
			"Status":   "save",
			// "IDItem":           v.IDItem,
			// "UoMCode":          v.UoMCode,
			// "ExpectedDelivery": v.ExpectedDelivery,
			// "TotalFirst":       v.TotalFirst,
			// "Discount":         v.Discount,
			// "TotalAfter":       v.TotalAfter,
		}
		list[i] = item
	}
	autoOrder, _ := models.GetORDERIDAuto_Model()
	item := map[string]interface{}{
		"tablename":    a.TableName,
		"tablenumber":  a.TableNumber,
		"docdate":      currentime,
		"remarks":      a.Remarks,
		"status":       a.Status,
		"type":         a.Type,
		"postingdate":  a.PostingDate,
		"total":        a.Total,
		"discount":     a.Discount,
		"doctotal":     a.DocTotal,
		"customerpay":  a.Customerpay,
		"refund":       a.Refund,
		"cardcode":     a.CardCode,
		"cardname":     a.CardName,
		"vat":          a.VAT,
		"creator":      a.Creator,
		"viewpayment":  a.ViewPayment,
		"codeauto":     autoOrder,
		"orderDetails": a.OrderDetails,
		// "idven":        a.IDVen,
		// "ponum":        a.PONum,
		// "uomunit":      a.UoMUnit,
		// "deliverydate": a.DeliveryDate,
	}
	if err := models.AddOrder(item, list); err != nil {
		return err
	}

	return nil
}

// Update Order
func (a *Order) UpdateOrder(id string) error {
	currentime := time.Now()
	item := map[string]interface{}{
		"tablename":   a.TableName,
		"tablenumber": a.TableNumber,
		"docdate":     currentime,
		"remarks":     a.Remarks,
		"status":      a.Status,
		"type":        a.Type,
		"postingdate": a.PostingDate,
		"total":       a.Total,
		"discount":    a.Discount,
		"doctotal":    a.DocTotal,
		"customerpay": a.Customerpay,
		"refund":      a.Refund,
		"cardcode":    a.CardCode,
		"cardname":    a.CardName,
		"vat":         a.VAT,
		"creator":     a.Creator,
		"viewpayment": a.ViewPayment,
		// "idven":        a.IDVen,
		// "ponum":        a.PONum,
		// "uomunit":      a.UoMUnit,
		// "deliverydate": a.DeliveryDate,
		// "orderDetails":    a.OrderDetails,
	}

	if err := models.UpdateOrder_Model(id, item); err != nil {
		return err
	}

	return nil
}

// Get all order
func (a *Order) GetOrder_Service() (*[]models.Order, error) {
	order, err := models.GetOrder_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return order, nil
}
func (a *Order) PostGetOrder_Service(day string, month string, year string) (*[]models.Order, error) {
	order, err := models.PostGetOrder_Model(day, month, year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return order, nil
}

// get id order

func (a *Order) GetByIdOrder_Service(id string) (*models.Order, error) {
	order, err := models.GetByIdOrder_Model(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// Search Order
func (a *Order) SearchOrder_Service(tablename string, tablenumber string) (*[]models.Order, error) {
	order, err := models.SearchOrder_Model(tablename, tablenumber)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (p *Order) GetOrderTableName() (*models.Order, error) {
	order, err := models.GetOrderTableName(p.TableName)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (a *Order) AppendOrderDetails_Service(id string) error {
	currentime := time.Now()
	list := make([]map[string]interface{}, len(a.OrderDetails))
	for i, v := range a.OrderDetails {
		fmt.Printf("v: %v\n", v)
		item := map[string]interface{}{
			"ItemCode": v.ItemCode,
			"ItemName": v.ItemName,
			"Quantity": v.Quantity,
			"Price":    v.Price,
			"Remarks":  v.Remarks,
			"Tax":      v.Tax,
			"WhsCode":  v.WhsCode,
			"Status":   "save",
		}
		list[i] = item
	}
	item := map[string]interface{}{
		"tablename":    a.TableName,
		"tablenumber":  a.TableNumber,
		"docdate":      currentime,
		"remarks":      a.Remarks,
		"status":       a.Status,
		"type":         a.Type,
		"postingdate":  a.PostingDate,
		"total":        a.Total,
		"discount":     a.Discount,
		"doctotal":     a.DocTotal,
		"customerpay":  a.Customerpay,
		"refund":       a.Refund,
		"cardcode":     a.CardCode,
		"cardname":     a.CardName,
		"vat":          a.VAT,
		"creator":      a.Creator,
		"viewpayment":  "save",
		"orderDetails": a.OrderDetails,
	}
	fmt.Println(item)
	if err := models.AppendOrderDetails_Model(item, list, id); err != nil {
		return err
	}

	return nil
}

func (p *Order) GetTablenumber_Service(code string, date string) (*[]models.Order, error) {
	order, err := models.GetTablenumber(code, date)

	if err != nil {
		return nil, err
	}

	return order, nil
}

// Get all order
func (a *Order) GetTablenumberNotThanhToan_Service(code string, date string) (*[]models.Order, error) {
	order, err := models.GetTablenumberNotThanhToan(code, date)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return order, nil
}
