package member_service

import (
	"sap-crm/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	Username          string
	Birthday          time.Time
	Email             string
	Image             string
	Firstname         string
	Lastname          string
	Status            string
	Phone             string
	City              string
	Address           string
	Address2          string
	Role              string
	IP                string
	NameDatabase      string
	ResetKey          string
	ResetCount        int
	ResetTimestamp    string
	ResetKeyTimestamp string
	Password          string
}

func (m *Member) RegisterPublicUser(password string) error {
	salt := 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)

	if err != nil {
		return err
	}

	member := map[string]interface{}{
		"username":          m.Username,
		"birthday":          time.Now(),
		"email":             m.Email,
		"image":             m.Image,
		"firstName":         m.Firstname,
		"lastName":          m.Firstname,
		"status":            "1",
		"phone":             m.Phone,
		"city":              m.City,
		"address":           m.Address,
		"address2":          "",
		"role":              "user",
		"ip":                m.IP,
		"namedatabase":      m.NameDatabase,
		"resetKey":          "",
		"resetKeyTimestamp": "",
		"resetTimestamp":    "",
		"resetCount":        0,
		"password":          string(hashedPassword),
	}

	if err := models.AddMember(member); err != nil {
		return err
	}

	return nil
}

func (m *Member) GetMemberUsername() (*models.Member, error) {
	member, err := models.GetMemberUserName(m.Username)

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (m *Member) GetMemberEmail() (*models.Member, error) {
	member, err := models.GetMemberEmail(m.Email)

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (m *Member) GetmemberLogin() (*models.Member, error) {
	member, err := models.GetmemberLogin(m.Username)

	if err != nil {
		return nil, err
	}

	return member, nil
}

// Get All User
func (a *Member) GetUser_Service() (*[]models.Member, error) {
	item, err := models.GetAllUser_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Get ID User
func (a *Member) GetByIdUser_Service(id string) (*models.Member, error) {
	item, err := models.GetByIdUser_Model(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// Search Username User
func (a *Member) SearchMember_Service(username string) (*[]models.Member, error) {
	item, err := models.SearchMember_Model(username)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Update User
func (a *Member) UpdatMembers(id string) error {

	member := map[string]interface{}{
		"Username":     a.Username,
		"Birthday":     a.Birthday,
		"Email":        a.Email,
		"Image":        a.Image,
		"Firstname":    a.Firstname,
		"Lastname":     a.Lastname,
		"Status":       a.Status,
		"Phone":        a.Phone,
		"City":         a.City,
		"Address":      a.Address,
		"Address2":     a.Address2,
		"Role":         a.Role,
		"IP":           a.IP,
		"NameDatabase": a.NameDatabase,
	}

	if err := models.UpdateMembers_Model(id, member); err != nil {
		return err
	}

	return nil
}
