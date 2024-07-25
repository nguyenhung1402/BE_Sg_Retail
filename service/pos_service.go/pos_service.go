package pos_service

import (
	"fmt"
	"sap-crm/models"
	"time"
)

type POS struct {
	POSCode     string
	DocDate     time.Time
	Type        string
	PostingDate string
	Total       float64
	Discount    float64
	DocTotal    float64
	VAT         float64
	CustomerPay string
	Refund      float64
	CardCode    string
	CardName    string
	Creator     string
	POSDetails  []POSDetails
	// IDVen        int
	// PONum        string
	// TableName   string
	// TableNumber string
	// Remarks     string
	// Status      bool
	// UoMUnit      string
	// DeliveryDate string
}

type POSDetails struct {
	ItemCode string
	ItemName string
	Quantity float64
	Price    float64
	Category string
	Tax      float64
	WhsCode  string
	// IDItem           int
	// UoMCode          string
	// ExpectedDelivery string
	// TotalFirst       float64
	// Discount         float64
	// TotalAfter       float64
}

// Add POS
func (a *POS) AddPOS() error {
	currentime := time.Now()
	list := make([]map[string]interface{}, len(a.POSDetails))
	for i, v := range a.POSDetails {
		fmt.Printf("v: %v\n", v)
		item := map[string]interface{}{
			"ItemCode": v.ItemCode,
			"ItemName": v.ItemName,
			"Quantity": v.Quantity,
			"Price":    v.Price,
			"Category": v.Category,
			"Tax":      v.Tax,
			"WhsCode":  v.WhsCode,
			// "IDItem":           v.IDItem,
			// "UoMCode":          v.UoMCode,
			// "ExpectedDelivery": v.ExpectedDelivery,
			// "TotalFirst":       v.TotalFirst,
			// "Discount":         v.Discount,
			// "TotalAfter":       v.TotalAfter,
		}
		list[i] = item
	}
	autoSo, _ := models.GetSOIDAuto_Model()
	item := map[string]interface{}{
		"poscode":     autoSo,
		"docdate":     currentime,
		"type":        a.Type,
		"postingdate": a.PostingDate,
		"total":       a.Total,
		"discount":    a.Discount,
		"doctotal":    a.DocTotal,
		"vat":         a.VAT,
		"customerpay": a.CustomerPay,
		"refund":      a.Refund,
		"cardcode":    a.CardCode,
		"cardname":    a.CardName,
		"creator":     a.Creator,
		"posDetails":  a.POSDetails,
		// "idven":        a.IDVen,
		// "ponum":        a.PONum,
		// "tablename":   a.TableName,
		// "tablenumber": a.TableNumber,
		// "remarks":     a.Remarks,
		// "status":      a.Status,
		// "uomunit":      a.UoMUnit,
		// "deliverydate": a.DeliveryDate,
	}
	if err := models.AddPOS(item, list); err != nil {
		return err
	}

	return nil
}

// Update POS
func (a *POS) UpdatePOS(id string) error {

	item := map[string]interface{}{
		"poscode":     a.POSCode,
		"docdate":     a.DocDate,
		"type":        a.Type,
		"postingdate": a.PostingDate,
		"total":       a.Total,
		"discount":    a.Discount,
		"doctotal":    a.DocTotal,
		"vat":         a.VAT,
		"customerpay": a.CustomerPay,
		"cardcode":    a.CardCode,
		"cardname":    a.CardName,
		"refund":      a.Refund,
		"creator":     a.Creator,
		// "idven":        a.IDVen,
		// "ponum":        a.PONum,
		// "tablename":   a.TableName,
		// "tablenumber": a.TableNumber,
		// "remarks":     a.Remarks,
		// "status":      a.Status,
		// "uomunit":      a.UoMUnit,
		// "deliverydate": a.DeliveryDate,
	}

	if err := models.UpdatePOS_Model(id, item); err != nil {
		return err
	}

	return nil
}

// Get all pos
func (a *POS) GetPOS_Service() (*[]models.POS, error) {
	pos, err := models.GetPOS_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return pos, nil
}

// get id pos

func (a *POS) GetByIdPOS_Service(id string) (*models.POS, error) {
	pos, err := models.GetByIdPOS_Model(id)
	if err != nil {
		return nil, err
	}

	return pos, nil
}

// Search POS
func (a *POS) SearchPOS_Service(poscode string) (*[]models.POS, error) {
	pos, err := models.SearchPOS_Model(poscode)
	if err != nil {
		return nil, err
	}
	return pos, nil
}

func (p *POS) GetPOSCode() (*models.POS, error) {
	pos, err := models.GetPOSTableName(p.POSCode)

	if err != nil {
		return nil, err
	}

	return pos, nil
}
