package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryCode string `gorm:"column:CategoryCode" json:"categorycode"`
	CategoryName string `gorm:"column:CategoryName" json:"categoryname"`
	Status       string `gorm:"column:Status" json:"status"`
}

func AddCategory(data map[string]interface{}) error {
	category := Category{
		CategoryCode: data["categorycode"].(string),
		CategoryName: data["categoryname"].(string),
		Status:       data["status"].(string),
	}

	fmt.Println(category)
	result := db.Create(&category)

	if err := result.Error; err != nil {
		//fmt.Println(result)
		return err
	}
	return nil
}

// Edit Member modify a single Member
func UpdateCategory_Model(id string, data map[string]interface{}) error {

	category := Category{
		CategoryName: data["CategoryName"].(string),
		Status:       data["status"].(string),
	}

	if err := db.Model(&category).Where("id = ?", id).Updates(map[string]interface{}{
		"CategoryName": category.CategoryName,
		"Status":       category.Status,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GetCategory_Model() (*[]Category, error) {

	category := []Category{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Find(&category).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &category, nil
}

func GetByIdCategory_Model(id string) (*Category, error) {
	var category Category

	err := db.Where("id = ?", id).First(&category).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &category, nil
}

func SearchCategory_Model(CategoryCode string, CategoryName string) (*[]Category, error) {
	category := []Category{}

	err := db.Raw(
		"select a.ID,a.CategoryCode,a.CategoryName,a.Status from category a " +
			"where (charindex(N'" + CategoryCode + "', a.CategoryCode) > 0) " +
			"union " +
			"select a.ID,a.CategoryCode,a.CategoryName,a.Status from category a " +
			"where (charindex(N'" + CategoryName + "', a.CategoryName) > 0) ",
	).Find(&category).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &category, nil
}
