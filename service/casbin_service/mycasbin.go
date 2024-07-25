package casbin_service

import "sap-crm/models"

type MyCasbin struct {
	Ptype    string `json:"ptype"`
	RoleName string `json:"roleName"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

// Add permissions
func (c *MyCasbin) AddCasbin() (bool, error) {
	item := map[string]interface{}{
		"ptype":    c.Ptype,
		"roleName": c.RoleName,
		"path":     c.Path,
		"method":   c.Method,
	}
	check, err := models.AddCasbin(item)
	if err != nil {
		return check, err
	}
	return check, nil
}
