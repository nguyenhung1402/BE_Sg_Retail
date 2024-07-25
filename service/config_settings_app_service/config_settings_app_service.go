package config_settings_app_service

import (
	"sap-crm/models"
)

type ConfigSettingApp struct {
	Company  string `gorm:"column:Company" json:"company"`
	Address  string `gorm:"column:Address" json:"address"`
	Phone    string `gorm:"column:Phone" json:"phone"`
	MST      string `gorm:"column:MST" json:"mst"`
	IP       string `gorm:"column:IP" json:"ip"`
	Database string `gorm:"column:Database" json:"database"`
	UserName string `gorm:"column:UserName" json:"username"`
	PassWord string `gorm:"column:PassWord" json:"password"`
	Image    string `gorm:"column:Image" json:"image"`
}

func (a *ConfigSettingApp) AddSettings_service() error {
	item := map[string]interface{}{
		"company":  a.Company,
		"address":  a.Address,
		"phone":    a.Phone,
		"mst":      a.MST,
		"ip":       a.IP,
		"database": a.Database,
		"username": a.UserName,
		"password": a.PassWord,
		"image":    a.Image,
	}
	if err := models.AddSettings(item); err != nil {
		return err
	}

	return nil
}

func (a *ConfigSettingApp) UpdateUpdateSettings_service(id string) error {

	item := map[string]interface{}{
		"Company":  a.Company,
		"Address":  a.Address,
		"Phone":    a.Phone,
		"MST":      a.MST,
		"IP":       a.IP,
		"Database": a.Database,
		"UserName": a.UserName,
		"PassWord": a.PassWord,
		"Image":    a.Image,
	}

	if err := models.UpdateSettings_Model(id, item); err != nil {
		return err
	}

	return nil
}

func (a *ConfigSettingApp) GetSettings_Service() (*[]models.ConfigSettingApp, error) {
	item, err := models.GetSettings_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *ConfigSettingApp) GetByIdSettings_Service(id string) (*models.ConfigSettingApp, error) {
	item, err := models.GetByIdSettings_Model(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}
