package models

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	WhsCode    string    `gorm:"column:WhsCode" json:"whscode"`
	WhsName    string    `gorm:"column:WhsName" json:"whsname"`
	Address    string    `gorm:"column:Address" json:"address"`
	Remarks    string    `gorm:"column:Remarks" json:"remarks"`
	DocDate    time.Time `gorm:"column:DocDate" json:"docdate"`
	UserAdd    string    `gorm:"column:UserAdd" json:"useradd"`
	UserUpdate string    `gorm:"column:UserUpdate" json:"userupdate"`
	Status     bool      `gorm:"column:Status" json:"status"`
	Type       string    `gorm:"column:Type" json:"type"`
}

func AddWhs(data map[string]interface{}) error {
	whs := Warehouse{
		WhsCode: data["whscode"].(string),
		WhsName: data["whsname"].(string),
		// Group:      data["group"].(string),
		Address: data["address"].(string),
		// Phone:      data["phone"].(string),
		Remarks:    data["remarks"].(string),
		DocDate:    data["docdate"].(time.Time),
		UserAdd:    data["useradd"].(string),
		UserUpdate: data["userupdate"].(string),
		Status:     data["status"].(bool),
		Type:       data["type"].(string),
	}

	result := db.Create(&whs)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateWhs_Model(id string, data map[string]interface{}) error {
	whs := Warehouse{
		WhsName: data["WhsName"].(string),
		// Group:      data["Group"].(string),
		Address: data["Address"].(string),
		// Phone:      data["Phone"].(string),
		Remarks:    data["Remarks"].(string),
		UserAdd:    data["UserAdd"].(string),
		UserUpdate: data["UserUpdate"].(string),
		Status:     data["Status"].(bool),
		Type:       data["Type"].(string),
	}

	if err := db.Model(&whs).Where("id = ?", id).Updates(map[string]interface{}{
		"WhsName": whs.WhsName,
		// "Group":    bps.Group,
		"Address": whs.Address,
		// "Phone":    bps.Phone,
		"Remarks": whs.Remarks,
		"Status":  whs.Status,
		"Type":    whs.Type}).Error; err != nil {
		return err
	}
	return nil
}

func GetWhs_Model() (*[]Warehouse, error) {

	item := []Warehouse{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func GetByIdWhs_Model(id string) (*Warehouse, error) {
	var item Warehouse

	err := db.Where("id = ?", id).First(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &item, nil
}

func SearchWhs_Model(whsCode string, whsName string, types string) (*[]Warehouse, error) {
	whs := []Warehouse{}

	err := db.Raw(
		"select a.ID,a.CardCode,a.CardName,, a.Address,,a.Remarks, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type from warehouse a " +
			"where (charindex(N'" + whsCode + "', a.CardCode) > 0) and a.Type = '" + types + "' " +
			"union " +
			"select a.ID,a.CardCode,a.CardName,, a.Address, ,a.Remarks, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type from warehouse a " +
			"where (charindex(N'" + whsName + "', a.CardName) > 0) and a.Type = '" + types + "' ",
	).Find(&whs).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &whs, nil
}

func GetWhsCodeIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast((RIGHT(REPLICATE('0',1) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),created_at,12),4) ORDER BY ID)) countID from warehouses where LEN(created_at) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
