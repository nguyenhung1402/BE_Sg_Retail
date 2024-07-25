package models

import (
	"time"

	"gorm.io/gorm"
)

type POS struct {
	gorm.Model
	POSCode     string       `gorm:"column:POSCode" json:"poscode"`
	DocDate     time.Time    `gorm:"column:DocDate" json:"docDate"`
	Type        string       `gorm:"column:Type" json:"type"`
	PostingDate string       `gorm:"column:PostingDate" json:"postingdate"`
	Total       float64      `gorm:"column:Total" json:"total"`
	Discount    float64      `gorm:"column:Discount" json:"discout"`
	DocTotal    float64      `gorm:"column:DocTotal" json:"doctotal"`
	VAT         float64      `gorm:"column:VAT" json:"vat"`
	CustomerPay string       `gorm:"column:CustomerPay" json:"customerpay"`
	Refund      float64      `gorm:"column:Refund" json:"refund"`
	CardCode    string       `gorm:"column:CardCode" json:"cardcode"`
	CardName    string       `gorm:"column:CardName" json:"cardname"`
	Creator     string       `gorm:"column:Creator" json:"creator"`
	POSDetails  []POSDetails `gorm:"foreignKey:IDPOS;" json:"posDetails"`
	// IDVen        int         `gorm:"column:IDVen" json:"idven"`
	// PONum        string      `gorm:"column:PONum" json:"ponum"`
	// ObjectType   string      `gorm:"column:ObjectType" json:"objecttype"`
	// TableName   string    `gorm:"column:TableName" json:"tablename"`
	// TableNumber string    `gorm:"column:TableNumber" json:"tablenumber"`
	// Remarks     string    `gorm:"column:Remarks" json:"remarks"`
	// Status      bool      `gorm:"column:Status" json:"status"`
	// UoMUnit      string      `gorm:"column:UoMUnit" json:"uomunit"`
	// DeliveryDate string      `gorm:"column:DeliveryDate" json:"deliverydate"`
}

type IDAuto struct {
	countID2 string `gorm:"column:countID2" json:"countid2"`
}

func AddPOS(data map[string]interface{}, dataDetail []map[string]interface{}) error {
	posDetails := make([]POSDetails, len(dataDetail))
	for i, v := range dataDetail {
		item := POSDetails{
			ItemCode: v["ItemCode"].(string),
			ItemName: v["ItemName"].(string),
			Quantity: v["Quantity"].(float64),
			Price:    v["Price"].(float64),
			Category: v["Category"].(string),
			Tax:      v["Tax"].(float64),
			WhsCode:  v["WhsCode"].(string),
		}
		posDetails[i] = item
	}
	pos := POS{
		POSCode:     data["poscode"].(string),
		DocDate:     data["docdate"].(time.Time),
		Type:        data["type"].(string),
		PostingDate: data["postingdate"].(string),
		Total:       data["total"].(float64),
		Discount:    data["discount"].(float64),
		DocTotal:    data["doctotal"].(float64),
		VAT:         data["vat"].(float64),
		CustomerPay: data["customerpay"].(string),
		Refund:      data["refund"].(float64),
		CardCode:    data["cardcode"].(string),
		CardName:    data["cardname"].(string),
		Creator:     data["creator"].(string),
		POSDetails:  posDetails,
		// IDVen:        data["idven"].(int),
		// PONum:        data["ponum"].(string),
		// ObjectType:   "22",
		// TableName:   data["tablename"].(string),
		// TableNumber: data["tablenumber"].(string),
		// Remarks:     data["remarks"].(string),
		// Status:      data["status"].(bool),
		// UoMUnit:      data["uomunit"].(string),
		// DeliveryDate: data["deliverydate"].(string),
	}
	result := db.Create(&pos)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Edit Member modify a single Member
func UpdatePOS_Model(id string, data map[string]interface{}) error {
	pos := POS{
		POSCode:     data["poscode"].(string),
		DocDate:     data["docdate"].(time.Time),
		Type:        data["type"].(string),
		PostingDate: data["postingdate"].(string),
		Total:       data["total"].(float64),
		Discount:    data["discount"].(float64),
		DocTotal:    data["doctotal"].(float64),
		VAT:         data["vat"].(float64),
		CustomerPay: data["customerpay"].(string),
		Refund:      data["refund"].(float64),
		CardCode:    data["cardcode"].(string),
		CardName:    data["cardname"].(string),
		Creator:     data["creator"].(string),
		// IDVen:        data["idven"].(int),
		// PONum:        data["ponum"].(string),
		// ObjectType:   "22",
		// TableName:   data["tablename"].(string),
		// TableNumber: data["tablenumber"].(string),
		// Remarks:     data["remarks"].(string),
		// Status:      data["status"].(bool),
		// UoMUnit:      data["uomunit"].(string),
		// DeliveryDate: data["deliverydate"].(string),
	}
	if err := db.Model(&pos).Where("id = ?", id).Updates(map[string]interface{}{
		"POSCode":     pos.POSCode,
		"DocDate":     pos.DocDate,
		"Type":        pos.Type,
		"PostingDate": pos.PostingDate,
		"Total":       pos.Total,
		"Discount":    pos.Discount,
		"DocTotal":    pos.DocTotal,
		"VAT":         pos.VAT,
		"CustomerPay": pos.CustomerPay,
		"Refund":      pos.Refund,
		"CardCode":    pos.CardCode,
		"CardName":    pos.CardName,
		// "IDVen":        po.IDVen,
		// "PONum":        po.PONum,
		// "UoMUnit":      po.UoMUnit,
		// "DeliveryDate": po.DeliveryDate,
		"Creator": pos.Creator,
	}).Error; err != nil {
		return err
	}
	return nil
}

// get all pos
func GetPOS_Model() (*[]POS, error) {

	pos := []POS{}
	err := db.Debug().Order("created_at desc").Find(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &pos, nil
}

// Get id po
func GetByIdPOS_Model(id string) (*POS, error) {
	var pos POS

	err := db.Where("id = ?", id).Preload("POSDetails").First(&pos).Error
	// err := db.Raw("select * from sos t0  " +
	// 	" inner join so_details t1 " +
	// 	" on t0.id = t1.IDSO " +
	// 	" where t0.id = " + id + " ",
	// ).Find(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &pos, nil
}

func SearchPOS_Model(poscode string) (*[]POS, error) {
	pos := []POS{}

	err := db.Raw(
		"select * from pos a " +
			"where (charindex(N'" + poscode + "', a.POSCode) > 0) ",
	).Find(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &pos, nil
}

func GetPOSTableName(poscode string) (*POS, error) {
	var pos POS
	err := db.Model(&POS{}).Where("username = ?", poscode).First(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &pos, nil
}

func GetSOIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast(('HD'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from pos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(POSCode) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
