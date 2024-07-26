package v1

import (
	"sap-crm/service/staff_service"

	"time"

	"github.com/gofiber/fiber/v2"
)

type Staff struct {
	StaffCode  string    `form:"StaffCode" validate:"required"`
	StaffName  string    `form:"StaffName" validate:"required"`
	Group      string    `form:"Group"`
	Address    string    `form:"Address"`
	Phone      string    `form:"Phone"`
	Remarks    string    `form:"Remarks"`
	DocDate    time.Time `form:"DocDate"`
	UserAdd    string    `form:"UserAdd"`
	UserUpdate string    `form:"UserUpdate"`
	Status     bool      `form:"Status"`
	Type       string    `form:"Type"`
}

type FormSearchStaff struct {
	StaffCode string `form:"StaffCode" valid:"required json:"StaffCode"`
	StaffName string `form:"StaffName" valid:"required" json:"StaffName"`
	Type      string `form:"Type" valid:"required" json:"Type"`
}

func PostStaff(c *fiber.Ctx) error {
	form := &Staff{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	staffService := staff_service.Staff{
		StaffCode:  form.StaffCode,
		StaffName:  form.StaffName,
		Group:      form.Group,
		Address:    form.Address,
		Phone:      form.Phone,
		Remarks:    form.Remarks,
		UserAdd:    form.UserAdd,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
	}

	if err := staffService.AddStaff(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)
	data["StaffCode"] = form.StaffCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutStaff(c *fiber.Ctx) error {
	form := &Staff{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	staffService := staff_service.Staff{
		StaffCode:  form.StaffCode,
		StaffName:  form.StaffName,
		Group:      form.Group,
		Address:    form.Address,
		Phone:      form.Phone,
		Remarks:    form.Remarks,
		UserAdd:    form.UserAdd,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
	}

	err := staffService.UpdateStaff(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)
	data["StaffCode"] = form.StaffCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetStaff_Router(c *fiber.Ctx) error {
	item := staff_service.Staff{}
	data, err := item.GetStaff_Service()

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

func GetByIdStaff_Router(c *fiber.Ctx) error {
	item := staff_service.Staff{}
	data, err := item.GetByIdStaff_Service(c.Params("id"))

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

func SearchStaff_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchStaff)
	item := staff_service.Staff{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := item.SearchStaff_Service(formSearch.StaffCode, formSearch.StaffName, formSearch.Type)
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
