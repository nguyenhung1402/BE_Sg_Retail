package category_service

import "sap-crm/models"

type Category struct {
	CategoryCode string `gorm:"column:CategoryCode" json:"categorycode"`
	CategoryName string `gorm:"column:CategoryName" json:"categoryname"`
	Status       string `gorm:"column:Status" json:"status"`
}

type FormSearch struct {
	CategoryCode string `form:"CategoryCode" valid:"required"`
	CategoryName string `form:"CategoryName" valid:"required"`
	Status       string `form:"Status" valid:"required"`
}

func (a *Category) AddCategory() error {
	// currentime := time.Now()
	// theTime := currentime.Format("2006-1-2 15:4:5")
	category := map[string]interface{}{
		"categorycode": a.CategoryCode,
		"categoryname": a.CategoryName,
		"status":       a.Status,
	}
	if err := models.AddCategory(category); err != nil {
		return err
	}

	return nil
}

func (a *Category) UpdateCategory(id string) error {

	category := map[string]interface{}{
		"CategoryName": a.CategoryName,
		"Status":       a.Status,
		// "Quantity":   a.CatagoryName,

	}

	if err := models.UpdateCategory_Model(id, category); err != nil {
		return err
	}

	return nil
}

func (a *Category) GetCategory_Service() (*[]models.Category, error) {
	category, err := models.GetCategory_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return category, nil
}

func (a *Category) GetByIdCategory_Service(id string) (*models.Category, error) {
	category, err := models.GetByIdCategory_Model(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (a *Category) SearchCategory_Service(CategoryCode string, CategoryName string) (*[]models.Category, error) {
	category, err := models.SearchCategory_Model(CategoryCode, CategoryName)
	if err != nil {
		return nil, err
	}
	return category, nil
}
