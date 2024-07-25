package v1

import (
	"sap-crm/service/config_settings_app_service"

	"github.com/gofiber/fiber/v2"
)

type ConfigSettingApp struct {
	Company  string `form:"Company" json:"company"`
	Address  string `form:"Address" json:"address"`
	Phone    string `form:"Phone" json:"phone"`
	MST      string `form:"MST" json:"mst"`
	IP       string `form:"IP" json:"ip"`
	Database string `form:"Database" json:"database"`
	UserName string `form:"UserName" json:"username"`
	PassWord string `form:"PassWord" json:"password"`
	Image    string `form:"Image" json:"image"`
}

func PostConfigSettings_v1(c *fiber.Ctx) error {
	form := &ConfigSettingApp{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	itemService := config_settings_app_service.ConfigSettingApp{
		Company:  form.Company,
		Address:  form.Address,
		Phone:    form.Phone,
		MST:      form.MST,
		IP:       form.IP,
		Database: form.Database,
		UserName: form.UserName,
		PassWord: form.PassWord,
		Image:    form.Image,
	}
	if err := itemService.AddSettings_service(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Add false",
		})
	}

	data := make(map[string]string)
	data["Company"] = form.Company

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutConfigSettings(c *fiber.Ctx) error {
	form := &ConfigSettingApp{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	itemService := config_settings_app_service.ConfigSettingApp{
		Company:  form.Company,
		Address:  form.Address,
		Phone:    form.Phone,
		MST:      form.MST,
		IP:       form.IP,
		Database: form.Database,
		UserName: form.UserName,
		PassWord: form.PassWord,
		Image:    form.Image,
	}

	err := itemService.UpdateUpdateSettings_service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update false",
		})
	}

	data := make(map[string]string)
	data["Company"] = form.Company

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetConfigSettings_Router(c *fiber.Ctx) error {
	item := config_settings_app_service.ConfigSettingApp{}
	data, err := item.GetSettings_Service()

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
func GetByIdSettings_Router(c *fiber.Ctx) error {
	item := config_settings_app_service.ConfigSettingApp{}
	data, err := item.GetByIdSettings_Service(c.Params("id"))

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
