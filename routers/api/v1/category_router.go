package v1

import (
	category_service "sap-crm/service/category_service"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	CategoryCode string `form:"CategoryCode" valid:"required"`
	CategoryName string `form:"CategoryName" valid:"required"`
	Status       string `form:"Status" valid:"required"`
}

type CategoryUpdate struct {
	CategoryCode string `form:"CategoryCode" valid:"required"`
	CategoryName string `form:"CategoryName" valid:"required"`
	Status       string `form:"Status" valid:"required"`
}

type FormSearchCategory struct {
	CategoryCode string `form:"CategoryCode" valid:"required json:"CategoryCode"`
	CategoryName string `form:"CategoryName" valid:"required" json:"CategoryName"`
}

func PostCategory(c *fiber.Ctx) error {
	form := &Category{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	categoryService := category_service.Category{
		CategoryCode: form.CategoryCode,
		CategoryName: form.CategoryName,
		Status:       form.Status,
	}
	if err := categoryService.AddCategory(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Add false",
		})
	}

	data := make(map[string]string)
	data["CategoryCode"] = form.CategoryCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutCategory(c *fiber.Ctx) error {
	form := &Category{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	categoryService := category_service.Category{
		CategoryName: form.CategoryName,
	}

	err := categoryService.UpdateCategory(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update false",
		})
	}

	data := make(map[string]string)
	data["CategoryCode"] = form.CategoryCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetCategory_Router(c *fiber.Ctx) error {
	category := category_service.Category{}
	data, err := category.GetCategory_Service()

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

func GetByIdCategory_Router(c *fiber.Ctx) error {
	category := category_service.Category{}
	data, err := category.GetByIdCategory_Service(c.Params("id"))

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

func SearchCategory_Router(c *fiber.Ctx) error {
	FormSearchCategory := new(FormSearchCategory)
	category := category_service.Category{}
	if err := c.BodyParser(FormSearchCategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := category.SearchCategory_Service(FormSearchCategory.CategoryCode, FormSearchCategory.CategoryName)
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
