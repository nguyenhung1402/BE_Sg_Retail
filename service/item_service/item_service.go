package item_service

import (
	"sap-crm/models"
	"time"
)

type Items struct {
	ItemCode   string    `gorm:"column:ItemCode" json:"itemcode"`
	ItemName   string    `gorm:"column:ItemName" json:"itemname"`
	Quantity   int       `gorm:"column:Quantity" json:"quantity"`
	Group      string    `gorm:"column:Group" json:"group"`
	InStock    float64   `gorm:"column:InStock" json:"instock"`
	Price      float64   `gorm:"column:Price" json:"price"`
	DocDate    time.Time `gorm:"column:DocDate" json:"docdate"`
	UserAdd    string    `gorm:"column:UserAdd" json:"useradd"`
	UserUpdate string    `gorm:"column:UserUpdate" json:"userupdate"`
	Status     bool      `gorm:"column:Status" json:"status"`
	Type       string    `gorm:"column:Type" json:"type"`
	Barcode    string    `gorm:"column:Barcode" json:"barcode"`
	Image      string    `gorm:"column:Image" json:"image"`
	WhsCode    string    `gorm:"column:WhsCode" json:"whscode"`
	CardCode   string    `gorm:"column:CardCode" json:"cardcode"`
}

type FormSearch struct {
	ItemCode string `form:"ItemCode" valid:"required"`
	ItemName string `form:"ItemName" valid:"required"`
}

func (a *Items) AddItems() error {
	currentime := time.Now()
	// theTime := currentime.Format("2006-1-2 15:4:5")
	_autoCode, _ := models.GetItemsIDAuto_Model()
	item := map[string]interface{}{
		"itemcode":   _autoCode,
		"itemname":   a.ItemName,
		"quantity":   a.Quantity,
		"group":      a.Group,
		"instock":    a.InStock,
		"price":      a.Price,
		"docdate":    currentime,
		"useradd":    a.UserAdd,
		"userupdate": a.UserUpdate,
		"status":     a.Status,
		"type":       a.Type,
		"barcode":    a.Barcode,
		"image":      a.Image,
		"whscode":    a.WhsCode,
		"cardcode":   a.CardCode,
	}
	if err := models.AddItems(item); err != nil {
		return err
	}

	return nil
}

func (a *Items) UpdateItems(id string) error {

	item := map[string]interface{}{
		// "ItemCode":   a.ItemCode,
		"ItemName":   a.ItemName,
		"Quantity":   a.Quantity,
		"Group":      a.Group,
		"InStock":    a.InStock,
		"Price":      a.Price,
		"UserUpdate": a.UserUpdate,
		"Status":     a.Status,
		"Type":       a.Type,
		"Barcode":    a.Barcode,
		"Image":      a.Image,
		"WhsCode":    a.WhsCode,
		"CardCode":   a.CardCode,
	}

	if err := models.UpdateItem_Model(id, item); err != nil {
		return err
	}

	return nil
}

func (a *Items) GetItems_Service() (*[]models.Items, error) {
	item, err := models.GetItems_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *Items) GetById_Service(id string) (*models.Items, error) {
	item, err := models.GetById_Model(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (a *Items) SearchItems_Service(itemCode string, itemname string) (*[]models.Items, error) {
	item, err := models.SearchItems_Model(itemCode, itemname)
	if err != nil {
		return nil, err
	}
	return item, nil
}
