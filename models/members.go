package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	BaseModel
	Username          string            `gorm:"column:username" json:"username"`
	Birthday          time.Time         `gorm:"column:birthday" json:"birthday"`
	Email             string            `gorm:"column:email" json:"email"`
	Image             string            `gorm:"column:image" json:"image"`
	Firstname         string            `gorm:"column:first_name" json:"firstName"`
	Lastname          string            `gorm:"column:last_name" json:"lastName"`
	Status            string            `gorm:"column:status" json:"status"`
	Phone             string            `gorm:"column:phone" json:"phone"`
	City              string            `gorm:"column:city" json:"city"`
	Address           string            `gorm:"column:address" json:"address"`
	Address2          string            `gorm:"column:address2" json:"address2"`
	Role              string            `gorm:"column:role;" json:"role"`
	IP                string            `gorm:"column:ip;" json:"ip"`
	NameDatabase      string            `gorm:"column:namedatabase;" json:"namedatabase"`
	ResetKey          string            `gorm:"column:reset_key" json:"resetKey"`
	ResetCount        int               `gorm:"column:reset_count" json:"resetCount"`
	ResetTimestamp    string            `gorm:"column:reset_timestamp" json:"resetTimestamp"`
	ResetKeyTimestamp string            `gorm:"column:reset_key_timestamp" json:"resetKeyTimestamp"`
	MemberCredentials MemberCredentials `gorm:"foreignKey:member_id;" json:"memberCredentials"`
}

func (a *Member) GetMembers(maps interface{}) (*Member, error) {
	var member Member

	if maps != nil {
		err := db.Where("id = ?", a.Id).Where(maps).First(&member).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		err := db.Where("id = ?", a.Id).First(&member).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &member, nil
}

func GetMember(nickName string) (*Member, error) {
	var member Member

	err := db.Where("username = ?", nickName).First(&member).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &member, nil
}

// Edit Member modify a single Member
func EditMember(id int, data interface{}) error {
	if err := db.Model(&Member{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditMemberByCondition(condition interface{}, data interface{}) error {
	if err := db.Model(&Member{}).Where(condition).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddMember(data map[string]interface{}) error {
	member := Member{
		Username:          data["username"].(string),
		Birthday:          data["birthday"].(time.Time),
		Email:             data["email"].(string),
		Image:             data["image"].(string),
		Firstname:         data["firstName"].(string),
		Lastname:          data["lastName"].(string),
		Status:            data["status"].(string),
		Phone:             data["phone"].(string),
		City:              data["city"].(string),
		Address:           data["address"].(string),
		Address2:          data["address2"].(string),
		Role:              data["role"].(string),
		IP:                data["ip"].(string),
		NameDatabase:      data["namedatabase"].(string),
		ResetKey:          data["resetKey"].(string),
		ResetCount:        data["resetCount"].(int),
		ResetTimestamp:    data["resetTimestamp"].(string),
		ResetKeyTimestamp: data["resetKeyTimestamp"].(string),
		MemberCredentials: MemberCredentials{Password: data["password"].(string)},
	}
	if err := db.Create(&member).Error; err != nil {
		return err
	}

	return nil
}

func UpdateMembers_Model(id string, data map[string]interface{}) error {
	member := Member{
		Birthday:     data["birthday"].(time.Time),
		Email:        data["email"].(string),
		Image:        data["image"].(string),
		Firstname:    data["firstName"].(string),
		Lastname:     data["lastName"].(string),
		Status:       data["status"].(string),
		Phone:        data["phone"].(string),
		City:         data["city"].(string),
		Address:      data["address"].(string),
		Address2:     data["address2"].(string),
		Role:         data["role"].(string),
		IP:           data["ip"].(string),
		NameDatabase: data["namedatabase"].(string),
		// ResetKey:          data["resetKey"].(string),
		// ResetCount:        data["resetCount"].(int),
		// ResetTimestamp:    data["resetTimestamp"].(string),
		// ResetKeyTimestamp: data["resetKeyTimestamp"].(string),
		//MemberCredentials: MemberCredentials{Password: data["password"].(string)},
	}

	if err := db.Model(&member).Where("id = ?", id).Updates(map[string]interface{}{
		"Birthday":     member.Birthday,
		"Email":        member.Email,
		"Image":        member.Image,
		"Firstname":    member.Firstname,
		"Lastname":     member.Lastname,
		"Status":       member.Status,
		"Phone":        member.Phone,
		"City":         member.City,
		"Address":      member.Address,
		"Address2":     member.Address2,
		"Role":         member.Role,
		"IP":           member.IP,
		"NameDatabase": member.NameDatabase}).Error; err != nil {
		return err
	}
	return nil
}

func GetMemberUserName(username string) (*Member, error) {
	var member Member
	err := db.Model(&Member{}).Where("username = ?", username).First(&member).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &member, nil
}

func GetMemberEmail(email string) (*Member, error) {
	var member Member

	err := db.Model(&Member{}).Where("email = ?", email).First(&member).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &member, nil
}

func GetmemberLogin(username string) (*Member, error) {
	var member Member
	err := db.Model(&Member{}).Preload("MemberCredentials").Where("username = ?", username).First(&member).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &member, nil
}

// Get All user
func GetAllUser_Model() (*[]Member, error) {

	item := []Member{}
	// check loi database
	// err := db.Debug().Find(&item).Error
	err := db.Debug().Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

// Get ID User
func GetByIdUser_Model(id string) (*Member, error) {
	var item Member

	err := db.Where("id = ?", id).First(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &item, nil
}

// Search Username user
func SearchMember_Model(username string) (*[]Member, error) {
	item := []Member{}

	err := db.Where("username LIKE ?", "%"+username+"%").Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}
