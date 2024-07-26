package models

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	StaffCode  string    `gorm:"column:StaffCode" json:"staffcode"`
	StaffName  string    `gorm:"column:StaffName" json:"staffname"`
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

func AddStaff(data map[string]interface{}) error {
	staff := Staff{
		StaffCode:  data["staffcode"].(string),
		StaffName:  data["staffname"].(string),
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

	result := db.Create(&staff)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Edit BP modify a single Member
func UpdateStaff_Model(id string, data map[string]interface{}) error {
	staff := Staff{
		// CardCode:   data["CardCode"].(string),
		StaffName:  data["StaffName"].(string),
		Group:      data["Group"].(string),
		Address:    data["Address"].(string),
		Phone:      data["Phone"].(string),
		Remarks:    data["Remarks"].(string),
		UserAdd:    data["UserAdd"].(string),
		UserUpdate: data["UserUpdate"].(string),
		Status:     data["Status"].(bool),
		Type:       data["Type"].(string),
	}

	if err := db.Model(&staff).Where("id = ?", id).Updates(map[string]interface{}{
		// "CardCode": bps.CardCode,
		"StaffName": staff.StaffName,
		"Group":     staff.Group,
		"Address":   staff.Address,
		"Phone":     staff.Phone,
		"Remarks":   staff.Remarks,
		"Status":    staff.Status,
		"Type":      staff.Type}).Error; err != nil {
		return err
	}
	return nil
}

func GetStaff_Model() (*[]Staff, error) {

	item := []Staff{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func GetByIdStaff_Model(id string) (*Staff, error) {
	var item Staff

	err := db.Where("id = ?", id).First(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &item, nil
}

func SearchStaff_Model(staffCode string, staffName string, types string) (*[]Staff, error) {
	bp := []Staff{}

	err := db.Raw(
		"select a.ID,a.StaffCode,a.StaffName,a.[Group] , a.Address, a.Phone,a.Remarks, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type from staffs a " +
			"where (charindex(N'" + staffName + "', a.StaffCode) > 0) and a.Type = '" + types + "' " +
			"union " +
			"select a.ID,a.StaffCode,a.StaffName,a.[Group] , a.Address, a.Phone,a.Remarks, a.DocDate, a.UserAdd,a.UserUpdate, a.Status, a.Type from staffs a " +
			"where (charindex(N'" + staffName + "', a.StaffName) > 0) and a.Type = '" + types + "' ",
	).Find(&bp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &bp, nil
}
func GetStaffIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast(('ST' + '' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),created_at,12),4) ORDER BY ID)) countID from staffs where LEN(created_at) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
