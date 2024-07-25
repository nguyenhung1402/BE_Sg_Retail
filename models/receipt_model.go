package models

import (
	// "time"

	"fmt"
	"time"

	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	WhsCode        string           `gorm:"column:WhsCode" json:"whscode"`
	CardCode       string           `gorm:"column:CardCode" json:"cardcode"`
	CardName       string           `gorm:"column:CardName" json:"cardname"`
	Address        string           `gorm:"column:Address" json:"address"`
	Phone          string           `gorm:"column:Phone" json:"phone"`
	Creator        string           `gorm:"column:Creator" json:"creator"`
	AutoCode       string           `gorm:"column:AutoCode" json:"autocode"`
	ReceiptDetails []ReceiptDetails `gorm:"foreignKey:IDReceipt;" json:"receiptDetails"`
}

func AddReceipt(data map[string]interface{}, dataDetail []map[string]interface{}) error {
	receiptDetails := make([]ReceiptDetails, len(dataDetail))
	for i, v := range dataDetail {
		item := ReceiptDetails{
			ItemCode: v["ItemCode"].(string),
			ItemName: v["ItemName"].(string),
			Quantity: v["Quantity"].(float64),
		}
		receiptDetails[i] = item
	}
	receipt := Receipt{
		WhsCode:        data["whscode"].(string),
		CardCode:       data["cardcode"].(string),
		CardName:       data["cardname"].(string),
		Address:        data["address"].(string),
		Phone:          data["phone"].(string),
		Creator:        data["creator"].(string),
		AutoCode:       data["autocode"].(string),
		ReceiptDetails: receiptDetails,
	}
	result := db.Create(&receipt)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Edit Member modify a single Member
func UpdateReceipt_Model(id string, data map[string]interface{}) error {
	receipt := Receipt{

		WhsCode:  data["whscode"].(string),
		CardCode: data["cardcode"].(string),
		CardName: data["cardname"].(string),
		Address:  data["address"].(string),
		Phone:    data["phone"].(string),
		Creator:  data["creator"].(string),
	}
	if err := db.Model(&receipt).Where("id = ?", id).Updates(map[string]interface{}{
		"WhsCode":  receipt.WhsCode,
		"CardCode": receipt.CardCode,
		"CardName": receipt.CardName,
		"Address":  receipt.Address,
		"Phone":    receipt.Phone,
		"Creator":  receipt.Creator,
	}).Error; err != nil {
		return err
	}
	return nil
}

// get all po
func GetReceipt_Model() (*[]Receipt, error) {
	currentTime := time.Now()
	var abc = currentTime.Format("01-02-2006")
	fmt.Println(abc)
	receipt := []Receipt{}
	err := db.Debug().Where("cast(created_at as date) = ?", abc).Order("created_at desc").Find(&receipt).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &receipt, nil
}

// Get id po
func GetByIdReceipt_Model(id string) (*Receipt, error) {
	var receipt Receipt

	err := db.Where("id = ?", id).Preload("ReceiptDetails").First(&receipt).Error
	// err := db.Raw("select * from sos t0  " +
	// 	" inner join so_details t1 " +
	// 	" on t0.id = t1.IDSO " +
	// 	" where t0.id = " + id + " ",
	// ).Find(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &receipt, nil
}

func SearchReceipt_Model(whscode string, cardcode string) (*[]Receipt, error) {
	receipt := []Receipt{}

	err := db.Raw(
		"select * from receipt a " +
			"where (charindex(N'" + whscode + "', a.TableName) > 0) " +
			"union " +
			"select * from receipt a " +
			"where (charindex(N'" + cardcode + "', a.TableNumber) > 0) ",
	).Find(&receipt).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &receipt, nil
}

func GetReceiptTableName(tablename string) (*Receipt, error) {
	var receipt Receipt
	err := db.Model(&Receipt{}).Where("username = ?", tablename).First(&receipt).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &receipt, nil
}

func GetReceiptIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast(('NK'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),created_at,12),4) ORDER BY ID)) countID from receipts where (MONTH(created_at) = MONTH(getdate()) and YEAR(created_at) = YEAR(GETDATE())) and LEN(created_at) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
