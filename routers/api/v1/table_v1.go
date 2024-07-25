package v1

import (
	table_service "sap-crm/service/table_service"

	"github.com/gofiber/fiber/v2"
)

type Table struct {
	TableNumber string `form:"TableNumber" valid:"required"`
	Remark      string `form:"Remark" valid:"required"`
	Status      string `form:"Status" valid:"required"`
}

type TableUpdate struct {
	TableNumber string `form:"TableNumber" valid:"required"`
	Remark      string `form:"Remark" valid:"required"`
	Status      string `form:"Status" valid:"required"`
}

type FormSearchTable struct {
	TableNumber string `form:"TableNumber" valid:"required json:"TableNumber"`
}

func PostTable(c *fiber.Ctx) error {
	form := &Table{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	tableService := table_service.Table{
		TableNumber: form.TableNumber,
		Remark:      form.Remark,
		Status:      form.Status,
	}
	if err := tableService.AddTable(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Add false",
		})
	}

	data := make(map[string]string)
	data["TableNumber"] = form.TableNumber

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutTable(c *fiber.Ctx) error {
	form := &Table{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	tableService := table_service.Table{
		//TableNumber: form.TableNumber,
		Remark: form.Remark,
		Status: form.Status,
	}

	err := tableService.UpdateTable(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update false",
		})
	}

	data := make(map[string]string)
	data["TableNumber"] = form.TableNumber

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetTable_Router(c *fiber.Ctx) error {
	table := table_service.Table{}
	data, err := table.GetTable_Service()

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

func GetByIdTable_Router(c *fiber.Ctx) error {
	table := table_service.Table{}
	data, err := table.GetByIdTable_Service(c.Params("id"))

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

func SearchTable_Router(c *fiber.Ctx) error {
	FormSearchTable := new(FormSearchTable)
	table := table_service.Table{}
	if err := c.BodyParser(FormSearchTable); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := table.SearchTable_Service(FormSearchTable.TableNumber)
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
