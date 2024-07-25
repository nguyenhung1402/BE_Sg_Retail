package receiptdetails_service

import "sap-crm/models"

type ReceiptDetails struct {
	// IDItem           int
	ItemCode string
	ItemName string
	Quantity float64
}

func (a *ReceiptDetails) GetByIdReceiptDetail_Service(id string) (*models.ReceiptDetails, error) {
	receiptdetails, err := models.GetByIdReceiptDetail_Model(id)
	if err != nil {
		return nil, err
	}

	return receiptdetails, nil
}

// --- Update Item ReceiptDetails ---
func (a *ReceiptDetails) UpdateReceiptDetaiItem_Service(id string) error {

	item := map[string]interface{}{
		// "IDItem":           a.IDItem,
		"ItemCode": a.ItemCode,
		"ItemName": a.ItemName,
		"Quantity": a.Quantity,

		// "UoMCode":          a.UoMCode,
		// "ExpectedDelivery": a.ExpectedDelivery,
		// "TotalFirst":       a.TotalFirst,
		// "Tax":              a.Tax,
		// "Discount":         a.Discount,
		// "TotalAfter":       a.TotalAfter,
	}

	if err := models.UpdateReceiptDetailItem_Model(id, item); err != nil {
		return err
	}

	return nil
}
