package models

import (
	"gorm.io/gorm"
)

type ConfigSettingApp struct {
	gorm.Model
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

func AddSettings(data map[string]interface{}) error {
	item := ConfigSettingApp{
		Company:  data["company"].(string),
		Address:  data["address"].(string),
		Phone:    data["phone"].(string),
		MST:      data["mst"].(string),
		IP:       data["ip"].(string),
		Database: data["database"].(string),
		UserName: data["username"].(string),
		PassWord: data["password"].(string),
		Image:    data["image"].(string),
	}

	result := db.Create(&item)

	if err := result.Error; err != nil {
		//fmt.Println(result)
		return err
	}
	return nil
}

func UpdateSettings_Model(id string, data map[string]interface{}) error {

	item := ConfigSettingApp{
		Company:  data["Company"].(string),
		Address:  data["Address"].(string),
		Phone:    data["Phone"].(string),
		MST:      data["MST"].(string),
		IP:       data["IP"].(string),
		Database: data["Database"].(string),
		UserName: data["UserName"].(string),
		PassWord: data["PassWord"].(string),
		Image:    data["Image"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"Company":  item.Company,
		"Address":  item.Address,
		"Phone":    item.Phone,
		"MST":      item.MST,
		"IP":       item.IP,
		"Database": item.Database,
		"UserName": item.UserName,
		"PassWord": item.PassWord,
		"Image":    item.Image}).Error; err != nil {
		return err
	}
	return nil
}

func GetSettings_Model() (*[]ConfigSettingApp, error) {

	item := []ConfigSettingApp{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Order("created_at desc").Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func GetByIdSettings_Model(id string) (*ConfigSettingApp, error) {
	var item ConfigSettingApp

	err := db.Where("id = ?", id).First(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &item, nil
}
