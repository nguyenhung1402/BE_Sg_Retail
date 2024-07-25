package v1

import (
	"sap-crm/service/viewnxt_service"

	"github.com/gofiber/fiber/v2"
)

type FormViewNXT struct {
	WhsCode  string  `json:"WhsCode"`
	ItemCode string  `json:"ItemCode"`
	Quantity float64 `json:"Quantity"`
}

type FormPostDataWhsCodeViewNXT struct {
	WhsCode string `json:"WhsCode"`
}

type FormPostDataViewNXT struct {
	WhsCode  string `json:"WhsCode"`
	ItemCode string `json:"ItemCode"`
}

func GetAllNXT_Router(c *fiber.Ctx) error {
	viewNXT := viewnxt_service.ViewNXT{}
	data, err := viewNXT.GetAllNXT_Service()

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

func GetWhsCodeNXT_Router(c *fiber.Ctx) error {
	formSearch := new(FormPostDataWhsCodeViewNXT)
	viewNXT := viewnxt_service.ViewNXT{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := viewNXT.GetWhsCodeNXT_Service(formSearch.WhsCode)
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

func GetWhsItemCodeNXT_Router(c *fiber.Ctx) error {
	formSearch := new(FormPostDataViewNXT)
	viewNXT := viewnxt_service.ViewNXT{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := viewNXT.GetWhsItemCodeNXT_Service(formSearch.WhsCode, formSearch.ItemCode)
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
