package staff_service

import (
	"sap-crm/models"
	"time"
)

type Staff struct {
	StaffCode  string    `gorm:"column:StaffCode" json:"staffcode"`
	StaffName  string    `gorm:"column:StaffName" json:"staffname"`
	Group      string    `gorm:"column:Group" json:"group"`
	Resources  string    `gorm:"column:Resources" json:"resources"`
	Tax        string    `gorm:"column:Tax" json:"tax"`
	Address    string    `gorm:"column:Address" json:"address"`
	Phone      string    `gorm:"column:Phone" json:"phone"`
	Remarks    string    `gorm:"column:Remarks" json:"remarks"`
	DocDate    time.Time `gorm:"column:DocDate" json:"docdate"`
	UserAdd    string    `gorm:"column:UserAdd" json:"useradd"`
	UserUpdate string    `gorm:"column:UserUpdate" json:"userupdate"`
	Status     bool      `gorm:"column:Status" json:"status"`
	Type       string    `gorm:"column:Type" json:"type"`
}

func (a *Staff) AddStaff() error {
	currentime := time.Now()
	_autoCode, _ := models.GetStaffIDAuto_Model()
	// theTime := currentime.Format("2006-1-2 15:4:5")
	item := map[string]interface{}{
		"staffcode":  _autoCode,
		"staffname":  a.StaffName,
		"group":      a.Group,
		"resources":  a.Resources,
		"tax":        a.Tax,
		"address":    a.Address,
		"phone":      a.Phone,
		"remarks":    a.Remarks,
		"docdate":    currentime,
		"useradd":    a.UserAdd,
		"userupdate": a.UserUpdate,
		"status":     a.Status,
		"type":       a.Type,
	}
	if err := models.AddStaff(item); err != nil {
		return err
	}

	return nil
}

func (a *Staff) UpdateStaff(id string) error {

	staff := map[string]interface{}{
		// "CardCode":   a.CardCode,
		"StaffName":  a.StaffName,
		"Group":      a.Group,
		"Resources":  a.Resources,
		"Tax":        a.Tax,
		"Address":    a.Address,
		"Phone":      a.Phone,
		"Remarks":    a.Remarks,
		"UserAdd":    a.UserAdd,
		"UserUpdate": a.UserUpdate,
		"Status":     a.Status,
		"Type":       a.Type,
	}

	if err := models.UpdateStaff_Model(id, staff); err != nil {
		return err
	}

	return nil
}

func (a *Staff) GetStaff_Service() (*[]models.Staff, error) {
	item, err := models.GetStaff_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (a *Staff) GetByIdStaff_Service(id string) (*models.Staff, error) {
	item, err := models.GetByIdStaff_Model(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (a *Staff) SearchStaff_Service(staffCode string, staffName string, types string) (*[]models.Staff, error) {
	item, err := models.SearchStaff_Model(staffCode, staffCode, types)
	if err != nil {
		return nil, err
	}
	return item, nil
}
