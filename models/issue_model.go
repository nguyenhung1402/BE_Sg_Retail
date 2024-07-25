package models

import (
	// "time"

	"fmt"
	"time"

	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	WhsCode      string         `gorm:"column:WhsCode" json:"whscode"`
	CardCode     string         `gorm:"column:CardCode" json:"cardcode"`
	CardName     string         `gorm:"column:CardName" json:"cardname"`
	Address      string         `gorm:"column:Address" json:"address"`
	Phone        string         `gorm:"column:Phone" json:"phone"`
	Creator      string         `gorm:"column:Creator" json:"creator"`
	AutoCode     string         `gorm:"column:AutoCode" json:"autocode"`
	IssueDetails []IssueDetails `gorm:"foreignKey:IDIssue;" json:"issueDetails"`
}

func AddIssue(data map[string]interface{}, dataDetail []map[string]interface{}) error {
	issueDetails := make([]IssueDetails, len(dataDetail))
	for i, v := range dataDetail {
		issue := IssueDetails{
			ItemCode: v["ItemCode"].(string),
			ItemName: v["ItemName"].(string),
			Quantity: v["Quantity"].(float64),
		}
		issueDetails[i] = issue
	}
	issue := Issue{
		WhsCode:      data["whscode"].(string),
		CardCode:     data["cardcode"].(string),
		CardName:     data["cardname"].(string),
		Address:      data["address"].(string),
		Phone:        data["phone"].(string),
		Creator:      data["creator"].(string),
		AutoCode:     data["autocode"].(string),
		IssueDetails: issueDetails,
	}
	result := db.Create(&issue)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Edit Member modify a single Member
func UpdateIssue_Model(id string, data map[string]interface{}) error {
	issue := Issue{

		WhsCode:  data["whscode"].(string),
		CardCode: data["cardcode"].(string),
		CardName: data["cardname"].(string),
		Address:  data["address"].(string),
		Phone:    data["phone"].(string),
		Creator:  data["creator"].(string),
	}
	if err := db.Model(&issue).Where("id = ?", id).Updates(map[string]interface{}{
		"WhsCode":  issue.WhsCode,
		"CardCode": issue.CardCode,
		"CardName": issue.CardName,
		"Address":  issue.Address,
		"Phone":    issue.Phone,
		"Creator":  issue.Creator,
	}).Error; err != nil {
		return err
	}
	return nil
}

// get all po
func GetIssue_Model() (*[]Issue, error) {
	currentTime := time.Now()
	var abc = currentTime.Format("01-02-2006")
	fmt.Println(abc)
	issue := []Issue{}
	err := db.Debug().Where("cast(created_at as date) = ?", abc).Order("created_at desc").Find(&issue).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &issue, nil
}

// Get id po
func GetByIdIssue_Model(id string) (*Issue, error) {
	var issue Issue

	err := db.Where("id = ?", id).Preload("IssueDetails").First(&issue).Error
	// err := db.Raw("select * from sos t0  " +
	// 	" inner join so_details t1 " +
	// 	" on t0.id = t1.IDSO " +
	// 	" where t0.id = " + id + " ",
	// ).Find(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &issue, nil
}

func SearchIssue_Model(whscode string, cardcode string) (*[]Issue, error) {
	issue := []Issue{}

	err := db.Raw(
		"select * from issue a " +
			"where (charindex(N'" + whscode + "', a.TableName) > 0) " +
			"union " +
			"select * from issue a " +
			"where (charindex(N'" + cardcode + "', a.TableNumber) > 0) ",
	).Find(&issue).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &issue, nil
}

func GetIssueTableName(tablename string) (*Issue, error) {
	var issue Issue
	err := db.Model(&Issue{}).Where("username = ?", tablename).First(&issue).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &issue, nil
}

func GetIssueIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast(('XK'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),created_at,12),4) ORDER BY ID)) countID from issues where (MONTH(created_at) = MONTH(getdate()) and YEAR(created_at) = YEAR(GETDATE())) and LEN(created_at) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
