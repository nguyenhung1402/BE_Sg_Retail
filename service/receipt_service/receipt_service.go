package receipt_service

import (
	"fmt"
	"sap-crm/models"
)

type Receipt struct {
	WhsCode        string
	CardCode       string
	CardName       string
	Address        string
	Phone          string
	Creator        string
	AutoCode       string
	ReceiptDetails []ReceiptDetails
}

type ReceiptDetails struct {
	// IDItem           int
	ItemCode string
	ItemName string
	Quantity float64
	// UoMCode          string
	// ExpectedDelivery string
	// TotalFirst       float64
	// Tax              float64
	// Discount         float64
	// TotalAfter       float64
}

// Add Receipt
func (a *Receipt) AddReceipt() error {
	// currentime := time.Now()
	list := make([]map[string]interface{}, len(a.ReceiptDetails))
	for i, v := range a.ReceiptDetails {
		fmt.Printf("v: %v\n", v)
		item := map[string]interface{}{
			"ItemCode": v.ItemCode,
			"ItemName": v.ItemName,
			"Quantity": v.Quantity,
		}
		list[i] = item
	}
	_autoCode, _ := models.GetReceiptIDAuto_Model()
	item := map[string]interface{}{
		"whscode":        a.WhsCode,
		"cardcode":       a.CardCode,
		"cardname":       a.CardName,
		"address":        a.Address,
		"phone":          a.Phone,
		"creator":        a.Creator,
		"autocode":       _autoCode,
		"receiptDetails": a.ReceiptDetails,
	}
	if err := models.AddReceipt(item, list); err != nil {
		return err
	}

	return nil
}

// Update Receipt
func (a *Receipt) UpdateReceipt(id string) error {
	// currentime := time.Now()
	item := map[string]interface{}{
		"whscode":  a.WhsCode,
		"cardcode": a.CardCode,
		"cardname": a.CardName,
		"address":  a.Address,
		"phone":    a.Phone,

		"creator": a.Creator,
	}

	if err := models.UpdateReceipt_Model(id, item); err != nil {
		return err
	}

	return nil
}

// Get all receipt
func (a *Receipt) GetReceipt_Service() (*[]models.Receipt, error) {
	receipt, err := models.GetReceipt_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return receipt, nil
}

// get id receipt

func (a *Receipt) GetByIdReceipt_Service(id string) (*models.Receipt, error) {
	receipt, err := models.GetByIdReceipt_Model(id)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

// Search Receipt
func (a *Receipt) SearchReceipt_Service(whscode string, cardcode string) (*[]models.Receipt, error) {
	receipt, err := models.SearchReceipt_Model(whscode, cardcode)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (p *Receipt) GetReceiptID() (*models.Receipt, error) {
	receipt, err := models.GetReceiptTableName(p.WhsCode)

	if err != nil {
		return nil, err
	}

	return receipt, nil
}
