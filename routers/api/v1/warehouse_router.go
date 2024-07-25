package v1

import (
	"sap-crm/service/warehouse_service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Whs struct {
	WhsCode string `form:"WhsCode" validate:"required"`
	WhsName string `form:"WhsName" validate:"required"`
	// Group      string    `form:"Group"`
	Address string `form:"Address"`
	// Phone      string    `form:"Phone"`
	Remarks    string    `form:"Remarks"`
	DocDate    time.Time `form:"DocDate"`
	UserAdd    string    `form:"UserAdd"`
	UserUpdate string    `form:"UserUpdate"`
	Status     bool      `form:"Status"`
	Type       string    `form:"Type"`
}

type FormSearchWhs struct {
	WhsCode string `form:"WhsCode" valid:"required json:"WhsCode"`
	WhsName string `form:"WhsName" valid:"required" json:"WhsName"`
	Type    string `form:"Type" valid:"required" json:"Type"`
}

func PostWhs(c *fiber.Ctx) error {
	form := &Whs{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	whsService := warehouse_service.Whs{
		WhsCode: form.WhsCode,
		WhsName: form.WhsName,
		// Group:      form.Group,
		Address: form.Address,
		// Phone:      form.Phone,
		Remarks:    form.Remarks,
		UserAdd:    form.UserAdd,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
	}

	if err := whsService.AddWhs(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)
	data["WhsCode"] = form.WhsCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutWhs(c *fiber.Ctx) error {
	form := &Whs{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	whsService := warehouse_service.Whs{
		WhsName: form.WhsName,
		// Group:      form.Group,
		Address: form.Address,
		// Phone:      form.Phone,
		Remarks:    form.Remarks,
		UserAdd:    form.UserAdd,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
	}

	err := whsService.UpdatWhs(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)
	data["WhsCode"] = form.WhsCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetWhs_Router(c *fiber.Ctx) error {
	item := warehouse_service.Whs{}
	data, err := item.GetWhs_Service()

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

func GetByIdWhs_Router(c *fiber.Ctx) error {
	item := warehouse_service.Whs{}
	data, err := item.GetByIdWhs_Service(c.Params("id"))

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

func SearchWhs_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchWhs)
	item := warehouse_service.Whs{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := item.SearchWhs_Service(formSearch.WhsCode, formSearch.WhsName, formSearch.Type)
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
