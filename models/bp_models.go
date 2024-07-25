package models

import (
	"time"

	"gorm.io/gorm"
)

type BPs struct {
	gorm.Model
	CardCode   string    `gorm:"column:CardCode" json:"cardcode"`
	CardName   string    `gorm:"column:CardName" json:"cardname"`
	Group      string    `gorm:"column:Group" json:"group"`
	Address    string    `gorm:"column:Address" json:"address"`
	Phone      string    `gorm:"column:Phone" json:"phone"`
	Remarks    string    `gorm:"column:Remarks" json:"remarks"`
	DocDate    time.Time `gorm:"column:DocDate" json:"docdate"`
	UserAdd    string    `gorm:"column:UserAdd" json:"useradd"`
	UserUpdate string    `gorm:"column:UserUpdate" json:"userupdate"`
	Status     bool      `gorm:"column:Status" json:"status"`
	Type       string    `gorm:"column:Type" json:"type"`
}

func AddBPs(data map[string]interface{}) error {
	bps := BPs{
		CardCode:   data["cardcode"].(string),
		CardName:   data["cardname"].(string),
		Group:      data["group"].(string),
		Address:    data["address"].(string),
		Phone:      data["phone"].(string),
		Remarks:    data["remarks"].(string),
		DocDate:    data["docdate"].(time.Time),
		UserAdd:    data["useradd"].(string),
		UserUpdate: data["userupdate"].(string),
		Status:     data["status"].(bool),
		Type:       data["type"].(string),
	}

	result := db.Create(&bps)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Edit BP modify a single Member
func UpdateBPs_Model(id string, data map[string]interface{}) error {
	bps := BPs{
		// CardCode:   data["CardCode"].(string),
		CardName:   data["CardName"].(string),
		Group:      data["Group"].(string),
		Address:    data["Address"].(string),
		Phone:      data["Phone"].(string),
		Remarks:    data["Remarks"].(string),
		UserAdd:    data["UserAdd"].(string),
		UserUpdate: data["UserUpdate"].(string),
		Status:     data["Status"].(bool),
		Type:       data["Type"].(string),
	}

	if err := db.Model(&bps).Where("id = ?", id).Updates(map[string]interface{}{
		// "CardCode": bps.CardCode,
		"CardName": bps.CardName,
		"Group":    bps.Group,
		"Address":  bps.Address,
		"Phone":    bps.Phone,
		"Remarks":  bps.Remarks,
		"Status":   bps.Status,
		"Type":     bps.Type}).Error; err != nil {
		return err
	}
	return nil
}

func GetBPs_Model() (*[]BPs, error) {

	item := []BPs{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func GetByIdBPs_Model(id string) (*BPs, error) {
	var item BPs

	err := db.Where("id = ?", id).First(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &item, nil
}

func SearchBPs_Model(cardCode string, cardName string, types string) (*[]BPs, error) {
	bp := []BPs{}

	err := db.Raw(
		"select a.ID,a.CardCode,a.CardName,a.[Group] , a.Address, a.Phone,a.Remarks, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type from b_ps a " +
			"where (charindex(N'" + cardCode + "', a.CardCode) > 0) and a.Type = '" + types + "' " +
			"union " +
			"select a.ID,a.CardCode,a.CardName,a.[Group] , a.Address, a.Phone,a.Remarks, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type from b_ps a " +
			"where (charindex(N'" + cardName + "', a.CardName) > 0) and a.Type = '" + types + "' ",
	).Find(&bp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &bp, nil
}
func GetBPSIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast(('BP' + '' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),created_at,12),4) ORDER BY ID)) countID from b_ps where LEN(created_at) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
