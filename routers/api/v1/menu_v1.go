package v1

import (
	"sap-crm/service/menu_service"

	"github.com/gofiber/fiber/v2"
)

type FormParentMenu struct {
	Title     string          `json:"Title"`
	Icon      string          `json:"Icon"`
	Position  string          `json:"Position"`
	Rule      string          `json:"Rule"`
	ChildMenu []FormChildMenu `json:"ChildMenu"`
}

type FormChildMenu struct {
	IDMenu   int    `json:"IDMenu"`
	Title    string `json:"Title"`
	Icon     string `json:"Icon"`
	Position string `json:"Position"`
	Rule     string `json:"Rule"`
	Url      string `json:"Url"`
}

// ADD Menu
func PostMenu(c *fiber.Ctx) error {
	form := &FormParentMenu{}

	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var menuDetails []menu_service.ChildMenu = make([]menu_service.ChildMenu, len(form.ChildMenu))
	for i, v := range form.ChildMenu {
		item := menu_service.ChildMenu{
			IDMenu:   v.IDMenu,
			Title:    v.Title,
			Icon:     v.Icon,
			Position: v.Position,
			Rule:     v.Rule,
			Url:      v.Url,
		}
		menuDetails[i] = item
	}
	menuService := menu_service.ParentMenu{
		Title:     form.Title,
		Icon:      form.Icon,
		Position:  form.Position,
		Rule:      form.Rule,
		ChildMenu: menuDetails,
	}

	if err := menuService.AddParentMenu_Service(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Post Add Menu false",
		})
	}

	data := make(map[string]string)
	data["Title"] = form.Title

	return c.Status(fiber.StatusOK).JSON(data)
}

// get all po
func GetMenu_Router(c *fiber.Ctx) error {
	parentMenu := menu_service.ParentMenu{}
	data, err := parentMenu.GetMenu_Service(c.Params("id"))

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

func GetByIdMenu_Router(c *fiber.Ctx) error {
	parentMenu := menu_service.ParentMenu{}
	data, err := parentMenu.GetByIdMenu_Service(c.Params("id"))

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
