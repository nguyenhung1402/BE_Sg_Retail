package models

import (
	"gorm.io/gorm"
)

type ViewNXT struct {
	WhsCode  string `gorm:"column:WhsCode" json:"whscode"`
	ItemCode string `gorm:"column:ItemCode" json:"itemcode"`
	Quantity string `gorm:"column:Quantity" json:"quantity"`
}

func GetAllNXT_Model() (*[]ViewNXT, error) {
	viewNXT := []ViewNXT{}
	err := db.Raw("select WhsCode,ItemCode, Quantity from ViewNXT_RETAILS_ONE").First(&viewNXT).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &viewNXT, nil
}

func GetWhsCodeNXT_Model(whscode string) (*[]ViewNXT, error) {
	viewNXT := []ViewNXT{}
	err := db.Raw("select WhsCode,ItemCode, Quantity from ViewNXT_RETAILS_ONE where WhsCode = ?", whscode).First(&viewNXT).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &viewNXT, nil
}

func GetWhsItemCodeNXT_Model(whscode string, itemcode string) (*ViewNXT, error) {
	viewNXT := ViewNXT{}
	err := db.Raw("select WhsCode,ItemCode, Quantity from ViewNXT_RETAILS_ONE where WhsCode = ? and ItemCode = ?", whscode, itemcode).First(&viewNXT).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &viewNXT, nil
}

// CREATE VIEW ViewNXT_RETAILS_ONE AS
// select A.WhsCode, A.ItemCode, SUM(A.Quantity) Quantity from (
// select t0.WhsCode, t1.ItemCode, SUM(t1.Quantity) Quantity from receipts t0 inner join receipt_details t1 on t0.id = t1.IDReceipt
// where (t0.WhsCode is not null)
// group by t0.WhsCode, t1.ItemCode
// union all
// select t0.WhsCode, t1.ItemCode, -SUM(t1.Quantity) Quantity from receipts t0 inner join issue_details t1 on t0.id = t1.IDIssue
// where (t0.WhsCode is not null)
// group by t0.WhsCode, t1.ItemCode
// union all
// select t0.WhsCode, t0.ItemCode, -SUM(t0.Quantity) Quantity from order_details t0
// where (t0.WhsCode is not null)
// group by t0.WhsCode, t0.ItemCode
// ) A
// --where A.WhsCode = 'DEMO02' and A.ItemCode = ''
// group by A.WhsCode, A.ItemCode
