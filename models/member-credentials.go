package models

import (
	"gorm.io/gorm"
)

type MemberCredentials struct {
	BaseModel
	Password string `gorm:"column:password" json:"password"`
	MemberId uint   `gorm:"column:member_id" json:"memberId"`
}

func (u *MemberCredentials) FindCredentials(memberId string) (*MemberCredentials, error) {
	var memberCredentials MemberCredentials
	err := db.Where("member_id = ?", memberId).First(&memberCredentials).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &memberCredentials, err
}
