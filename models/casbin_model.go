package models

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

// Increase the column size to 512.
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

// Authority structure
type CasbinModel struct {
	BaseModel
	Ptype    string `json:"ptype"`
	RoleName string `json:"roleName"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

// Add permissions
func AddCasbin(data map[string]interface{}) (bool, error) {
	cm := CasbinModel{
		Ptype:    data["ptype"].(string),
		RoleName: data["roleName"].(string),
		Path:     data["path"].(string),
		Method:   data["method"].(string),
	}
	e := Casbin()
	return e.AddPolicy(cm.RoleName, cm.Path, cm.Method)

}

// Persist to the database
func Casbin() *casbin.Enforcer {
	// a, err := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	e, err := casbin.NewEnforcer("config/model.conf", "config/rbac_policy.csv")
	// Use pattern matching in RBAC
	e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
	if err != nil {
		log.Fatal(err)
	}

	err = e.LoadPolicy()

	if err != nil {
		log.Fatal(err)
	}
	// Save the policy back to DB.
	//e.SavePolicy()
	return e
}
