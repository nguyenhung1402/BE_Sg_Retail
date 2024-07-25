package v1

import (
	"sap-crm/service/menudetail_service"

	"github.com/gofiber/fiber/v2"
)

type ChildMenu struct {
	IDMenu   uint   `json:"IDMenu"`
	Title    string `json:"Title"`
	Icon     string `json:"Icon"`
	Position string `json:"Position"`
	Rule     string `json:"Rule"`
	Url      string `json:"Url"`
}

func GetByIdChilMenu_Router(c *fiber.Ctx) error {
	podetail := menudetail_service.ChildMenu{}
	data, err := podetail.GetByIdChilMenu_Service(c.Params("id"))

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
