package posdetails_service

import "sap-crm/models"

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

// --- Get ID POSDetails ---
func (a *POSDetails) GetByIdPOSDetail_Service(id string) (*models.POSDetails, error) {
	posdetails, err := models.GetByIdPOSDetail_Model(id)
	if err != nil {
		return nil, err
	}

	return posdetails, nil
}

func (a *POSDetails) UpdatePOSDetaiItem_Service(id string) error {

	item := map[string]interface{}{
		"ItemCode": a.ItemCode,
		"ItemName": a.ItemName,
		"Quantity": a.Quantity,
		"Price":    a.Price,
		"Category": a.Category,
		"Tax":      a.Tax,
		"WhsCode":  a.WhsCode,
		// "IDItem":           a.IDItem,
		// "UoMCode":          a.UoMCode,
		// "ExpectedDelivery": a.ExpectedDelivery,
		// "TotalFirst":       a.TotalFirst,
		// "Discount":         a.Discount,
		// "TotalAfter":       a.TotalAfter,
	}

	if err := models.UpdatePOSDetailItem_Model(id, item); err != nil {
		return err
	}

	return nil
}
