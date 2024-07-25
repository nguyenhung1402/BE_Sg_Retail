package v1

import (
	"sap-crm/pkg/app"
	"sap-crm/pkg/utils"
	"sap-crm/service/member_service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func SignIn(c *fiber.Ctx) error {
	form := &LoginForm{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err.Error(),
			"message": err.Error(),
		})
	}

	// Validate
	errors := app.ValidateStruct(*form)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"errors":  errors,
			"message": errors,
		})

	}

	// JWT
	check, role, ip, nameDB := ProcessUserLogin(form.Username, form.Password)
	if check && (len(role) > 0) {
		token, err := utils.GenerateToken(form.Username, role, ip, nameDB)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    500,
				"errors":  err.Error(),
				"message": err.Error(),
			})
		}

		data := make(map[string]string)
		data["access_token"] = token

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"errors":  nil,
			"message": "Success login",
			"data":    data,
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    401,
		"errors":  nil,
		"message": "Error Auth",
	})
}

func ProcessUserLogin(username string, password string) (bool, string, string, string) {
	memberService := member_service.Member{Username: username}
	member, err := memberService.GetmemberLogin()
	if err != nil {
		return false, "", "", ""
	}

	return utils.PasswordVerify(password, member.MemberCredentials.Password), member.Role, member.IP, member.NameDatabase
}

type MemberUpdate struct {
	Birthday     time.Time `json:"birthday" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Image        string    `json:"image" validate:"required"`
	Firstname    string    `json:"firstName" validate:"required"`
	Lastname     string    `json:"lastName" validate:"required"`
	Phone        string    `json:"phone" validate:"required"`
	City         string    `json:"city" validate:"required"`
	Address      string    `json:"address" validate:"required"`
	Address2     string    `json:"address2" validate:"required"`
	Role         string    `json:"role" validate:"required"`
	IP           string    `json:"ip"`
	NameDatabase string    `json:"namedatabase"`
	Status       string    `json:"status" validate:"required"`
}

type SignUpForm struct {
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Image        string `json:"image" validate:"required"`
	Firstname    string `json:"firstName" validate:"required"`
	Lastname     string `json:"lastName" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	City         string `json:"city" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Address2     string `json:"address2" validate:"required"`
	Password     string `json:"password" validate:"required"`
	IP           string `json:"ip"`
	NameDatabase string `json:"namedatabase"`
}

func SignUp(c *fiber.Ctx) error {
	form := &SignUpForm{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err.Error(),
			"message": err.Error(),
		})
	}

	// Validate
	errors := app.ValidateStruct(*form)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  errors,
			"message": errors,
		})

	}

	// Checking
	valid := checkMember(c, *form)

	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": "Username exists.",
		})
	}

	memberService := member_service.Member{
		Username:          form.Username,
		Birthday:          time.Time{},
		Email:             form.Email,
		Image:             form.Image,
		Firstname:         form.Firstname,
		Lastname:          form.Lastname,
		Status:            "",
		Phone:             form.Phone,
		City:              form.City,
		Address:           form.Address,
		Address2:          form.Address2,
		Role:              "",
		IP:                form.IP,
		NameDatabase:      form.NameDatabase,
		ResetKey:          "",
		ResetCount:        0,
		ResetTimestamp:    "",
		ResetKeyTimestamp: "",
		Password:          "",
	}

	if err := memberService.RegisterPublicUser(form.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err,
			"message": "Register false",
		})
	}

	data := make(map[string]string)
	data["member_code"] = form.Username

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"errors":  nil,
		"message": "Success",
		"data":    data,
	})

}

func checkMember(c *fiber.Ctx, form SignUpForm) bool {
	isDuplicateMemberUsername, err := isDuplicateMemberUsername(form.Username)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
		})
		return false
	}

	if isDuplicateMemberUsername {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
		})
		return false
	}

	return true
}

func isDuplicateMemberUsername(username string) (bool, error) {
	memberService := member_service.Member{Username: username}
	member, err := memberService.GetMemberUsername()

	if err != nil {
		return true, err
	}

	if member.Username == username {
		return true, nil
	}

	return false, nil
}

// Get All User
func GetAllUser_Router(c *fiber.Ctx) error {
	item := member_service.Member{}
	data, err := item.GetUser_Service()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Get ID User
func GetByIdUser_Router(c *fiber.Ctx) error {
	item := member_service.Member{}
	data, err := item.GetByIdUser_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Search Username User
func SearchUser_Router(c *fiber.Ctx) error {
	formSearch := new(LoginForm)
	item := member_service.Member{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := item.SearchMember_Service(formSearch.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Update User
func PutMembers(c *fiber.Ctx) error {
	form := &MemberUpdate{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	memberService := member_service.Member{
		Birthday:     form.Birthday,
		Email:        form.Email,
		Image:        form.Image,
		Firstname:    form.Firstname,
		Lastname:     form.Lastname,
		Phone:        form.Phone,
		City:         form.City,
		Address:      form.Address,
		Address2:     form.Address2,
		Role:         form.Role,
		IP:           form.IP,
		NameDatabase: form.NameDatabase,
		Status:       form.Status,
	}

	err := memberService.UpdatMembers(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)
	data["ItemCode"] = form.Email

	return c.Status(fiber.StatusOK).JSON(data)
}
