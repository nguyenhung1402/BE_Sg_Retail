package bp_service

import (
	"sap-crm/models"
	"time"
)

type BPs struct {
	CardCode   string    `gorm:"column:CardCode" json:"cardcode"`
	CardName   string    `gorm:"column:CardName" json:"cardname"`
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

func (a *BPs) AddBPs() error {
	currentime := time.Now()
	_autoCode, _ := models.GetBPSIDAuto_Model()
	// theTime := currentime.Format("2006-1-2 15:4:5")
	item := map[string]interface{}{
		"cardcode":   _autoCode,
		"cardname":   a.CardName,
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
	if err := models.AddBPs(item); err != nil {
		return err
	}

	return nil
}

func (a *BPs) UpdatBPs(id string) error {

	bps := map[string]interface{}{
		// "CardCode":   a.CardCode,
		"CardName":   a.CardName,
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

	if err := models.UpdateBPs_Model(id, bps); err != nil {
		return err
	}

	return nil
}

func (a *BPs) GetBPs_Service() (*[]models.BPs, error) {
	item, err := models.GetBPs_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (a *BPs) GetByIdBPs_Service(id string) (*models.BPs, error) {
	item, err := models.GetByIdBPs_Model(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (a *BPs) SearchBPs_Service(cardCode string, cardName string, types string) (*[]models.BPs, error) {
	item, err := models.SearchBPs_Model(cardCode, cardName, types)
	if err != nil {
		return nil, err
	}
	return item, nil
}
