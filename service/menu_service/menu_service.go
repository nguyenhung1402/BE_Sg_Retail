package menu_service

import (
	"sap-crm/models"
)

type ParentMenu struct {
	Title     string
	Icon      string
	Position  string
	Rule      string
	ChildMenu []ChildMenu
}

type ChildMenu struct {
	IDMenu   int
	Title    string
	Icon     string
	Position string
	Rule     string
	Url      string
}

// Add PO
func (a *ParentMenu) AddParentMenu_Service() error {
	list := make([]map[string]interface{}, len(a.ChildMenu))
	for i, v := range a.ChildMenu {
		item := map[string]interface{}{
			"Title":    v.Title,
			"Icon":     v.Icon,
			"Position": v.Position,
			"Rule":     v.Rule,
			"Url":      v.Url,
		}
		list[i] = item
	}
	item := map[string]interface{}{
		"title":     a.Title,
		"icon":      a.Icon,
		"position":  a.Position,
		"rule":      a.Rule,
		"childmenu": a.ChildMenu,
	}
	if err := models.AddParentMenu(item, list); err != nil {
		return err
	}

	return nil
}

// Get all po
func (a *ParentMenu) GetMenu_Service(id string) (*[]models.ParentMenu, error) {
	menu, err := models.GetMenu_Model(id)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return menu, nil
}

// get id po

func (a *ParentMenu) GetByIdMenu_Service(id string) (*models.ParentMenu, error) {
	menu, err := models.GetByIdParentMenu_Model(id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}
