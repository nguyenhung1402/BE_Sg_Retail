package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	ItemCode     string    `gorm:"column:ItemCode" json:"itemcode"`
	ItemName     string    `gorm:"column:ItemName" json:"itemname"`
	Quantity     int       `gorm:"column:Quantity" json:"quantity"`
	Group        string    `gorm:"column:Group" json:"group"`
	InStock      float64   `gorm:"column:InStock" json:"instock"`
	Price        float64   `gorm:"column:Price" json:"price"`
	DocDate      time.Time `gorm:"column:DocDate" json:"docdate"`
	UserAdd      string    `gorm:"column:UserAdd" json:"useradd"`
	UserUpdate   string    `gorm:"column:UserUpdate" json:"userupdate"`
	Status       bool      `gorm:"column:Status" json:"status"`
	Type         string    `gorm:"column:Type" json:"type"`
	Barcode      string    `gorm:"column:Barcode" json:"barcode"`
	Image        string    `gorm:"column:Image" json:"image"`
	WhsCode      string    `gorm:"column: WhsCode" json:"whscode"`
	CardCode     string    `gorm:"column: CardCode" json:"cardcode"`
	Bonus        float64   `gorm:"column: Bonus" json:"bonus"`
	IdentifiCode string    `gorm:"column: IdentifiCode" json:"identificode"`
}

func AddItems(data map[string]interface{}) error {
	item := Items{
		ItemCode:     data["itemcode"].(string),
		ItemName:     data["itemname"].(string),
		Quantity:     data["quantity"].(int),
		Group:        data["group"].(string),
		InStock:      data["instock"].(float64),
		Price:        data["price"].(float64),
		DocDate:      data["docdate"].(time.Time),
		UserAdd:      data["useradd"].(string),
		UserUpdate:   data["userupdate"].(string),
		Status:       data["status"].(bool),
		Type:         data["type"].(string),
		Barcode:      data["barcode"].(string),
		Image:        data["image"].(string),
		WhsCode:      data["whscode"].(string),
		CardCode:     data["cardcode"].(string),
		Bonus:        data["bonus"].(float64),
		IdentifiCode: data["identificode"].(string),
	}

	fmt.Println(item)
	result := db.Create(&item)

	if err := result.Error; err != nil {
		//fmt.Println(result)
		return err
	}
	return nil
}

// Edit Member modify a single Member
func UpdateItem_Model(id string, data map[string]interface{}) error {

	item := Items{
		// ItemCode:   data["ItemCode"].(string),
		ItemName:     data["ItemName"].(string),
		Quantity:     data["Quantity"].(int),
		Group:        data["Group"].(string),
		InStock:      data["InStock"].(float64),
		Price:        data["Price"].(float64),
		UserUpdate:   data["UserUpdate"].(string),
		Status:       data["Status"].(bool),
		Type:         data["Type"].(string),
		Barcode:      data["Barcode"].(string),
		Image:        data["Image"].(string),
		WhsCode:      data["WhsCode"].(string),
		CardCode:     data["CardCode"].(string),
		Bonus:        data["Bonus"].(float64),
		IdentifiCode: data["IdentifiCode"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		// "ItemCode": item.ItemCode,
		"ItemName":     item.ItemName,
		"Quantity":     item.Quantity,
		"Group":        item.Group,
		"InStock":      item.InStock,
		"Price":        item.Price,
		"Status":       item.Status,
		"Type":         item.Type,
		"Barcode":      item.Barcode,
		"Image":        item.Image,
		"WhsCode":      item.WhsCode,
		"IdentifiCode": item.IdentifiCode,
		"Bonus":        item.Bonus,
		"CardCode":     item.CardCode},
	).Error; err != nil {
		return err
	}
	return nil
}

func GetItems_Model() (*[]Items, error) {

	item := []Items{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Order("created_at desc").Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func GetById_Model(id string) (*Items, error) {
	var item Items

	err := db.Where("id = ?", id).First(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &item, nil
}

func SearchItems_Model(itemCode string, itemName string) (*[]Items, error) {
	item := []Items{}

	err := db.Raw(
		"select a.ID,a.ItemCode,a.ItemName,a.[Group] ,a.Quantity, a.InStock, a.Price, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type, a.Barcode, a.Image, a.WhsCode, a.CardCode from items a " +
			"where (charindex(N'" + itemCode + "', a.ItemCode) > 0) " +
			"union " +
			"select a.ID,a.ItemCode,a.ItemName,a.[Group], a.Quantity , a.InStock, a.Price, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type, a.Barcode, a.Image, a.WhsCode,a.CardCode from items a " +
			"where (charindex(N'" + itemName + "', a.ItemName) > 0) ",
	).Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func GetItemsIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast(('I' + '' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),created_at,12),4) ORDER BY ID)) countID from items where LEN(created_at) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
