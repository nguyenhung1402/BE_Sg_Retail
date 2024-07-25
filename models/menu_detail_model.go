package models

import (
	"gorm.io/gorm"
)

type ChildMenu struct {
	BaseModel
	IDMenu   uint   `gorm:"column:IDMenu" json:"idmenu"`
	Title    string `gorm:"column:Title" json:"title"`
	Icon     string `gorm:"column:Icon" json:"icon"`
	Position string `gorm:"column:Position" json:"position"`
	Rule     string `gorm:"column:Rule" json:"rule"`
	Url      string `gorm:"column:Url" json:"url"`
}

func (u *ChildMenu) FindChilMenu(poid string) (*ChildMenu, error) {
	var items ChildMenu
	err := db.Where("IDPO = ?", poid).First(&items).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &items, err
}

// Get id po
func GetByIdChilMenu_Model(id string) (*ChildMenu, error) {
	var items ChildMenu

	err := db.Where("id = ?", id).First(&items).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &items, nil
}
