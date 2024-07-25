package models

import "gorm.io/gorm"

type IssueDetails struct {
	BaseModel
	IDIssue  uint    `gorm:"column:IDIssue" json:"idissue"`
	ItemCode string  `gorm:"column:ItemCode" json:"itemcode"`
	ItemName string  `gorm:"column:ItemName" json:"itemname"`
	Quantity float64 `gorm:"column:Quantity" json:"quantity"`
}

// --- Tìm 1 hoặc nhiều mã có cùng ID ---
func (u *IssueDetails) FindIssueDetails(issueid string) (*IssueDetails, error) {
	var issuedetail IssueDetails
	err := db.Where("IDIssue = ?", issueid).First(&issuedetail).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &issuedetail, err
}

// --- get id issue ---

func GetByIdIssueDetail_Model(id string) (*IssueDetails, error) {
	var issueDetail IssueDetails

	err := db.Where("id = ?", id).First(&issueDetail).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &issueDetail, nil
}

// Update Item Issue
func UpdateIssueDetailItem_Model(id string, data map[string]interface{}) error {

	item := IssueDetails{

		ItemCode: data["ItemCode"].(string),
		ItemName: data["ItemName"].(string),
		Quantity: data["Quantity"].(float64),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		// "IDItem":           item.IDItem,
		"ItemCode": item.ItemCode,
		"ItemName": item.ItemName,
		"Quantity": item.Quantity,
	}).Error; err != nil {
		return err
	}
	return nil
}
