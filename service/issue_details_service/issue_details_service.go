package receiptdetails_service

import "sap-crm/models"

type IssueDetails struct {
	// IDItem           int
	ItemCode string
	ItemName string
	Quantity float64
}

func (a *IssueDetails) GetByIdIssueDetail_Service(id string) (*models.IssueDetails, error) {
	issuedetails, err := models.GetByIdIssueDetail_Model(id)
	if err != nil {
		return nil, err
	}

	return issuedetails, nil
}

// --- Update Item IssueDetails ---
func (a *IssueDetails) UpdateIssueDetaiItem_Service(id string) error {

	issue := map[string]interface{}{
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

	if err := models.UpdateIssueDetailItem_Model(id, issue); err != nil {
		return err
	}

	return nil
}
