package table_service

import "sap-crm/models"

type Table struct {
	TableNumber string `gorm:"column:TableNumber" json:"tablenumber"`
	Remark      string `gorm:"column:Remark" json:"remark"`
	Status      string `gorm:"column:Status" json:"status"`
}

type FormSearch struct {
	TableNumber string `form:"TableNumber" valid:"required"`
	Remark      string `form:"Remark" valid:"required"`
	Status      string `form:"Status" valid:"required"`
}

func (a *Table) AddTable() error {
	// currentime := time.Now()
	// theTime := currentime.Format("2006-1-2 15:4:5")

	_autoCode, _ := models.GetTableIDAuto_Model()
	table := map[string]interface{}{
		"tablenumber": _autoCode,
		"remark":      a.Remark,
		"status":      a.Status,
	}
	if err := models.AddTable(table); err != nil {
		return err
	}

	return nil
}

func (a *Table) UpdateTable(id string) error {

	table := map[string]interface{}{
		//"TableNumber": a.TableNumber,
		"Remark": a.Remark,
		"Status": a.Status,
		// "Quantity":   a.CatagoryName,

	}

	if err := models.UpdateTable_Model(id, table); err != nil {
		return err
	}

	return nil
}

func (a *Table) GetTable_Service() (*[]models.Table, error) {
	table, err := models.GetTable_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return table, nil
}

func (a *Table) GetByIdTable_Service(id string) (*models.Table, error) {
	table, err := models.GetByIdTable_Model(id)
	if err != nil {
		return nil, err
	}

	return table, nil
}

func (a *Table) SearchTable_Service(TableNumber string) (*[]models.Table, error) {
	table, err := models.SearchTable_Model(TableNumber)
	if err != nil {
		return nil, err
	}
	return table, nil
}
