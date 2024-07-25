package warehouse_service

import (
	"sap-crm/models"
	"time"
)

type Whs struct {
	WhsCode string `gorm:"column:WhsCode" json:"whscode"`
	WhsName string `gorm:"column:WhsName" json:"whsname"`
	// Group      string    `gorm:"column:Group" json:"group"`
	// Resources  string    `gorm:"column:Resources" json:"resources"`
	// Tax        string    `gorm:"column:Tax" json:"tax"`
	Address string `gorm:"column:Address" json:"address"`
	// Phone      string    `gorm:"column:Phone" json:"phone"`
	Remarks    string    `gorm:"column:Remarks" json:"remarks"`
	DocDate    time.Time `gorm:"column:DocDate" json:"docdate"`
	UserAdd    string    `gorm:"column:UserAdd" json:"useradd"`
	UserUpdate string    `gorm:"column:UserUpdate" json:"userupdate"`
	Status     bool      `gorm:"column:Status" json:"status"`
	Type       string    `gorm:"column:Type" json:"type"`
}

func (a *Whs) AddWhs() error {
	currentime := time.Now()
	// theTime := currentime.Format("2006-1-2 15:4:5")
	_autoCode, _ := models.GetWhsCodeIDAuto_Model()
	item := map[string]interface{}{
		"whscode": _autoCode,
		"whsname": a.WhsName,
		// "group":      a.Group,
		// "resources":  a.Resources,
		// "tax":        a.Tax,
		"address": a.Address,
		// "phone":      a.Phone,
		"remarks":    a.Remarks,
		"docdate":    currentime,
		"useradd":    a.UserAdd,
		"userupdate": a.UserUpdate,
		"status":     a.Status,
		"type":       a.Type,
	}
	if err := models.AddWhs(item); err != nil {
		return err
	}

	return nil
}

func (a *Whs) UpdatWhs(id string) error {

	whs := map[string]interface{}{
		"WhsName": a.WhsName,
		// "Group":      a.Group,
		// "Resources":  a.Resources,
		// "Tax":        a.Tax,
		"Address": a.Address,
		// "Phone":      a.Phone,
		"Remarks":    a.Remarks,
		"UserAdd":    a.UserAdd,
		"UserUpdate": a.UserUpdate,
		"Status":     a.Status,
		"Type":       a.Type,
	}

	if err := models.UpdateWhs_Model(id, whs); err != nil {
		return err
	}

	return nil
}

func (a *Whs) GetWhs_Service() (*[]models.Warehouse, error) {
	item, err := models.GetWhs_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (a *Whs) GetByIdWhs_Service(id string) (*models.Warehouse, error) {
	item, err := models.GetByIdWhs_Model(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (a *Whs) SearchWhs_Service(whsCode string, whsName string, types string) (*[]models.Warehouse, error) {
	item, err := models.SearchWhs_Model(whsCode, whsName, types)
	if err != nil {
		return nil, err
	}
	return item, nil
}
