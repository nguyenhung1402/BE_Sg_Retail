package models

import (
	"gorm.io/gorm"
)

type ParentMenu struct {
	gorm.Model
	Title     string      `gorm:"column:Title" json:"title"`
	Icon      string      `gorm:"column:Icon" json:"icon"`
	Position  string      `gorm:"column:Position" json:"position"`
	Rule      string      `gorm:"column:Rule" json:"rule"`
	ChildMenu []ChildMenu `gorm:"foreignKey:IDMenu;" json:"childmenu"`
}

func AddParentMenu(data map[string]interface{}, dataDetail []map[string]interface{}) error {
	menuDetails := make([]ChildMenu, len(dataDetail))
	for i, v := range dataDetail {
		item := ChildMenu{
			Title:    v["Title"].(string),
			Icon:     v["Icon"].(string),
			Position: v["Position"].(string),
			Rule:     v["Rule"].(string),
			Url:      v["Url"].(string),
		}
		menuDetails[i] = item
	}
	menu := ParentMenu{
		Title:     data["title"].(string),
		Icon:      data["icon"].(string),
		Position:  data["position"].(string),
		Rule:      data["rule"].(string),
		ChildMenu: menuDetails,
	}
	result := db.Create(&menu)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// get all menu
func GetMenu_Model(id string) (*[]ParentMenu, error) {

	item := []ParentMenu{}
	values := parseIDToValues(id)
	// err := db.Debug().Preload("ChildMenu").Find(&item).Error

	err := db.Debug().Where("[Rule] LIKE ?", "%"+values+"%").Preload("ChildMenu").Find(&item).Error
	// err := db.Preload(clause.Associations).Where("[Rule] LIKE ?", "%"+values+"%").Preload("ParentMenu." + clause.Associations).Find(&item).Error

	//err := db.Table("parent_menus").Preload("childmenu").Joins("INNER JOIN child_menus a ON parent_menus.id = a.IDMenu").Where("parent_menus.[Rule] LIKE " + "'" + " %" + values + "%" + "'" + " OR a.[Rule] LIKE " + "'" + "%" + values + "%" + "'" + " ").Find(&item).Error

	// err := db.Raw(
	// 	"select t0.Title AS 'title', t0.Icon as 'icon', t1.Title AS 'titleChild', t1.Icon as 'iconChild', t1.Url as 'url' from parent_menus t0 inner join child_menus t1 on t0.id = t1.IDMenu  " +
	// 		" where (t0.[Rule] like (case when(" + id + " = '1') then '%admin%' when(" + id + " = '2') then '%manager%' when(" + id + "='3') then '%sale%' when(" + id + "='4') then '%purchase%' when(" + id + "='5') then '%customer%' when(" + id + "='6') then '%user%' else '' end))  " +
	// 		" or (t1.[Rule] like (case when(" + id + " = '1') then '%admin%' when(" + id + " = '2') then '%manager%' when(" + id + "='3') then '%sale%' when(" + id + "='4') then '%purchase%' when(" + id + "='5') then '%customer%' when(" + id + "='6') then '%user%' else '' end))  ",
	// ).Find(&item).Error

	// err := db.Raw(
	// 	"select t0.Title AS 'title', t0.Icon as 'icon', t1.Title AS 'titleChild', t1.Icon as 'iconChild', t1.Url as 'url' from parent_menus t0 inner join child_menus t1 on t0.id = t1.IDMenu  " +
	// 	" where (t0." + '"' + "Rule" + '"' + " like (case when(" + id + " = '1') then '%admin%' when(" + id + " = '2') then '%manager%' when(" + id + "='3') then '%sale%' when(" + id + "='4') then '%purchase%' when(" + id + "='5') then '%customer%' when(" + id + "='6') then '%user%' else '' end))  " +
	// 	" or (t1." + '"' + "Rule" + '"' + " like (case when(" + id + " = '1') then '%admin%' when(" + id + " = '2') then '%manager%' when(" + id + "='3') then '%sale%' when(" + id + "='4') then '%purchase%' when(" + id + "='5') then '%customer%' when(" + id + "='6') then '%user%' else '' end))  ",
	// ).Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

func parseIDToValues(values string) string {
	if values == "1" {
		return "admin"
	} else if values == "2" {
		return "manager"
	} else if values == "3" {
		return "sale"
	} else if values == "4" {
		return "purchase"
	} else if values == "5" {
		return "customer"
	} else {
		return "user"
	}
}

// Get id meun
func GetByIdParentMenu_Model(id string) (*ParentMenu, error) {
	var parentMenu ParentMenu

	err := db.Where("id = ?", id).Preload("ChildMenu").First(&parentMenu).Error
	// err := db.Raw("select * from sos t0  " +
	// 	" inner join so_details t1 " +
	// 	" on t0.id = t1.IDSO " +
	// 	" where t0.id = " + id + " ",
	// ).Find(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &parentMenu, nil
}
