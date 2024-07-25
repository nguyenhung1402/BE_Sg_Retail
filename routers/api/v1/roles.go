package v1

import (
	"sap-crm/service/casbin_service"

	"github.com/gofiber/fiber/v2"
)

type RoleForm struct {
	Ptype    string `json:"ptype" validate:"required,min=1,max=1"`
	RoleName string `json:"roleName" validate:"required"`
	Path     string `json:"path" validate:"required"`
	Method   string `json:"method" validate:"required"`
}

func AddRole(c *fiber.Ctx) error {
	form := &RoleForm{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err.Error(),
			"message": err.Error(),
		})
	}

	casbinService := casbin_service.MyCasbin{
		Ptype:    form.Ptype,
		RoleName: form.RoleName,
		Path:     form.Path,
		Method:   form.Method,
	}

	check, err := casbinService.AddCasbin()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err,
			"message": "Register roles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"errors":  nil,
		"message": "Success",
		"data":    check,
	})
}
