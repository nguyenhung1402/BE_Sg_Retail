package viewnxt_service

import (
	"sap-crm/models"
)

type ViewNXT struct {
	WhsCode  string
	ItemCode string
	Quantity string
}

func (a *ViewNXT) GetAllNXT_Service() (*[]models.ViewNXT, error) {
	viewNXT, err := models.GetAllNXT_Model()
	if err != nil {
		return nil, err
	}
	return viewNXT, nil
}

func (a *ViewNXT) GetWhsCodeNXT_Service(whscode string) (*[]models.ViewNXT, error) {
	viewNXT, err := models.GetWhsCodeNXT_Model(whscode)
	if err != nil {
		return nil, err
	}
	return viewNXT, nil
}

func (a *ViewNXT) GetWhsItemCodeNXT_Service(whscode string, itemcode string) (*models.ViewNXT, error) {
	viewNXT, err := models.GetWhsItemCodeNXT_Model(whscode, itemcode)
	if err != nil {
		return nil, err
	}
	return viewNXT, nil
}
