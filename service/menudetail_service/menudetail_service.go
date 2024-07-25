package menudetail_service

import (
	"sap-crm/models"
)

type ChildMenu struct {
	IDMenu   int
	Title    string
	Icon     string
	Position string
	Rule     string
	Url      string
}

func (a *ChildMenu) GetByIdChilMenu_Service(id string) (*models.ChildMenu, error) {
	items, err := models.GetByIdChilMenu_Model(id)
	if err != nil {
		return nil, err
	}

	return items, nil
}
