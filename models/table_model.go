package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	TableNumber string `gorm:"column:TableNumber" json:"tablenumber"`
	Remark      string `gorm:"column:Remark" json:"remark"`
	Status      string `gorm:"column:Status" json:"status"`
}

func AddTable(data map[string]interface{}) error {
	table := Table{
		TableNumber: data["tablenumber"].(string),
		Remark:      data["remark"].(string),
		Status:      data["status"].(string),
	}

	fmt.Println(table)
	result := db.Create(&table)

	if err := result.Error; err != nil {
		//fmt.Println(result)
		return err
	}
	return nil
}

func UpdateTable_Model(id string, data map[string]interface{}) error {

	table := Table{
		// TableNumber: data["TableNumber"].(string),
		Remark: data["Remark"].(string),
		Status: data["Status"].(string),
	}

	if err := db.Model(&table).Where("id = ?", id).Updates(map[string]interface{}{
		// "TableNumber": table.TableNumber,
		"Remark": table.Remark,
		"Status": table.Status,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GetTable_Model() (*[]Table, error) {

	table := []Table{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Find(&table).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &table, nil
}

func GetByIdTable_Model(id string) (*Table, error) {
	var table Table

	err := db.Where("id = ?", id).First(&table).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &table, nil
}

func SearchTable_Model(TableNumber string) (*[]Table, error) {
	table := []Table{}

	err := db.Raw(
		"select a.ID,a.TableNumber from table a " +
			"where (charindex(N'" + TableNumber + "', a.TableNumber) > 0) ",
	).Find(&table).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &table, nil
}

func GetTableIDAuto_Model() (string, error) {
	var countID2 string
	rows, err := db.Raw("select cast(('TB' + '' +  RIGHT(REPLICATE('0',3) +  cast((COUNT(countID)+1) as nvarchar(10)),3)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),created_at,12),4) ORDER BY ID)) countID from tables where LEN(created_at) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}
