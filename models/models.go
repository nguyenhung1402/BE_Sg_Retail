package models

import (
	"fmt"
	"log"
	"sap-crm/pkg/setting"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

// gorm.Model definition
type BaseModel struct {
	Id        uint           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deletedAt"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	// dns := fmt.Sprintf("%s://%s:%s@%s?database=RETAIL_V0", setting.DatabaseSetting.Type, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host)
	//dns := fmt.Sprintf("%s://%s:%s@%s?database=RETAIL_DEMO", setting.DatabaseSetting.Type, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host)
	dns := fmt.Sprintf("%s://%s:%s@%s?database=RETAIL_DEMO_V0", setting.DatabaseSetting.Type, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host)

	db, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{
		PrepareStmt: true,
	})

	db.AutoMigrate(&Member{}, &MemberCredentials{}, &Items{}, &BPs{}, &ParentMenu{}, &ChildMenu{}, &Order{}, &OrderDetails{}, &Warehouse{}, &POS{}, &POSDetails{}, &Receipt{}, &ReceiptDetails{}, &Issue{}, &IssueDetails{}, &Category{}, &Table{}, &ConfigSettingApp{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

}

// CloseDB closes database connection (unnecessary)
