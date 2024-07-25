package issue_service

import (
	"fmt"
	"sap-crm/models"
)

type Issue struct {
	WhsCode      string
	CardCode     string
	CardName     string
	Address      string
	Phone        string
	Creator      string
	AutoCode     string
	IssueDetails []IssueDetails
}

type IssueDetails struct {
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

// Add Issue
func (a *Issue) AddIssue() error {
	// currentime := time.Now()
	list := make([]map[string]interface{}, len(a.IssueDetails))
	for i, v := range a.IssueDetails {
		fmt.Printf("v: %v\n", v)
		issue := map[string]interface{}{
			"ItemCode": v.ItemCode,
			"ItemName": v.ItemName,
			"Quantity": v.Quantity,
		}
		list[i] = issue
	}
	_autoCode, _ := models.GetIssueIDAuto_Model()
	issue := map[string]interface{}{
		"whscode":      a.WhsCode,
		"cardcode":     a.CardCode,
		"cardname":     a.CardName,
		"address":      a.Address,
		"phone":        a.Phone,
		"autocode":     _autoCode,
		"creator":      a.Creator,
		"issueDetails": a.IssueDetails,
	}
	if err := models.AddIssue(issue, list); err != nil {
		return err
	}

	return nil
}

// Update Issue
func (a *Issue) UpdateIssue(id string) error {
	// currentime := time.Now()
	item := map[string]interface{}{
		"whscode":  a.WhsCode,
		"cardcode": a.CardCode,
		"cardname": a.CardName,
		"address":  a.Address,
		"phone":    a.Phone,

		"creator": a.Creator,
	}

	if err := models.UpdateIssue_Model(id, item); err != nil {
		return err
	}

	return nil
}

// Get all issue
func (a *Issue) GetIssue_Service() (*[]models.Issue, error) {
	issue, err := models.GetIssue_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return issue, nil
}

// get id issue

func (a *Issue) GetByIdIssue_Service(id string) (*models.Issue, error) {
	issue, err := models.GetByIdIssue_Model(id)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// Search issue
func (a *Issue) SearchIssue_Service(whscode string, cardcode string) (*[]models.Issue, error) {
	issue, err := models.SearchIssue_Model(whscode, cardcode)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (p *Issue) GetIssueID() (*models.Issue, error) {
	issue, err := models.GetIssueTableName(p.WhsCode)

	if err != nil {
		return nil, err
	}

	return issue, nil
}
